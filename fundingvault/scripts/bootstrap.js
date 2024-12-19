const { ethers, network } = require("hardhat");
const { vars } = require("hardhat/config");
const fs = require('fs');


/* bootstrap commands:
npx hardhat vars set FUNDINGVAULT_OWNER_PRIVATE_KEY
npx hardhat vars set FUNDINGVAULT_OWNER_ADDRESS
npx hardhat run scripts/bootstrap.js --network ephemery
*/

async function main() {
    const [deployer] = await ethers.getSigners();
    console.log("deployer address (owner): " + deployer.address)
    console.log("")

    // First deploy proxy
    console.log("deploying FundingVaultProxy...")
    let Proxy = await ethers.getContractFactory("FundingVaultProxy");
    let proxy = await Proxy.connect(deployer).deploy();
    await proxy.waitForDeployment();
    let proxyAddress = await proxy.getAddress();
    console.log("  success: " + proxyAddress);

    // Then deploy token
    console.log("deploying FundingVaultToken...")
    let FundingVaultToken = await ethers.getContractFactory("FundingVaultToken");
    let token = await FundingVaultToken.deploy(proxyAddress);
    await token.waitForDeployment();
    let tokenAddress = await token.getAddress();
    console.log("  success: " + tokenAddress);

    // Lastly, deploy vault implementation
    console.log("deploying FundingVaultV1...")
    let FundingVault = await ethers.getContractFactory("FundingVaultV1");
    let vault = await FundingVault.connect(deployer).deploy();
    await vault.waitForDeployment();
    let vaultAddress = await vault.getAddress();
    console.log("  success: " + vaultAddress);

    // Create instance of vault implementation, attached to the proxy contract
    let proxiedVault = FundingVault.attach(proxyAddress);

    // Initialize vault thorugh proxy's upgradeToAndCall
    console.log("calling upgradeToAndCall on FundingVaultProxy...")
    const vaultConfig = readVaultConfig();
    const initData = vault.interface.encodeFunctionData("initialize(address,uint32,uint128,uint64,uint32,uint32)", [
        tokenAddress,
        vaultConfig.claimTransferLockTime, // in seconds
        vaultConfig.managerLimitAmount, // ETH
        vaultConfig.managerLimitInterval, // in seconds
        vaultConfig.managerLimitCooldown, // in seconds
        vaultConfig.managerLimitCooldownLock // in seconds
    ]);
    console.log("init data: " + initData)
    await proxy.connect(deployer).upgradeToAndCall(vaultAddress, initData, {
        gasLimit: 150000,
    });
    console.log("  success.");

    // change owner if set
    if(vars.has("FUNDINGVAULT_OWNER_ADDRESS")) {
        let ownerAddress = vars.get("FUNDINGVAULT_OWNER_ADDRESS");
        let adminRole = "0x0000000000000000000000000000000000000000000000000000000000000000";
        console.log("changing FundingVault admin to: " + ownerAddress);

        console.log("calling grantRole & setProxyManager on FundingVault...");
        await proxiedVault.connect(deployer).grantRole(adminRole, ownerAddress, {
            gasLimit: 60000,
        });
        await proxiedVault.connect(deployer).setProxyManager(ownerAddress, {
            gasLimit: 40000,
        });
        console.log("  success.");

        console.log("calling revokeRole on FundingVault...");
        await proxiedVault.connect(deployer).revokeRole(adminRole, deployer.address, {
            gasLimit: 60000,
        });
        console.log("  success.");
    }

    const chainId = network.config.chainId;
    saveDeployment(chainId, proxyAddress, tokenAddress, vaultAddress);
}

function readVaultConfig() {
    const path = 'vault.config.json';
    const obj = JSON.parse(fs.readFileSync(path, 'utf8'));
    return obj;
}

function saveDeployment(
    chainId,
    proxyAddr,
    tokenAddr,
    implAddr
) {
    const path = `deployment/${chainId}.json`;
    const obj = {
        "FundingVaultProxy": proxyAddr,
        "FundingVaultToken": tokenAddr,
        "FudingVaultV1": implAddr
    };
    fs.writeFileSync(
        path,
        JSON.stringify(obj),
        "utf8"
    );
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });