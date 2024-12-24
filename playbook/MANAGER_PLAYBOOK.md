# Grant Manager Playbook

Account holders who are granted with the `GRANT_MANAGER_ROLE` should use this playbook as a reference.

Grant managers may perform the following actions:
- Create grants
- Update the amount per time period entitled to an existing grant recipient
- Rename grants
- Transfer grants from one recipient to another
- Delete grants

The next section covers the actions listed above more in-depth.

Grant managers can perform all actions above either via the website or CLI.

This playbook provides the tutorial on how you can manage grants using the UI provided by the website. To use the CLI, you may refer to the [FundingVault CLI Guide](../fundingvault/README.md).

---
## FundingVault Website Tutorial

### Wallet Configuration

When landed on the website, you must first connect to your wallet.

Make sure that you are connected to the intended network and account. Once you have done so, you should see the homescreen with the "Manage Grants" button as shown below:

![manager-homepage](./assets/managers/manager-homepage.png)

Click on "Manage Grants" to see the Grants List.

![manager-grants-list-0](./assets/managers/manager-grants-list-0.png)

The Grants List is currently empty, so let's go ahead and create one.

### Create Grant(s)

Click on "New Grant", the following prompt should appear:

![manager-create-grant-prompt](./assets/managers/manager-create-grant-prompt.png)

> **⚠️ IMPORTANT**:
>
> Grant managers are limited to a specified amount that they can manage within a configured interval. 
>
> Currently, defaults at 100K ETH per month for each manager. This is subject to change by the Vault owner.
>
> If a grant manager has created (or updated) grants with the total amount exceeded the limit, the grant manager may need to "cooldown" before they can continue managing more grants.
>
> See [Cooldown Deep-Dive](#cooldown-deep-dive) section for more details.
>

After filling up the form, clicking on "Create Grant" should prompt your wallet to submit the transaction. A transaction hash will appear after submitting the transaction.

![manager-create-grant-submission](./assets/managers/manager-create-grant-submission.png)

Close the prompt window, you should see the newly added grant in the list.

![manager-grants-list-1](./assets/managers/manager-grants-list-1.png)

>
> **ℹ️ NOTE**:
>
> The Grants List shows ALL grants that are being actively managed and is NOT scoped to individual managers.
> In other words, grants created by other manager(s) are also visible to all managers.
>

>
> **ℹ️ NOTE**:
>
> Newly created grants will not be immediately claimable, because of a "transfer lock" put in place.  
> Transfer lock is a delay measure to avoid funds from being drained due to human errors, such as:
>
> (1) setting the wrong recipient address, 
>
> (2) setting the wrong amount, 
>
> etc.
>
> The current default transfer lock time is 10 minutes, and is subject to change by the Vault owner.
>

After the transfer lock period has passed, the user should be able to claim the entitled amount.

![claimer-grant-eligible](./assets/claimers/claimer-grant-eligible.png)

>
> **✅ Knowledge Check**:
> 
> The "allowance" set for a grant is limited by fixed amount of tokens that the recipient(s) can claim in a per interval basis, e.g. 100k ETH per month.
>
> The actual "claimable" amount is linearly proportional to the number of seconds from current time to a `claimTime`, a recorded timestamp when the recipient had previously interacted with the grant.
>
> For example, Alice claimed the full amount (60 ETH) of her grant, which entitled her to 60 ETH per hour. Ten minutes later, Alice now has a claimable balance of 10 ETH, because that is the amount that she is supposedly entitled to based on the duration from her previous interaction.
>
> Two things to note from the example above:
>
> - If Alice did not make a claim when her claimable balance has reached the maximum amount (60 ETH), the claimable balance would plateau and she will "miss out" on funds that would've been claimable. In other words, any inactivity with her grant beyond an hour would still entitle her to 60 ETH.
>
> - The full amount would be available to Alice upon creation (plus duration of transfer lock) of the grant. This is because the `claimTime` is recorded as the `creationTimestamp - interval`.
>
> After a grant is created, the NFT representing the grant is minted and transferred to the recipient. Any account holding the NFT would be eligible to claim funds for the grant. Minting and transferring of NFT would trigger a transfer lock, although claimable funds would still accumulate during the lock period.
>

### Update Grant(s)

Managers can modify an existing grant, which allows:

- Changing the amount per interval entitled to the recipient 

>
> **⚠️ IMPORTANT**:
>
> - If the new `amount / interval` is higher than previous, this increases the cooldown timer.
>
> - Changing `amount / interval` may also affect the total unclaimed amount of previously claimable balance. E.g. if a recipient had accumulated X claimable amount and the `amount / interval` for the grant had doubled, the recipient would instead be entitled to 2X claimable amount after the update is in effect.
>

- Changing the grant name

Continuing from the example above, let's increase the amount and change the grant name.

Click on the ✏ pencil icon under "Granted Amount", and the following prompt should appear:

![manager-update-grant-prompt](./assets/managers/manager-update-grant-prompt.png)

After submitting the transaction, you should see the updated amount in "Granted Amount".

![manager-grants-list-2](./assets/managers/manager-grants-list-2.png)

Similarly, you can change the grant name by clicking on the pencil icon under "Name".

![manager-rename-grant-prompt](./assets/managers/manager-rename-grant-prompt.png)

![manager-grants-list-3](./assets/managers/manager-grants-list-3.png)

### Lock Grant(s)

Manager(s) can suspend an existing grant by clicking on the 🔒 pad lock icon. This would set the claimable balance for the grant to zero.

Managers can specify the lock time in seconds in the prompt shown below:

![manager-lock-grant-prompt](./assets/managers/manager-lock-grant-prompt.png)

The user shouldn't be able to claim the funds for 10 minutes.

![claimer-grant-locked-1](./assets/claimers/claimer-grant-locked-1.png)

### Transfer Grant(s)

An existing grant can either be transferred by:

(1) The recipient themselves, OR

(2) Grant managers (**⚠️ IMPORTANT**: Increases the cooldown timer)

> **ℹ️ Note**: A "transfer lock" will immediately be in effect after transferring of grants. 

A user can transfer grants by simply sending the NFT to a target address.

For grant managers, simply click on the arrow ↗ icon to initiate the transfer action.

Simply provide the target address in the prompt shown below:

![manager-transfer-grant-prompt](./assets/managers/manager-transfer-grant-prompt.png)

### Delete Grant(s)

Grant managers can delete (which in turn, burns the NFT representing) grants, and thereby permanently disqualifies the recipient from claiming funds from the Vault. 

To delete a grant, click on the trash 🗑️ icon. Then, click "Delete Grant" to submit the transaction.

![manager-delete-grant-prompt](./assets/managers/manager-delete-grant-prompt.png)

---

## Cooldown Deep-Dive

TODO