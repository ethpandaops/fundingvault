const {	loadFixture } = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { expect } = require("chai");

describe("Funding Vault Tests", function () {
	// We define a fixture to reuse the same setup in every test.
	// We use loadFixture to run this setup once, snapshot that state,
	// and reset Hardhat Network to that snapshot in every test.
	async function deployProxyFixture() {
		const [owner] = await ethers.getSigners();
		// First deploy proxy
		let Proxy = await ethers.getContractFactory("FundingVaultProxy");
		let proxy = await Proxy.connect(owner).deploy();
		let proxyAddress = await proxy.getAddress();

		// Then deploy token
		let FundingVaultToken = await ethers.getContractFactory("FundingVaultToken");
		let token = await FundingVaultToken.deploy(proxyAddress);
		let tokenAddress = await token.getAddress();

		// Lastly, deploy vault implementation
		let FundingVault = await ethers.getContractFactory("FundingVaultV1");
		let vault = await FundingVault.connect(owner).deploy();
		let vaultAddress = await vault.getAddress();

		// Initialize vault thorugh proxy's upgradeToAndCall
		const initData = vault.interface.encodeFunctionData("initialize(address)", [tokenAddress]);
		await proxy.connect(owner).upgradeToAndCall(vaultAddress, initData);
		return {proxy, token, vault, owner};
	}

	describe("Deployment", function () {
		it("Token vault should be proxy address", async function () {
			const {proxy, token, vault, owner} = await loadFixture(deployProxyFixture);
			expect(await token.getVault()).to.equal(await proxy.getAddress());
		});
		it("Proxy implementation should be vault impl address", async function () {
			const {proxy, token, vault, owner} = await loadFixture(deployProxyFixture);
			expect(await proxy.implementation()).to.equal(await vault.getAddress());
		});
		it("Vault implementation's token should be token address", async function () {
			const {proxy, token, vault, owner} = await loadFixture(deployProxyFixture);
			const initData = vault.interface.encodeFunctionData("getVaultToken()", []);
			const tokenAddress = await token.getAddress();
			const returnData = await owner.call({data: initData, to: await proxy.getAddress()});
			const stripZeroes = returnData.replace(/0x0+/, "0x");
			expect(tokenAddress.toLowerCase()).to.equal(stripZeroes);
		});
	});
});
  