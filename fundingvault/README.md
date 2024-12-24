# FundingVault CLI Guide

## Prerequisite:

Make sure you are on the current directly and have installed the dependencies, before proceeding:

Run `$ pwd` to make sure you are on the `fundingvault/fundingvault` directory.

```bash
npm install
```

Before deploying the FundingVault contracts to a network, please make sure that the network configuration has been added to `hardhat.config.js`.

You may refer to the [Technical Concept](./docs/TechnicalConcept.md) documentation to learm more about the available operations that can be performed on the Vault. 

---

# Deployment

Go to [`deployment`](./deployment/) to view all deployment addresses.

|  | Network | Address |
| --- | --- | --- |
| `FundingVaultProxy` | Automata Testnet | [0x74bBc82C68fc1e83BFA98cb9A2d6ef8241F46d28](https://explorer-testnet.ata.network/address/0x74bBc82C68fc1e83BFA98cb9A2d6ef8241F46d28) |
| | Automata Mainnet | [0x92B59abfE96C9E4bEe808476dD0975a9b89b6c45](https://explorer.ata.network/address/0x92B59abfE96C9E4bEe808476dD0975a9b89b6c45) |
| `FundingVaultToken` | Automata Testnet | [0x26540FCfd36262fbfb49Aa4eC6108B20595b796a](https://explorer-testnet.ata.network/address/0x26540FCfd36262fbfb49Aa4eC6108B20595b796a) |
| | Automata Mainnet | [0x8618395693B1e4BE231d29E338b2a1f9e645e56F](https://explorer.ata.network/address/0x8618395693B1e4BE231d29E338b2a1f9e645e56F) |
| `FudingVaultV1` | Automata Testnet | [0x60607656b0DBc91e038587967E4a5e20B0e9809F](https://explorer-testnet.ata.network/address/0x60607656b0DBc91e038587967E4a5e20B0e9809F) |
| | Automata Mainnet | [0x0E3Fe1a63E70aF7D4137172D8527912B5E11E02c](https://explorer.ata.network/address/0x0E3Fe1a63E70aF7D4137172D8527912B5E11E02c) |

---

# Running the CLI

>
> ℹ️ **Note**: Possible `network_name` is either `testnet` or `mainnet`.
>

## Owner Operations

Before perfoming any actions below, make sure to configure your keys by running:
```bash
npx hardhat vars set FUNDINGVAULT_OWNER_PRIVATE_KEY
```

Owners may perform the following operations:
- Contract Deployment
- Granting and revoking: 
    - `DEFAULT_ADMIN_ROLE` (role with the highest privilege)
    - `GRANT_MANAGER_ROLE` (accounts granted with this role, may create and/or modify grants)
- Designate a proxy manager, to perform contract upgrades.
- Pause and unpause the vault contract.
- Change the duration of lock time after a grant is created or transferred.
- Change managers' limit on funds to be managed within a defined interval.
    - Cooldown period on managers who exceeded their quota.
- Perform rescue calls on a specified function.

### Contract Deployment

Before running the deployment script, make sure you make the necessary changes to `vault.config.json` for your initialization configuration upon deployment of the Funding Vault contract. Otherwise, the default values will be used.

Default Values:
```javascript
{   
    // (seconds) 10 minutes of lock time applied after a grant is created or transferred
    "claimTransferLockTime": 600,
    // (ETH) max of 100k ETH can be managed within
    "managerLimitAmount": 100000,
    // (seconds) per ~month
    "managerLimitInterval": 2592000,
    // (seconds) 24 hours increment to cooldownClock
    "managerLimitCooldown": 86400,
    // (seconds) 12 hours threshold
    // Managers cannot create/update/transfer grants 
    // if cooldownClock - now() >= managerLimitCooldownLock
    "managerLimitCooldownLock": 43200
}
```

If you intend to transfer owner of the Vault to a different address on deployment, you must set the `FUNDINGVAULT_OWNER_ADDRESS` variable before running `bootstrap.js`.

```bash
npx hardhat vars set FUNDINGVAULT_OWNER_ADDRESS
```

The command above ensures that the deployer will not retain `DEFAULT_ADMIN_ROLE` and all admin privileges are to transfer to `FUNDINGVAULT_OWNER_ADDRESS`.

```bash
npx hardhat run scripts/bootstrap.js --network <network-name>
```

The deployment addresses will be written and stored at `/deployment/{chainId}.json`.

After deploying the contracts, owners must ensure that `FundingVaultProxy` is funded. This can be done by simply sending tokens directly to the contract. Or by running the command below:

```bash
npx hardhat fund-vault --network <network-name> --amount <ETH>
```

### Grant Roles

To grant `account` with `DEFAULT_ADMIN_ROLE`:

```bash
npx hardhat grant-role --network <network-name> --account <account> --owner-role
```

To grant `account` with `GRANT_MANAGER_ROLE`

```bash
npx hardhat grant-role --network <network-name> --account <account> --manager-role
```

### Revoke Roles

To revoke `account` with `DEFAULT_ADMIN_ROLE`:

```bash
npx hardhat revoke-role --network <network-name> --account <account> --owner-role
```

To revoke `account` with `GRANT_MANAGER_ROLE`

```bash
npx hardhat revoke-role --network <network-name> --account <account> --manager-role
```

### Set proxy manager:

```bash
npx hardhat set-proxy-manager --network <network-name> --account <account>
```

> ℹ️ **NOTE**:
> 
> You may pass an optional `--skip-admin-check` flag to this command.
> Skipping this check however, may result in the new proxy manager to NOT have a `DEFAULT_ADMIN_ROLE`,
> which is required for `set-proxy-manager`.
>

### Pause/Unpause Vault

To pause the Vault:

```bash
npx hardhat set-paused --network <network-name> --pause
```

To unpause the Vault:

```bash
npx hardhat set-paused --network <network-name> --unpause
```

### Set claim lock:

```bash
npx hardhat set-claim-lock --network <network-name> --lock-time <seconds>
```

### Set manager limit:

```bash
npx hardhat set-manager-limits --network <network-name> --amount <ETH> --interval <seconds> --cooldown <seconds> --cooldown-lock <seconds>
```

### Rescue calls:

```bash
npx hardhat rescue-call --network <network-name> --addr <address> --amount <wei> --data <calldata-hexstring>
```

Omit the `--data` argument to perform emergency fund withdrawals.

---

## Grant Manager Operations

Before perfoming any actions below, make sure to configure your keys by running:
```bash
npx hardhat vars set FUNDINGVAULT_MANAGER_PRIVATE_KEY
```

Grant Managers may `create`, `lock`, `remove`, `rename`, `transfer` and `update` grants.

### Create Grant

Grants can be created, with an allocated ETH per defined interval for the specified account. (as long as it does not exceed manager limit, otherwise may trigger cooldown)

```bash
npx hardhat create-grant --network <network-name> --grantee <address> --amount <ETH> --interval <seconds> --name <string>
```

### Lock Grant

Lock grant with ID for a specified duration (in seconds).

```bash
npx hardhat lock-grant --network <network-name> --id <number> --time <seconds>
```

### Remove Grant

```bash
npx hardhat remove-grant --network <network-name> --id <number>
```

### Rename Grant

```bash
npx hardhat rename-grant --network <network-name> --id <number> --name <string>
```

### Transfer Grant

Transfers the grant with ID to a specified target.

```bash
npx hardhat transfer-grant --network <network-name> --id <number> --target <address>
```

### Update Grant

Update allowance of grant with ID to a specified amount in ETH per interval. (increased allowance may trigger a cooldown)

```bash
npx hardhat update-grant --network <network-name> --id <number> --amount <ETH> --interval <seconds>
```

---

## User (Grantee) Operations

Before perfoming any actions below, make sure to configure your keys by running:
```bash
npx hardhat vars set FUNDINGVAULT_USER_PRIVATE_KEY
```

Grantees can either make claims on all grants that they are entitled for, or they may specify an ID to only get funds from that particular grant.

Grantees may pass 0 to --amount to claim all available funds from grant(s).

Grantees can also sends funds to a specified target.

```bash
npx hardhat claim --network <network-name> --amount <wei> --id <OPTIONAL: number> --target <OPTIONAL: address>
```