// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract FundingVaultProxyStorage {
  // slot 0x00 - manager address (admin)
  address internal _manager;
  uint96 internal __unused0;
  // slot 0x01 - implementation address
  address internal _implementation;
  uint96 internal __unused1;
}
