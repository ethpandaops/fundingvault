// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

/*
##################################################################
#                Holešovice Funding Vault                        #
#                                                                #
# This contract is used to distribute fund reserves to faucets   #
# or other projects that have a ongoing need for testnet funds.  #
#                                                                #
#  Vault contract:  0x610866c6089768dA95524bcc4cE7dB61eDa3931c   #
#                                                                #
# see https://dev.pk910.de/ethvault               by pk910.eth   #
##################################################################
*/

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "./ReentrancyGuard.sol";
import "./FundingVaultProxyStorage.sol";
import "./IFundingVaultToken.sol";
import "./IFundingVault.sol";

struct Grant {
  uint64 claimTime;
  uint64 claimInterval;
  uint128 claimLimit;
  uint256 dustBalance;
}

contract FundingVaultStorage {
  // slot 0x05
  address internal _vaultTokenAddr;
  uint64 internal _grantIdCounter;
  uint32 internal _claimTransferLockTime;

  // slot 0x06
  uint128 internal _managerLimitAmount;
  uint64 internal _managerLimitInterval;
  uint32 internal _managerGrantCooldown;
  uint32 internal _managerGrantCooldownLock;

  // slot 0x07
  mapping(uint64 => Grant) internal _grants;

  // slot 0x08
  mapping(uint64 => uint64) internal _grantClaimLock;

  // slot 0x09
  mapping(address => uint64) internal _managerCooldown;

  // slot 0x0a
  mapping(uint64 => uint256) internal _grantTotalClaimed;

  // slot 0x0b
  mapping(uint64 => bytes32) internal _grantNames;
}

