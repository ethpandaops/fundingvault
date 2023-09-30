# Holešovice Funding Vault

This is the home for the holesky funding vault contract: [0x610866c6089768dA95524bcc4cE7dB61eDa3931c](https://holesky.etherscan.io/address/0x610866c6089768da95524bcc4ce7db61eda3931c)

The purpose of this project is to give smaller faucets and projects with a constant need for testnet funds a reliable source for these funds in the [Holešovice Testnet](https://github.com/eth-clients/holesky).

Golden Rule: Funds may never be sold for profit or hoarded with no use.\
Testnet funds are public goods and should be shared freely with entities that have an actual need for them.

## Technical concept

The funding vault consists of 2 contracts:
* The [`FundingVaultProxy`](https://holesky.etherscan.io/address/0x610866c6089768da95524bcc4ce7db61eda3931c) / `FundingVaultV1` upgradable contract, which holds the funding reserves and fulfills fund requests (claims).
* The [`FundingVaultToken`](https://holesky.etherscan.io/address/0x97652a83cc29043fa9be2781cc0038eba70de911) contract, that provides a ERC721 token which gives permission to claim the allowed funds from the vault.

If you're running a project that has a ongoing need for testnet funds (either for development teams or low audience faucets),
open a issue in this repository to apply for a funding grant. Supply some information about what you need these funds for and how much ETH per month you need. The needs shouldn't exceed a max. of 10k HolETH/month.

If your application is accepted, you'll receive a ERC721 token ("NFT") on holesky, which gives you access to the granted amount of funds when needed. \
You're responsible for keeping this NFT secure, but you're allowed to transfer it where-ever you want.

If no rules are broken and your project remains active, the grant will persist till the planned end of holesky in Dec 2028.


The holder of the NFT is allowed to claim funds from the funding vault (in respect of the granted limits).\
To do so, one of the claim functions on the [Funding Vault Contract](https://holesky.etherscan.io/address/0x610866c6089768da95524bcc4ce7db61eda3931c) needs to be called from the wallet that holds the NFT.

Available `claim` functions:
* `claim(uint256 amount)` - request the amount of funds specified by `amount` (in wei) and send to the sender wallet.
* `claimTo(uint256 amount, address target)` - request the amount of funds specified by `amount` (in wei) and send to `target`.

When supplying a amount of `0`, the whole allowed amount will be paid out.

The contract is time based. When claiming the full available amount, you can already claim again a few secs later. But in that case, only the funds that piled up during these few secs will be available.

If you've been granted 10k HolETH/month, you have to wait a day to claim ~333 HolETH, or wait 2 days to claim 666 HolETH, etc...