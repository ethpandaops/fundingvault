const { ethers, network } = require("hardhat");
const fs = require('fs');

async function main() {
    const [owner] = await ethers.getSigners();
    console.log("deployer address (owner): " + owner.address)
    console.log("")

    // First deploy proxy
    console.log("deploying FundingVaultProxy...")
    let Proxy = await ethers.getContractFactory("FundingVaultProxy");
    let proxy = await Proxy.connect(owner).deploy();
    let proxyAddress = await proxy.getAddress();
    console.log("  success: " + proxyAddress);

    // Then deploy token
    console.log("deploying FundingVaultToken...")
    let FundingVaultToken = await ethers.getContractFactory("FundingVaultToken");
    let token = await FundingVaultToken.deploy(proxyAddress);
    let tokenAddress = await token.getAddress();
    console.log("  success: " + tokenAddress);

    // Lastly, deploy vault implementation
    console.log("deploying FundingVaultV1...")
    let FundingVault = await ethers.getContractFactory("FundingVaultV1");
    let vault = await FundingVault.connect(owner).deploy();
    let vaultAddress = await vault.getAddress();
    console.log("  success: " + vaultAddress);

    // Create instance of vault implementation, attached to the proxy contract
    let proxiedVault = FundingVault.attach(proxyAddress);

    // Initialize vault thorugh proxy's upgradeToAndCall
    console.log("calling upgradeToAndCall on FundingVaultProxy...")
    const initData = vault.interface.encodeFunctionData("initialize(address)", [tokenAddress]);
    await proxy.connect(owner).upgradeToAndCall(vaultAddress, initData, {
        gasLimit: 100000,
    });
    console.log("  success.");
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });