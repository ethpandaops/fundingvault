const { vars } = require("hardhat/config");
const FundingVaultABI = require("../../FundingVaultABI.json");
const fs = require('fs');

function getVaultInstance() {
    const chainId = network.config.chainId;
    const deploymentJsonStr = fs.readFileSync(
      `deployment/${chainId}.json`,
      'utf8'
    );
    const deployment = JSON.parse(deploymentJsonStr);
    return new ethers.Contract(deployment.FundingVaultProxy, FundingVaultABI, ethers.provider);
}

function getManagerWallet() {
    let managerKey = vars.get('FUNDINGVAULT_MANAGER_PRIVATE_KEY');
    let manageWallet = new ethers.Wallet(managerKey, ethers.provider);
    return manageWallet;
}

function getOwnerWallet() {
    let ownerKey = vars.get('FUNDINGVAULT_OWNER_PRIVATE_KEY');
    let ownerWallet = new ethers.Wallet(ownerKey, ethers.provider);
    return ownerWallet;
}

module.exports = {
    getVaultInstance: getVaultInstance,
    getManagerWallet: getManagerWallet,
    getOwnerWallet: getOwnerWallet
}