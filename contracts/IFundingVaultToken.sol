// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

import "@openzeppelin/contracts/token/ERC721/extensions/IERC721Enumerable.sol";

interface IFundingVaultToken is IERC721Enumerable {
  function tokenUpdate(uint64 tokenId, address targetAddr) external;
}
