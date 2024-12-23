# Admin Playbook

Account holders who are granted with the `DEFAULT_ADMIN_ROLE` should use this playbook as a reference to learn about (1) the deployment of the FundingVault proxy, implementation and token contracts, and (2) all available operations that can be executed by admins.

Currently, all admin operations can only be done in the CLI.

You may refer to the [FundingVault CLI Guide](../fundingvault/README.md) to see details about the available commands for all admin operations.

## Step 1: Contract Deployment

> TL;DR, before contract deployment, make sure:
> 
> - `FUNDINGVAULT_OWNER_PRIVATE_KEY` variable has been configured
> - Optional: `FUNDINGVAULT_OWNER_ADDRESS` variable has been configured
> - `vault.config.json` should provide the expected initial configuration for the vault.

By default, the deployer will be the sole admin of FundingVault. In other words, it can do the following:
- Grant other accounts with `DEFAULT_ADMIN_ROLE`
- Assign other accounts as the proxy manager.

The deployer address is derived from the `FUNDINGVAULT_OWNER_PRIVATE_KEY` variable.

If the deployer intends to transfer ownership on deployment (a.k.a. renouncing DEFAULT_ADMIN_ROLE and transferring proxy manager role) to a different address, they should configure the `FUNDINGVAULT_OWNER_ADDRESS` variable.

Lastly before running `bootstrap.js`, make sure you make the necessary changes to `vault.config.json` for your initialization configuration upon deployment of the Funding Vault contract. Otherwise, the default values will be used.

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

## Step 2: Role Assignments

Admins can grant and revoke accounts for:
- `DEFAULT_ADMIN_ROLE` : role with the highest privilege
- `GRANT_MANAGER_ROLE` : role responsible for creating, modifying and deleting grants. Bounded by certain configurations described in details in the next section.

Admins may also designate an account as a proxy manager, which is responsible for performing upgrades.

## Step 3: Global Configurations

Admins are responsible for configuring "limits" that are put in place for grant managers:

### Claim Transfer Lock Time

Grants that are created or transferred to other recipients are immediately subjected to a lock time. The amount of funds entitled to the recipient is still accumulating within the lock time period but will not be claimable until the lock is lifted.

### Manager amount and interval

The amount of funds can be managed within a time period.

### Manager Cooldown

Admins can specify a "cooldown" in seconds.

The actual number of seconds added to the cooldown clock is linearly proportional to the (change in amount / interval) vs (manager's limit / interval).

The following action(s) could increase the manager's cooldown clock:

- Creating a new grant
- Increasing the amount for an existing grant

For example, given Alice (the manager) has a limit of 100k ETH per 2592000 seconds (about a month) for grants, with a cooldown of 86400 seconds (24 hours):

- Alice creates a grant for Bob with 25k ETH per month, increases 21600 seconds to her cooldown clock
- Alice increases Bob's grant, which entitles him to 70k ETH per month (+50k ETH), which increases another 43200 seconds to her cooldown clock.
- Deleting grants and decreasing grant amount will not change Alice's cooldown clock.

### Manager Cooldown Lock

The threshold to lock managers based on their cooldown clock.

If `managerCooldownClock - now() >= cooldownLock`, it temporarily locks the manager and prevents them from creating, updating and transferring grants, until the lock period is over.

Using the same example as above, but now we set a cooldown lock of 43200 seconds (12 hours).

- Alice creates a grant with the full 100k ETH/month for Charlie, which increases a cooldown of 24 hours.

- Alice can no longer create/update/transfer grants for 12 hours.

- 12 hours later, even though Alice still has to "cooldown" for 12 more hours, she may now perform any actions that are allowable. However, the locking time now gets applied linearly following the rules above.