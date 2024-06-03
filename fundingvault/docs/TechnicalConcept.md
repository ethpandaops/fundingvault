# Technical Concept

## Description

The FundingVault contract provides a way to distribute continuous limited amounts of funds to authorized entities. The distribution is time gated and a specific limit per grant is enforced (eg. 50k ETH per month).

Grants are represented as ERC721 Tokens ("NFTs"), which can be created by contract managers based on a accepted application / approved funding need.

The Grant NFT allows access to the granted amount of funds. Grantees are expected to store the NFT safely, but may transfer it to other wallets based on their needs.

The fund request process is pull based. Funds can be requested from the FundingVault contract whenever needed. Whoever holds the Grant NFT is able to request the granted amount of funds.

## Contracts

The FundingVault consists of three contracts:
* `FundingVaultProxy`  [0x610866c6089768dA95524bcc4cE7dB61eDa3931c](https://holesky.etherscan.io/address/0x610866c6089768da95524bcc4ce7db61eda3931c)
Proxy contract that holds Funds and serves as entrypoint for FundingVault calls.
* `FundingVaultV1`  [0x93Af84598dda401de8c2ecC87052B8506E83D064](https://holesky.etherscan.io/address/0x93Af84598dda401de8c2ecC87052B8506E83D064)
Implementation of the FundingVault logic (V1).
* `FundingVaultToken` [0x97652A83CC29043fA9Be2781cc0038EBa70de911](https://holesky.etherscan.io/address/0x97652A83CC29043fA9Be2781cc0038EBa70de911)
Token contract that provides a ERC721 token which gives permission to claim the allowed funds from the vault.

## Upgradability

The FundingVault contracts are deployed in a way that allows upgrades & fixes of the core logic when needed.
This also allows changing / extending the distribution logic if needed in future.

## User Groups
There are 3 user groups (security roles) in the contract:
* Grantee
A wallet that holds a Grant NFT and is allowed to request funds (up to the granted limits).

* Manager
A wallet that is allowed to manage grants (create, update, delete, transfer) up to a certain limit.

* Admin
A (cold/multisig-)wallet with full access to all funds and emergency functions.
May upgrade the contract or grant/revoke manager access.
May create/update/transfer grants without any limitation.


The manager & admin groups are represented by `AccessControl` roles within the contract.
The grantee group is represented by the Grant NFT ownership, so whoever holds a Grant NFT is part of this group.

## Security Concept

The FundingVault contract is intended to hold a large amount of funds over the lifetime of the network and drip it slowly over time to authorized entities.
It should be ensured, that even in case of a security breach of one of the grantees or managers hot wallets, the majority of funds should be safely stored in the contract.
It should be impossible to drain the contract completely in a short period of time just via a low secured user wallet.

To achieve that, there are two limitation stategies:
- Grants are limited to a specific amount of funds per time period.
  This is the most obvious limitation, which prevents requesting a big amounts of funds in a short time frame.
  In case of a security breach or if the grant NFT gets stolen, only the allowed amount of funds at the moment of the breach is lost.
  The NFT can be locked and transfered back to a verified wallet by contract managers afterwards.
- Contract managers are limited to a specific amount of funds per time period too.
  This limit applies to the creation, modification and forceful transfer of grants.
  Actions with negative impact like deleting or locking a grant are not affected.
  The manager limit defaults to managing grants worth 100k ETH per month within one day, but can be adjusted by the contract admin when needed.
  Given the limit of 100k ETH/month, a manager may either:
    - create 2 grants with 50k ETH per month
    - create 1 grant with ~12.5k ETH per week
    - create 1 grant with 50k ETH per month and transfer another grant of 50k/month to a new wallet
    - delete or lock all grants
  
  ... or other combinations of actions, as long as they're not exceeding the limit.
  There is a lock threshold of 12h to avoid unnecesarry delay when managing multiple grants in one go. 
  Given a limit of 100k/Month per day, that means locking only applies after managing grants worth 50k/Month. 
  Beyond the threshold, the locking time gets applied lineary following the rules above.
  
  The manager limit enforces, that even in case of a manager wallet breach, the majority of funds in the Vault cannot be stolen immediatly.
  When exceeding the specified limit, managers are locked for a certain period of time before being allowed to create/update/transfer further grants ("cooldown").
  Grants that exceed the manager limits on its own (eg. 200k/Month) cannot be created by managers themselves. Such big grants are expected to be created by the owner wallet, which is intended to be a high secured team multisig.


## Vault Contract

### Grantee Functions
Grantee functions are accessible to all wallets that hold at least one grant NFT.
The grantee functions provide different ways to claim the granted funds.

| Identifer | Arguments | Description |
|----------|------------|-------------|
| claim<br>`0x379607f5` | amount | Claim `amount` wei from any grant the sender wallet holds and send it to the sender wallet.<br>Claim all available funds when amount = `0`.<br>Rejects if `amount` exceeds allowance or no funds are available |
| claim<br>`0x503914db` | grantId<br>amount| Claim `amount` wei from grant with ID `grantId` and send it to the sender wallet.<br>Claim all available funds when amount = `0`<br>Rejects if `amount` exceeds allowance or no funds are available |
| claimTo<br>`0x1fca9342` | amount<br/>target | Claim `amount` wei from any grant the sender wallet holds and send it to `target`.<br>Claim all available funds when amount = `0`<br>Rejects if `amount` exceeds allowance or no funds are available |
| claimTo<br>`0x30e1198b` | grantId<br>amount<br/>target | Claim `amount` wei from grant with ID `grantId` and send it to `target`.<br>Claim all available funds when amount = `0`<br>Rejects if `amount` exceeds allowance or no funds are available |

### Manager Functions

Manager functions are accessible to managers only (AccessControl role).
Manager wallets are intended to be normal hot wallets of devops / otherwise permissioned users that are allowed to manage grants within specific limits.

| Identifer | Arguments | Description |
|----------|------------|-------------|
| createGrant<br>`0xa92e6f2b` | addr<br>amount<br/>interval<br/>name | Create new grant with allowance of `amount` ETH per `interval` sec and send grant NFT to `addr`.<br>Rejects if allowance exceeds the manager limits |
| lockGrant<br>`0x23d2dad7` | grantId<br>lockTime | Lock grant with ID `grantId` for `lockTime` seconds. This prevents the owner of the grant from claiming funds from the vault for the specified time. |
| removeGrant<br>`0x362e7e13` | grantId | Remove grant with ID `grantId` and burn the NFT that represents the grant. |
| renameGrant<br>`0x362e7e13` | grantId<br/>name | Update name for grant with ID `grantId`. |
| transferGrant<br>`0xdc0533ef` | grantId<br/>target | Transfer the grant NFT that represents the grant with ID `grantId` to `target`.<br>Grantees are expected to transfer the NFT themselves when needed. This is intended to recover stolen NFTs.<br>Rejects if grant allowance exceeds the manager limits.
| updateGrant<br>`0x8a1a14eb` | grantId<br/>amount<br/>interval | Update allowance of grant with ID `grantId` to `amount` ETH per `interval` sec.<br>Rejects if new allowance exceeds the manager limits.

### Owner Functions

Owner functions are accessible to owners only (AccessControl role).
Owner wallets are intended to be high security cold- or multisig wallets.
Owners are allowed to manage grants with no limits and may call recovery functions / upgrade the contract.

| Identifer | Arguments | Description |
|----------|------------|-------------|
| grantRole<br>`0x2f2ff15d` | role<br>account | AccessControl: Grant `role` to `account` |
| revokeRole<br>`0xd547741f` | role<br>account | AccessControl: Revoke `role` from `account` |
| rescueCall<br>`0x96dfe5de` | address<br>amount<br>data | Rescue function, do specified call |
| setPaused<br>`0x16c38b3c` | pause | Disable claiming funds from the contract |
| setClaimTransferLockTime<br>`0xa6a1cb4c` | pause | Set the number of seconds a claim gets locked for when the grant NFT gets transferred |
| setManagerGrantLimits<br>`0x08cf0ebb` | amount<br>interval<br>cooldown<br>cooldownLock | Change manager limits. |
| setProxyManager<br>`0xfe7f3505` | account | Set account that is allowed to upgrade the contract. |

### Grant structure

Grants are stored as a struct within the Vault contract.
There are 4 properties:
* `claimLimit`
Grant "amount" specified via `createGrant`/`updateGrant`. 
Defines the max amount of ETH that can accumulate in the grant over time.
* `claimInterval`
Grant "interval" specified via `createGrant`/`updateGrant`. 
Interval in seconds in which the full amount is available.
* `claimTime`
Specifies the current claim status. The status is stored as unix timestamp.
It basically describes when the full granted amount has been claimed last.
So, if `now() - claimTime` is 0, there are no funds available to claim.
If `now() - claimTime` is >= `claimInterval`, the full granted amount (`claimLimit`) is available to claim.
Each second represents a claimable balance of `claimLimit`/`claimInterval` ETH. 
The value of this property gets increased when claiming funds and may never exceed `now()`
* `dustBalance`
Dust balance that is avaiable to claim. 
This is used to handle claims with a amount smaller than `claimLimit`/`claimInterval` ETH. If a grantee claims ETH worth 1/2 seconds, `claimTime` is increased by 1 and the unclaimed rest (the remaining 1/2 sec) is added to the dustBalance. 

### Claim Process

The `_calculateClaim` function is the central piece of code that does the calculations for claiming funds via a grant. It is designed to handle claims from grantees accurately, while managing the balance available from a grant based on time. 
 
The claim calculation within `_calculateClaim` occurs in several steps:

1. **Initialization and Validation**
   - Retrieve the grant details from the `_grants` map using `grantId`.
   - Check that the grant's properties are valid (non-zero and logical).

2. **Handling Locked Grants**
   - If the grant is locked (`_grantClaimLock[grantId]` > current time), set output variables to reflect no funds can be claimed: `claimAmount` to 0, `newDustBalance` to the current `dustBalance`, and `usedTime` to 0.

3. **Calculating Available Funds**
   - Determine the `availableTime` since the last full claim by subtracting `claimTime` from the current time.
   - If `availableTime` exceeds `claimInterval`, reset it to `claimInterval` to limit the claim to the maximum replenished amount. Also, reset the `dustBalance` to 0 and adjust `baseClaimTime`.

4. **Using Dust Balance**
   - If there is a specific non-zero `requestAmount` and it is less than or equal to the `dustBalance`, use the dust balance to fulfill the claim entirely without affecting the `claimTime`.

5. **Calculating New Claims**
   - If the `requestAmount` is zero (indicating a request for all available funds) or if it exceeds the `dustBalance`, calculate the maximum claimable amount as `(claimLimit * availableTime / grant.claimInterval) + dustBalance`.
   - If a specific `requestAmount` is given and is less than the calculated maximum, adjust the `claimTime` based on the proportion of the `claimLimit` that the request represents. This involves:
     - Calculating the `usedTime` required to fulfill the request (minus the `dustBalance`).
     - Checking for rounding issues and adjusting the `usedTime` if necessary to ensure it accurately represents the funds being claimed.

6. **Setting Return Values**
   - Update `newClaimTime` by adding `usedTime` to `baseClaimTime`.
   - Set `claimAmount` to the lesser of the `requestAmount` or the calculated maximum.
   - Ensure that all calculations maintain logical consistency (e.g., `usedTime` does not exceed `availableTime`).