contract FundingVaultV1 is 
  FundingVaultProxyStorage, // 0x00 - 0x01
  AccessControl,            // 0x02
  Pausable,                 // 0x03
  ReentrancyGuard,          // 0x04
  FundingVaultStorage,      // 0x05 - 0x0b
  IFundingVault
{
  bytes32 public constant GRANT_MANAGER_ROLE = keccak256("GRANT_MANAGER_ROLE");

  event GrantLock(uint64 indexed grantId, uint64 lockTime, uint64 lockTimeout);
  event GrantUpdate(uint64 indexed grantId, uint128 amount, uint64 interval);
  event GrantClaim(uint64 indexed grantId, address indexed to, uint256 amount, uint64 grantTimeUsed);

  receive() external payable {
  }

  function initialize(address tokenAddr) public {
    require(_reentrancyStatus == 0 && _grantIdCounter == 0, "already initialized");
    require(_manager == _msgSender(), "access denied");
    _grantRole(DEFAULT_ADMIN_ROLE, _manager);
    _reentrancyStatus = 1;
    _vaultTokenAddr = tokenAddr;
    _grantIdCounter = 1;
    _claimTransferLockTime = 600;
    _managerLimitAmount = 100000;
    _managerLimitInterval = 2592000;
    _managerGrantCooldown = 86400;
    _managerGrantCooldownLock = 43200;
  }


  //## Admin configuration / rescue functions

  function rescueCall(address addr, uint256 amount, bytes calldata data) public onlyRole(DEFAULT_ADMIN_ROLE) {
    uint balance = address(this).balance;
    require(balance >= amount, "amount exceeds wallet balance");

    (bool sent, ) = payable(addr).call{value: amount}(data);
    require(sent, "call failed");
  }

  function setPaused(bool paused) public onlyRole(DEFAULT_ADMIN_ROLE) {
    if(paused) {
      _pause();
    } else {
      _unpause();
    }
  }

  function setProxyManager(address manager) public onlyRole(DEFAULT_ADMIN_ROLE) {
    _manager = manager;
  }

  function setClaimTransferLockTime(uint32 lockTime) public onlyRole(DEFAULT_ADMIN_ROLE) {
    _claimTransferLockTime = lockTime;
  }

  function setManagerGrantLimits(uint128 amount, uint64 interval, uint32 cooldown, uint32 cooldownLock) public onlyRole(DEFAULT_ADMIN_ROLE) {
    _managerLimitAmount = amount;
    _managerLimitInterval = interval;
    _managerGrantCooldown = cooldown;
    _managerGrantCooldownLock = cooldownLock;
  }

  
  //## Internal helper functions

  function _ownerOf(uint64 tokenId) internal view returns (address) {
    return IFundingVaultToken(_vaultTokenAddr).ownerOf(tokenId);
  }

  function _getTime() internal view returns (uint64) {
    return uint64(block.timestamp);
  }

  /*
  The _calculateClaim function is the central piece of code that does the calculations for claiming funds via a grant.
    Arguments: 
      grantId - the grant id the sender likes to claim from
      requestAmount - the desired amount of funds the sender likes to claim (0 to claim all available)
    Return Values:
      claimAmount - the amount of funds for payout, smaller or equal to requestedAmount, max available if requestAmount is 0
      newClaimTime - the new claimTime, must be set to the grant struct if claimAmount is payed out
      newDustBalance - the new dustBalance, must be set to the grant struct if claimAmount is payed out
      usedTime - the used claim time to fulfil the request (more informative and for debugging)
  */
  function _calculateClaim(uint64 grantId, uint256 requestAmount) public view 
    returns (uint256 claimAmount, uint64 newClaimTime, uint256 newDustBalance, uint64 usedTime) {
    Grant memory grant = _grants[grantId];
    require(grant.claimInterval > 0 && grant.claimLimit > 0 && grant.claimTime > 0, "invalid grant");
    
    uint256 claimLimit = grant.claimLimit * 1 ether;
    if(requestAmount > claimLimit) {
      requestAmount = claimLimit;
    }

    uint64 time = _getTime();
    if(_grantClaimLock[grantId] > time) {
      // grant locked
      newClaimTime = grant.claimTime;
      usedTime = 0;
      claimAmount = 0;
      newDustBalance = grant.dustBalance;
    }
    else {
      uint64 baseClaimTime = grant.claimTime;
      uint64 availableTime = time - baseClaimTime;
      uint256 dustBalance = grant.dustBalance;
      if(availableTime > grant.claimInterval) {
        // available time exceeds interval
        // the sender claimed less than granted, the unclaimed amount is no longer available 
        availableTime = grant.claimInterval;
        baseClaimTime = time - grant.claimInterval;
        dustBalance = 0;
      }

      if(requestAmount != 0 && requestAmount <= dustBalance) {
        // take from dust balance
        newClaimTime = baseClaimTime;
        usedTime = 0;
        claimAmount = requestAmount;
        newDustBalance = dustBalance - requestAmount;
      }
      else {
        // get max claimable amount
        claimAmount = (claimLimit * availableTime / grant.claimInterval) + dustBalance;

        if(requestAmount != 0 && requestAmount < claimAmount) {
          // sender requested less than available, "partial" claim
          uint256 requestClaimAmount = requestAmount - dustBalance;
          usedTime = uint64(requestClaimAmount * grant.claimInterval / claimLimit);
          if(usedTime * claimLimit / grant.claimInterval < requestClaimAmount) {
            usedTime++; // round up if there is a rounding gap in ETH amount
            newDustBalance = (usedTime * claimLimit / grant.claimInterval) - requestClaimAmount;
          }
          else {
            newDustBalance = 0;
          }
          require(usedTime <= availableTime, "calculation error: usedTime > availableTime");

          newClaimTime = baseClaimTime + usedTime;
          claimAmount = requestAmount;
        }
        else {
          // sender requested all available funds
          usedTime = availableTime;
          newClaimTime = time;
          newDustBalance = 0;
        }
      }
    }
  }


  //## Public view functions

  function getVaultToken() public view returns (address) {
    return _vaultTokenAddr;
  }

  function getGrants() public view returns (Grant[] memory) {
    IFundingVaultToken vaultToken = IFundingVaultToken(_vaultTokenAddr);
    uint64 grantCount = uint64(vaultToken.totalSupply());
    Grant[] memory grants = new Grant[](grantCount);
    for(uint64 grantIdx = 0; grantIdx < grantCount; grantIdx++) {
      uint64 grantId = uint64(vaultToken.tokenByIndex(grantIdx));
      grants[grantIdx] = _grants[grantId];
    }
    return grants;
  }

  function getGrant(uint64 grantId) public view returns (Grant memory) {
    require(_grants[grantId].claimTime > 0, "grant not found");
    return _grants[grantId];
  }

  function getGrantName(uint64 grantId) public view returns (bytes32) {
    require(_grants[grantId].claimTime > 0, "grant not found");
    return _grantNames[grantId];
  }

  function getGrantTotalClaimed(uint64 grantId) public view returns (uint256) {
    return _grantTotalClaimed[grantId];
  }

  function getGrantLockTime(uint32 grantId) public view returns (uint64) {
    require(_grants[grantId].claimTime > 0, "grant not found");
    if(_grantClaimLock[grantId] > uint64(block.timestamp)) {
      return _grantClaimLock[grantId] - uint64(block.timestamp);
    }
    else {
      return 0;
    }
  }

  function getClaimableBalance() public view returns (uint256) {
    uint256 claimableAmount = 0;
    IFundingVaultToken vaultToken = IFundingVaultToken(_vaultTokenAddr);

    uint64 grantCount = uint64(vaultToken.balanceOf(_msgSender()));
    for(uint64 grantIdx = 0; grantIdx < grantCount; grantIdx++) {
      uint64 grantId = uint64(vaultToken.tokenOfOwnerByIndex(_msgSender(), grantIdx));
      claimableAmount += _claimableBalance(grantId);
    }
    return claimableAmount;
  }

  function getClaimableBalance(uint64 grantId) public view returns (uint256) {
    require(_grants[grantId].claimTime > 0, "grant not found");
    return _claimableBalance(grantId);
  }

  function _claimableBalance(uint64 grantId) internal view returns (uint256) {
    (uint256 claimAmount, , , ) = _calculateClaim(grantId, 0);
    return claimAmount;
  }

  function getManagerCooldown(address manager) public view returns (uint64) {
    if(_managerCooldown[manager] <= _getTime()) {
      return 0;
    }
    return _managerCooldown[manager] - _getTime();
  }

  function getManagerGrantLimits() public view returns (uint128, uint64, uint32, uint32) {
    return (
      _managerLimitAmount,
      _managerLimitInterval,
      _managerGrantCooldown,
      _managerGrantCooldownLock
    );
  }

  //## Grant managemnet functions (Grant Manager)

  function createGrant(address addr, uint128 amount, uint64 interval, bytes32 name) public onlyRole(GRANT_MANAGER_ROLE) nonReentrant {
    require(amount > 0 && interval > 0, "invalid grant");
    uint256 grantQuota = uint256(amount) * 1 ether / interval;
    uint256 managerQuota = uint256(_managerLimitAmount) * 1 ether / _managerLimitInterval;

    if(!hasRole(DEFAULT_ADMIN_ROLE, _msgSender())) {
      _requireNotPaused();

      if(interval > _managerLimitInterval) {
        // special case, if a grant with an interval bigger than the manager limit interval is created
        // increase the grantQuota as if the grant would have been created with the manager limit interval
        // this avoids managers from exploiting the contract by creating multiple grants with extremely high intervals
        grantQuota = uint256(amount) * 1 ether / _managerLimitInterval;
      }

      // check if granted amount exceeds manager limits
      require(amount <= _managerLimitAmount, "amount exceeds manager limits");
      require(grantQuota <= managerQuota, "quota exceeds manager limits");
      require(_managerCooldown[_msgSender()] < _getTime() + _managerGrantCooldownLock, "manager cooldown");
    }
    if(_managerCooldown[_msgSender()] < _getTime()) {
      _managerCooldown[_msgSender()] = _getTime();
    }
    _managerCooldown[_msgSender()] += uint64(_managerGrantCooldown * grantQuota / managerQuota) + 1;

    uint64 grantId = _grantIdCounter++;
    _grants[grantId] = Grant({
      claimTime: _getTime() - interval,
      claimInterval: interval,
      claimLimit: amount,
      dustBalance: 0
    });
    _grantNames[grantId] = name;
    IFundingVaultToken(_vaultTokenAddr).tokenUpdate(grantId, addr);

    emit GrantUpdate(grantId, amount, interval);
  }

  function updateGrant(uint64 grantId, uint128 amount, uint64 interval) public onlyRole(GRANT_MANAGER_ROLE) nonReentrant {
    require(_grants[grantId].claimTime > 0, "grant not found");
    require(amount > 0 && interval > 0, "invalid grant");

    uint256 oldQuota = uint256(_grants[grantId].claimLimit) * 1 ether / _grants[grantId].claimInterval;
    uint256 newQuota = uint256(amount) * 1 ether / interval;
    uint256 managerQuota = uint256(_managerLimitAmount) * 1 ether / _managerLimitInterval;
    bool isIncrease = newQuota > oldQuota;

    if(!hasRole(DEFAULT_ADMIN_ROLE, _msgSender())) {
      _requireNotPaused();
      // check if granted amount exceeds manager limits
      require(amount <= _managerLimitAmount, "amount exceeds manager limits");
      require(newQuota <= managerQuota, "quota exceeds manager limits");

      if(isIncrease) {
        require(_managerCooldown[_msgSender()] < _getTime() + _managerGrantCooldownLock, "manager cooldown");
      }
    }
    if(isIncrease) {
      if(_managerCooldown[_msgSender()] < _getTime()) {
        _managerCooldown[_msgSender()] = _getTime();
      }
      _managerCooldown[_msgSender()] += uint64(_managerGrantCooldown * (newQuota - oldQuota) / managerQuota) + 1;
    }

    _grants[grantId].claimInterval = interval;
    _grants[grantId].claimLimit = amount;

    emit GrantUpdate(grantId, amount, interval);
  }

  function transferGrant(uint64 grantId, address addr) public onlyRole(GRANT_MANAGER_ROLE) nonReentrant {
    require(_grants[grantId].claimTime > 0, "grant not found");

    uint256 grantQuota = uint256(_grants[grantId].claimLimit) * 1 ether / _grants[grantId].claimInterval;
    uint256 managerQuota = uint256(_managerLimitAmount) * 1 ether / _managerLimitInterval;

    if(!hasRole(DEFAULT_ADMIN_ROLE, _msgSender())) {
      _requireNotPaused();
      // check if grant quota exceeds manager limits
      require(_grants[grantId].claimLimit <= _managerLimitAmount, "quota exceeds manager limits");
      require(grantQuota <= managerQuota, "quota exceeds manager limits");
      require(_managerCooldown[_msgSender()] < _getTime() + _managerGrantCooldownLock, "manager cooldown");
    }
    if(_managerCooldown[_msgSender()] < _getTime()) {
      _managerCooldown[_msgSender()] = _getTime();
    }
    _managerCooldown[_msgSender()] += uint64(_managerGrantCooldown * grantQuota / managerQuota) + 1;

    IFundingVaultToken(_vaultTokenAddr).tokenUpdate(grantId, addr);
  }

  function removeGrant(uint64 grantId) public onlyRole(GRANT_MANAGER_ROLE) nonReentrant {
    require(_grants[grantId].claimTime > 0, "grant not found");

    if(!hasRole(DEFAULT_ADMIN_ROLE, _msgSender())) {
      _requireNotPaused();
    }

    IFundingVaultToken(_vaultTokenAddr).tokenUpdate(grantId, address(0));
    delete _grants[grantId];
  }

  function renameGrant(uint64 grantId, bytes32 name) public onlyRole(GRANT_MANAGER_ROLE) nonReentrant {
    require(_grants[grantId].claimTime > 0, "grant not found");

    if(!hasRole(DEFAULT_ADMIN_ROLE, _msgSender())) {
      _requireNotPaused();
    }

    _grantNames[grantId] = name;
  }

  function lockGrant(uint64 grantId, uint64 lockTime) public nonReentrant {
    require(_grants[grantId].claimTime > 0, "grant not found");
    require(
      _msgSender() == _ownerOf(grantId) || 
      hasRole(GRANT_MANAGER_ROLE, _msgSender())
    , "not grant owner or manager");

    if(!hasRole(DEFAULT_ADMIN_ROLE, _msgSender())) {
      _requireNotPaused();
    }

    _lockGrant(grantId, lockTime);
  }

  function notifyGrantTransfer(uint64 grantId) public {
    require(_msgSender() == _vaultTokenAddr, "not token contract");
    _lockGrant(grantId, _claimTransferLockTime);
  }

  function _lockGrant(uint64 grantId, uint64 lockTime) internal {
    uint64 lockTimeout = _getTime() + lockTime;
    if(lockTimeout > _grantClaimLock[grantId] || hasRole(DEFAULT_ADMIN_ROLE, _msgSender())) {
      _grantClaimLock[grantId] = lockTimeout;
    }
    else {
      lockTime = 0;
      lockTimeout = _grantClaimLock[grantId];
    }
    emit GrantLock(grantId, lockTime, lockTimeout);
  }

  
  //## Public claim functions

  function claim(uint256 amount) public whenNotPaused nonReentrant returns (uint256) {
    uint256 claimAmount = _claimFrom(_msgSender(), amount, _msgSender());
    if(amount > 0) {
      require(claimAmount == amount, "claim failed");
    }
    else {
      require(claimAmount > 0, "claim failed");
    }
    return claimAmount;
  }

  function claim(uint64 grantId, uint256 amount) public whenNotPaused nonReentrant returns (uint256) {
    require(_grants[grantId].claimTime > 0, "grant not found");
    require(_ownerOf(grantId) == _msgSender(), "not owner of this grant");

    uint256 claimAmount = _claim(grantId, amount, _msgSender());
    if(amount > 0) {
      require(claimAmount == amount, "claim failed");
    }
    else {
      require(claimAmount > 0, "claim failed");
    }
    return claimAmount;
  }

  function claimTo(uint256 amount, address target) public whenNotPaused nonReentrant returns (uint256) {
    uint256 claimAmount = _claimFrom(_msgSender(), amount, target);
    if(amount > 0) {
      require(claimAmount == amount, "claim failed");
    }
    else {
      require(claimAmount > 0, "claim failed");
    }
    return claimAmount;
  }

  function claimTo(uint64 grantId, uint256 amount, address target) public whenNotPaused nonReentrant returns (uint256) {
    require(_grants[grantId].claimTime > 0, "grant not found");
    require(_ownerOf(grantId) == _msgSender(), "not owner of this grant");

    uint256 claimAmount = _claim(grantId, amount, target);
    if(amount > 0) {
      require(claimAmount == amount, "claim failed");
    }
    else {
      require(claimAmount > 0, "claim failed");
    }
    return claimAmount;
  }

  function _claimFrom(address owner, uint256 amount, address target) internal returns (uint256) {
    uint256 claimAmount = 0;
    IFundingVaultToken vaultToken = IFundingVaultToken(_vaultTokenAddr);

    uint64 grantCount = uint64(vaultToken.balanceOf(owner));
    for(uint64 grantIdx = 0; grantIdx < grantCount; grantIdx++) {
      uint64 grantId = uint64(vaultToken.tokenOfOwnerByIndex(owner, grantIdx));
      uint256 claimed = _claim(grantId, amount, target);
      claimAmount += claimed;
      if(amount > 0) {
        if(amount == claimed) {
          break;
        }
        else {
          amount -= claimed;
        }
      }
    }
    return claimAmount;
  }

  function _claim(uint64 grantId, uint256 amount, address target) internal returns (uint256) {
    (uint256 claimAmount, uint64 newClaimTime, uint256 newDustBalance, uint64 usedClaimTime) = _calculateClaim(grantId, amount);
    if(claimAmount == 0) {
      return 0;
    }

    // update grant struct
    _grants[grantId].claimTime = newClaimTime;
    _grants[grantId].dustBalance = newDustBalance;
    _grantTotalClaimed[grantId] += claimAmount;

    // send claim amount to target
    (bool sent, ) = payable(target).call{value: claimAmount}("");
    require(sent, "failed to send ether");

    // emit claim event
    emit GrantClaim(grantId, target, claimAmount, usedClaimTime);

    return claimAmount;
  }

}
