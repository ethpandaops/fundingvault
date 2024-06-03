# Funding Vault

Welcome to the repository for the Funding Vault contract, actively deployed on the [Holešovice](https://holesky.etherscan.io/address/0x610866c6089768da95524bcc4ce7db61eda3931c) and [Sepolia](https://sepolia.etherscan.io/address/0x610866c6089768da95524bcc4ce7db61eda3931c) Testnets. This project serves as a reliable source of testnet funds for smaller faucets and projects that require a steady influx of funds.

**Golden Rule**: Testnet funds may never be sold for profit or hoarded. They are public goods and should be utilized responsibly by entities with genuine needs.

## Usage

Entities eligible for funding can claim funds regularly either through the [Web UI](https://fundingvault.ethpandaops.io/) or programmatically via direct calls to the vault contract.

### Applying for a Grant

To ensure the integrity and purpose of the fund allocation, applicants must meet the following criteria to be eligible for a grant:

- **Website**: Provide a link to a functioning website with comprehensive information about the project or company.
- **Project Description**: Include a concise description of your project and a detailed explanation of how the funds will be used.
- **Working Demo/Implementation**: Showcase a working demo or an implementation of the project part that requires ongoing funding to demonstrate its functionality and relevance.
- **Protection Methods for Faucets**: If applying for a faucet, describe the methods employed to protect against abuse. Note that simple captcha protection is generally insufficient.

If your project meets these criteria and needs ongoing testnet funds (for development teams or low-traffic faucets), please open an issue in this repository with details about your requirements and the amount of ETH needed.

Upon approval, you will receive an ERC721 token ("NFT") that grants access to the specified funds. It is your responsibility to secure this NFT, although you may transfer it as needed.

**Ineligible Uses**:
- **Token Liquidity**: Providing liquidity for tokens is not an appropriate use of these funds.
- **Top-list Placement**: Funding intended to maintain a position on any kind of top-list is also considered inappropriate.

Grants will continue as long as:
- No rules are violated.
- Funds are used appropriately.
- Your project remains active.
The grant is designed to last until the planned end of Sepolia in December 2026 and Holešovice in December 2028.

## Programmatic Claims

Holders of the Grant NFT can claim funds within the granted limits by calling functions on the Funding Vault Contract.

**Contract Address**: `0x610866c6089768dA95524bcc4cE7dB61eDa3931c`

### Available Claim Functions:
- `claim(uint256 amount)`: Request and send the specified amount of funds (in wei) to the wallet initiating the call.
- `claimTo(uint256 amount, address target)`: Request and send the specified amount of funds (in wei) to a target address.

Specifying an amount of `0` will trigger a payout of all available funds.

### Timing of Claims
The contract operates on a time-based system. If you claim the full available amount, subsequent claims can be made within seconds, but only the funds that have accumulated since the last claim will be available. For instance, with a grant of 10k HolETH/month, you would need to wait a day to claim approximately 333 HolETH, or two days for about 666 HolETH, and so on.

## Credits

A big thanks to EF Testing for their testing efforts and to Nethermind Security for [auditing the smart contract](https://github.com/ethpandaops/fundingvault/blob/master/fundingvault/audit/NM-0234-Ethereum-Foundation-Final.pdf)!
