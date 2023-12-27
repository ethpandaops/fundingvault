const {	loadFixture } = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { expect } = require("chai");
// Check if HARDHAT_DEBUG is set. If set, debug prints will be enabled.
// Usage: HARDHAT_DEBUG=1 npx hardhat test to enable debug prints,
// omit the env var to disable debug prints.
const debug = process.env.HARDHAT_DEBUG !== undefined;

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
			const ownerAddress = await owner.getAddress();
			const granteeAddress = await grantee.getAddress();

			// add manager role to owner
			let managerRole = await proxiedVault.GRANT_MANAGER_ROLE();
			await proxiedVault.grantRole(managerRole, ownerAddress);

			// diable grant locking after creation / transfer for easier test
			// otherwise, we'd have to wait 10mins after grant creation to see some claimable balance
			await proxiedVault.setClaimTransferLockTime(0);

			// create grant (1000 ETH per hour)
			await proxiedVault.createGrant(granteeAddress, 1000, 3600);

			// check if grant token has been sent to grantee
			expect(await token.balanceOf(granteeAddress)).to.equal(1);

			// check claimable balance as grantee
			expect(await proxiedVault.connect(grantee).getClaimableBalance()).to.equal(1000000000000000000000n); // 1000 ETH should be claimable
		});
	});
	describe("Request funds", function () {
		it("Request max available balance", async function () {
			const {proxy, token, vault, proxiedVault} = await loadFixture(deployProxyFixture);
			const [owner, grantee] = await ethers.getSigners();
			const ownerAddress = await owner.getAddress();
			const granteeAddress = await grantee.getAddress();

			// add manager role to owner
			let managerRole = await proxiedVault.GRANT_MANAGER_ROLE();
			await proxiedVault.grantRole(managerRole, ownerAddress);

			// diable grant locking after creation / transfer for easier test
			// otherwise, we'd have to wait 10mins after grant creation to see some claimable balance
			await proxiedVault.setClaimTransferLockTime(0);

			// send some funds to the vault, otherwise there is nothing to claim
			await owner.sendTransaction({
				to: await proxiedVault.getAddress(),
				value: 2000000000000000000000n, // send 2000 ETH
			});

			// create grant (1000 ETH per hour)
			await proxiedVault.createGrant(granteeAddress, 1000, 3600);

			// claim all available balance
			let oldGranteeBalance = await ethers.provider.getBalance(granteeAddress);
			let claimTx = await proxiedVault.connect(grantee).claim(0);

			// await receipt
			let claimTxReceipt = await claimTx.wait();
			let claimTxFees = claimTxReceipt.gasUsed * claimTxReceipt.gasPrice;

			// check balance increase
			let newGranteeBalance = await ethers.provider.getBalance(granteeAddress);
			expect(newGranteeBalance - oldGranteeBalance).to.equal(1000000000000000000000n - claimTxFees); // 1000 ETH increase

			// check claimable balance as grantee
			expect(await proxiedVault.connect(grantee).getClaimableBalance()).to.equal(0); // 0 ETH should be claimable
		});
	});

	describe('Fuzz Testing', function () {
		it('Create Grant fuzzing', async function () {
			const {proxy, token, vault, proxiedVault} = await loadFixture(deployProxyFixture);
			const [owner] = await ethers.getSigners();
			const ownerAddress = await owner.getAddress();

			// add manager role to owner
			let managerRole = await proxiedVault.GRANT_MANAGER_ROLE();
			await proxiedVault.grantRole(managerRole, ownerAddress);

			// diable grant locking after creation / transfer for easier test
			// otherwise, we'd have to wait 10mins after grant creation to see some claimable balance
			await proxiedVault.setClaimTransferLockTime(0);

			for (let i = 0; i < 15; i++) {
				// Get a different signer for each iteration
				// Start from the third signer to avoid using the owner account
				// Note that if the number of fuzz iterations exceed 10, there may
				// be an out-of-bounds on the array returned by the ethers.getSigners()
				// method. By default, Hardhat Network creates 20 accounts for testing.
				const grantee = (await ethers.getSigners())[i + 1];
				const granteeAddress = await grantee.getAddress();

				// Generate random values for grant and claim
				const randomGrant = Math.floor(Math.random() * 1000);
				const randomTime = Math.floor(Math.random() * 3600);

				if (debug) {
					console.log(`Test ${i + 1}: Grantee = ${granteeAddress}, Grant = ${randomGrant}, Time = ${randomTime}`);
				}
		
				// Create grant with random values
				await proxiedVault.createGrant(granteeAddress, randomGrant, randomTime);
				if (debug) {
					console.log('Grant created');
				}

				// check if grant token has been sent to grantee
				expect(await token.balanceOf(granteeAddress)).to.equal(1);
				if (debug) {
					console.log('Token balance check passed');
				}

				// check claimable balance as grantee
				expect(await proxiedVault.connect(grantee).getClaimableBalance()).to.equal(ethers.parseEther(randomGrant.toString()));
				if (debug) {
					console.log('Claimable balance check passed')
				}
			}
		});
	});	
});
