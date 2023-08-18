// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

interface IFundingVault {
  function notifyGrantTransfer(uint64 grantId) external;
}
