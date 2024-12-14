const { task } = require("hardhat/config");
const {
  getVaultInstance,
  getManagerWallet
} = require("./common");

/* run this command before running any tasks below:
npx hardhat vars set FUNDINGVAULT_MANAGER_PRIVATE_KEY
*/

task("create-grant", "Creates a FundingVault Grant")
  .addParam("grantee", "The address of the Grantee")
  .addParam("amount", "The max allowance in ETH can be claimed within the defined interval")
  .addParam("interval", "Duration in seconds")
  .addParam("name", "The name of the Grant")
  .setAction(async (args) => {
    const managerWallet = getManagerWallet();
    const vaultContract = getVaultInstance();
    
    // pads zeroes to the right if shorter than 32 bytes
    const nameBytes32 = ethers.zeroPadBytes(
      ethers.toUtf8Bytes(args.name),
      32
    );

    const txReceipt = await vaultContract.connect(managerWallet).createGrant(
        args.grantee,
        args.amount,
        args.interval,  
        nameBytes32
    );

    // TODO: get ID from event to console log

    console.log("Grant created at tx: ", txReceipt.hash);
  })

task("lock-grant", "Locks grant with ID for a duration in seconds")
  .addParam("id", "Grant ID")
  .addParam("time", "Lock time in seconds")
  .setAction(async (args) => {
    const managerWallet = getManagerWallet();
    const vaultContract = getVaultInstance();

    const txReceipt = await vaultContract.connect(managerWallet).lockGrant(
        args.id,
        args.time
    );

    console.log(`Grant (ID: ${args.id}) locked at tx: `, txReceipt.hash);
  })

task("remove-grant", "Removes a grant with ID and burn the NFT that represents the grant")
  .addParam("id", "Grant ID")
  .setAction(async (args) => {
    const managerWallet = getManagerWallet();
    const vaultContract = getVaultInstance();

    const txReceipt = await vaultContract.connect(managerWallet).removeGrant(
        args.id,
    );

    console.log(`Grant (ID: ${args.id}) removed at tx: `, txReceipt.hash);
  })

task("rename-grant", "Update name for grant with ID")
  .addParam("id", "Grant ID")
  .addParam("name", "The name of the Grant")
  .setAction(async (args) => {
    const managerWallet = getManagerWallet();
    const vaultContract = getVaultInstance();

    // pads zeroes to the right if shorter than 32 bytes
    const nameBytes32 = ethers.zeroPadBytes(
      ethers.toUtf8Bytes(args.name),
      32
    );

    const txReceipt = await vaultContract.connect(managerWallet).renameGrant(
        args.id,
        nameBytes32
    );

    console.log(`Grant (ID: ${args.id}) removed at tx: `, txReceipt.hash);
  })

task("transfer-grant", "Transfer the grant NFT that represents the grant with ID grantId to target")
  .addParam("id", "Grant ID")
  .addParam("target", "The target recipient of the Grant to be transferred")
  .setAction(async (args) => {
    const managerWallet = getManagerWallet();
    const vaultContract = getVaultInstance();

    const txReceipt = await vaultContract.connect(managerWallet).transferGrant(
        args.id,
        args.target
    );

    console.log(`Grant (ID: ${args.id}) transferred to ${args.target} at tx: `, txReceipt.hash);
  })

task("update-grant", "Update allowance of grant with ID to a specified amount in ETH per interval in seconds")
  .addParam("id", "Grant ID")
  .addParam("amount", "The max allowance in ETH can be claimed within the defined interval")
  .addParam("interval", "Duration in seconds")
  .setAction(async (args) => {
    const managerWallet = getManagerWallet();
    const vaultContract = getVaultInstance();

    const txReceipt = await vaultContract.connect(managerWallet).updateGrant(
        args.id,
        args.amount,
        args.interval
    );

    console.log(`Grant (ID: ${args.id}) updated at tx: `, txReceipt.hash);
  })