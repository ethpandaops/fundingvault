// const { ethers, network } = require("hardhat");
const { vars , task } = require("hardhat/config");
const FundingVaultABI = require("../FundingVaultABI.json");

/* run this command before running any tasks below:
npx hardhat vars set FUNDINGVAULT_MANAGER_PRIVATE_KEY
*/

task("create-grant", "Creates a FundingVault Grant")
  .addParam("grantee", "The address of the Grantee")
  .addParam("amount", "The max amount in ETH granted within the defined interval")
  .addParam("interval", "Duration in seconds")
  .addParam("name", "The name of the Grant")
  .setAction(async (args) => {
    const managerWallet = getManagerWallet();
    const vaultContract = getVaultInstance();

    const txReceipt = await vaultContract.connect(managerWallet).createGrant(
        args.grantee,
        args.amount,
        args.interval,  
        args.name
    );

    console.log("Grant created at tx: ", txReceipt.hash);
  })

/// HELPER FUNCTIONS

function getVaultInstance() {
    // TEMP:npx hardhat vars set FUNDINGVAULT_ADDRESS
    let vaultAddress = vars.get("FUNDINGVAULT_ADDRESS");
    return new ethers.Contract(vaultAddress, FundingVaultABI, ethers.provider);
}

function getManagerWallet() {
    let managerKey = vars.get('FUNDINGVAULT_MANAGER_PRIVATE_KEY');
    let manageWallet = new ethers.Wallet(managerKey, ethers.provider);
    return manageWallet;
}