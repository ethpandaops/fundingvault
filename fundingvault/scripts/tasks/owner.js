const { task } = require("hardhat/config");
const {
  getVaultInstance,
  getOwnerWallet
} = require("./common");

/* run this command before running any tasks below:
npx hardhat vars set FUNDINGVAULT_OWNER_PRIVATE_KEY
*/

task("grant-role", "Grant specified role to account")
  .addFlag("managerRole", "Provide this flag to grant manager role")
  .addFlag("ownerRole", "Provide this flag to grant owner role")
  .addParam("account")
  .setAction(async (args) => {
    const ownerWallet = getOwnerWallet();
    const vaultContract = getVaultInstance();
    let role = ethers.ZeroHash;

    if (args.managerRole == args.ownerRole) {
        throw new Error("Must provide either manager-role or owner-role only");
    } else if (args.managerRole) {
        role = await vaultContract.GRANT_MANAGER_ROLE.staticcall();
    }

    const txReceipt = await vaultContract.connect(ownerWallet).grantRole(role, args.account);

    console.log(`Granted ${args.account} with role ${role}. Tx: ${txReceipt.hash}`);
  })

task("revoke-role", "Revoke specified role to account")
  .addFlag("managerRole", "Provide this flag to grant manager role")
  .addFlag("ownerRole", "Provide this flag to grant owner role")
  .addParam("account")
  .setAction(async (args) => {
    const ownerWallet = getOwnerWallet();
    const vaultContract = getVaultInstance();
    let role = ethers.ZeroHash;

    if (args.managerRole == args.ownerRole) {
        throw new Error("Must provide either manager-role or owner-role only");
    } else if (args.managerRole) {
        role = await vaultContract.GRANT_MANAGER_ROLE.staticcall();
    }

    const txReceipt = await vaultContract.connect(ownerWallet).revokeRole(role, args.account);

    console.log(`Revoked ${args.account} with role ${role}. Tx: ${txReceipt.hash}`);
  })

task('set-paused', "Enable/Disable claiming funds from the Vault")
  .addFlag("pause", "Disable fund claims")
  .addFlag("unpause", "Resume fund claims")
  .setAction(async (args) => {
    const ownerWallet = getOwnerWallet();
    const vaultContract = getVaultInstance();

    if (args.pause == args.unpause) {
        throw new Error("Must provide either pause or unpause flags only");
    }

    const paused = args.pause ? args.pause : args.unpause;

    const txReceipt = await vaultContract.connect(ownerWallet).setPaused(paused);

    console.log(`Vault is paused: ${paused}. Tx: ${txReceipt.hash}`);
  })

task('set-claim-lock', "Set the number of seconds a claim gets locked for when the grant NFT gets transferred")
  .addParam("lockTime")
  .setAction(async (args) => {
    const ownerWallet = getOwnerWallet();
    const vaultContract = getVaultInstance();

    const txReceipt = await vaultContract.connect(ownerWallet).setClaimTransferLockTime(args.lockTime);

    console.log(`Claim transfer lock time has been updated. Tx: ${txReceipt.hash}`);
  })

task('set-manager-limits', "Change the manager's grant limit in ETH per interval in seconds")
  .addParam("amount", "The amount in ETH within the defined interval")
  .addParam("interval", "Interval in seconds")
  .addParam("cooldown", "Number of seconds added to the cooldown clock")
  .addParam("cooldownLock", "The manager is locked if the cooldown clock is above this value")
  .setAction(async (args) => {
    const ownerWallet = getOwnerWallet();
    const vaultContract = getVaultInstance();

    const txReceipt = await vaultContract.connect(ownerWallet).setManagerGrantLimits(
        args.amount,
        args.interval,
        args.cooldown,
        args.cooldownLock
    );

    console.log(`Tx: ${txReceipt.hash}`);
  })
  
task('set-proxy-manager', "Set account that is allowed to upgrade the contract")
  .addParam("account")
  .setAction(async (args) => {
    const ownerWallet = getOwnerWallet();
    const vaultContract = getVaultInstance();

    const txReceipt = await vaultContract.connect(ownerWallet).setProxyManager(
        args.account
    );

    console.log(`Tx: ${txReceipt.hash}`);
  })
