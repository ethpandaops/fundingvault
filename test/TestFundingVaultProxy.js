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

		// Create instance of vault implementation, attached to the proxy contract
		let proxiedVault = FundingVault.attach(proxyAddress);

		// Initialize vault thorugh proxy's upgradeToAndCall
		const initData = vault.interface.encodeFunctionData("initialize(address)", [tokenAddress]);
		await proxy.connect(owner).upgradeToAndCall(vaultAddress, initData);
		return {proxy, token, vault, proxiedVault, owner};
	}

	describe("Deployment", function () {
		it("Token vault should be proxy address", async function () {
			const {proxy, token, vault, proxiedVault, owner} = await loadFixture(deployProxyFixture);
			expect(await token.getVault()).to.equal(await proxy.getAddress());
		});
		it("Proxy implementation should be vault impl address", async function () {
			const {proxy, token, vault, proxiedVault, owner} = await loadFixture(deployProxyFixture);
			expect(await proxy.implementation()).to.equal(await vault.getAddress());
		});
		it("Vault implementation's token should be token address", async function () {
			const {proxy, token, vault, proxiedVault, owner} = await loadFixture(deployProxyFixture);
			expect(await proxiedVault.getVaultToken()).to.equal(await token.getAddress());
		});
	});

	describe("Grant Management", function () {
		it("Create Grant", async function () {
			const {proxy, token, vault, proxiedVault} = await loadFixture(deployProxyFixture);
			const [owner, grantee] = await ethers.getSigners();

			// add manager role to owner
			let managerRole = await proxiedVault.GRANT_MANAGER_ROLE();
			await proxiedVault.grantRole(managerRole, owner.getAddress());

			// diable grant locking after creation / transfer for easier test
			// otherwise, we'd have to wait 10mins after grant creation to see some claimable balance
			await proxiedVault.setClaimTransferLockTime(0);

			// create grant (1000 ETH per hour)
			await proxiedVault.createGrant(grantee.getAddress(), 1000, 3600);

			// check if grant token has been sent to grantee
			expect(await token.balanceOf(grantee.getAddress())).to.equal(1);

			// check claimable balance as grantee
			expect(await proxiedVault.connect(grantee).getClaimableBalance()).to.equal(1000000000000000000000n); // 1000 ETH should be claimable
		});
	});
	describe("Request funds", function () {
		it("Request max available balance", async function () {
			const {proxy, token, vault, proxiedVault} = await loadFixture(deployProxyFixture);
			const [owner, grantee] = await ethers.getSigners();

			// add manager role to owner
			let managerRole = await proxiedVault.GRANT_MANAGER_ROLE();
			await proxiedVault.grantRole(managerRole, owner.getAddress());

			// diable grant locking after creation / transfer for easier test
			// otherwise, we'd have to wait 10mins after grant creation to see some claimable balance
			await proxiedVault.setClaimTransferLockTime(0);

			// send some funds to the vault, otherwise there is nothing to claim
			await owner.sendTransaction({
				to: proxiedVault.getAddress(),
				value: 2000000000000000000000n, // send 2000 ETH
			});

			// create grant (1000 ETH per hour)
			await proxiedVault.createGrant(grantee.getAddress(), 1000, 3600);

			// claim all available balance
			let oldGranteeBalance = await ethers.provider.getBalance(grantee.getAddress());
			let claimTx = await proxiedVault.connect(grantee).claim(0);

			// await receipt
			let claimTxReceipt = await claimTx.wait();
			let claimTxFees = claimTxReceipt.gasUsed * claimTxReceipt.gasPrice;

			// check balance increase
			let newGranteeBalance = await ethers.provider.getBalance(grantee.getAddress());
			expect(newGranteeBalance - oldGranteeBalance).to.equal(1000000000000000000000n - claimTxFees); // 1000 ETH increase

			// check claimable balance as grantee
			expect(await proxiedVault.connect(grantee).getClaimableBalance()).to.equal(0); // 0 ETH should be claimable
		});
	});
});
  