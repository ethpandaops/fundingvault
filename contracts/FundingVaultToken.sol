// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "./IFundingVaultToken.sol";
import "./IFundingVault.sol";

contract FundingVaultToken is ERC721Enumerable, IFundingVaultToken {
  address private _fundingVault;

  constructor(address fundingVault) ERC721("FundingVault Grant", "Funding Grant") {
    _fundingVault = fundingVault;
  }

  receive() external payable {
    if(msg.value > 0) {
      (bool sent, ) = payable(_fundingVault).call{value: msg.value}("");
      require(sent, "failed to forward ether");
    }
  }

  function getVault() public view returns (address) {
    return _fundingVault;
  }

  function _baseURI() internal view override returns (string memory) {
    return string(abi.encodePacked("https://dev.pk910.de/ethvault?c=", Strings.toString(block.chainid), "&v=",  Strings.toHexString(uint160(_fundingVault), 20), "&p="));
  }

  function _beforeTokenTransfer(address from, address to, uint256 tokenId, uint256 batchSize) internal virtual override {
    super._beforeTokenTransfer(from, to, tokenId, batchSize);

    IFundingVault(_fundingVault).notifyGrantTransfer(uint64(tokenId));
  }

  function tokenUpdate(uint64 tokenId, address targetAddr) public {
    require(_msgSender() == _fundingVault, "not vault contract");

    if(targetAddr != address(0)) {
      if(!_exists(tokenId)) {
        _safeMint(targetAddr, tokenId);
      }
      else if(_ownerOf(tokenId) != targetAddr) {
        _safeTransfer(_ownerOf(tokenId), targetAddr, tokenId, "");
      }
    }
    else if(_exists(tokenId)) {
      _burn(tokenId);
    }
  }

}
