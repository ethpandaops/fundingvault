const {	loadFixture } = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { time } = require("@nomicfoundation/hardhat-network-helpers");
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

	describe("Grant Management (as owner)", function () {
		async function prepareTest(proxiedVault) {
			const [owner, grantee] = await ethers.getSigners();
			const ownerAddress = await owner.getAddress();

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
		}

		it("Create Grant", async function () {
			const {proxy, token, vault, proxiedVault} = await loadFixture(deployProxyFixture);
			const [owner, grantee] = await ethers.getSigners();
			await prepareTest(proxiedVault);

			// create grant (1000 ETH per hour)
			await proxiedVault.createGrant(grantee.getAddress(), 1000, 3600);

			// check if grant token has been sent to grantee
			expect(await token.balanceOf(grantee.getAddress())).to.equal(1);

			// check claimable balance as grantee
			expect(await proxiedVault.connect(grantee).getClaimableBalance()).to.equal(1000000000000000000000n); // 1000 ETH should be claimable
		});
		it("Update Grant amount (max unclaimed balance)", async function () {
			const {proxy, token, vault, proxiedVault} = await loadFixture(deployProxyFixture);
			const [owner, grantee] = await ethers.getSigners();
			await prepareTest(proxiedVault);

			// create grant (1000 ETH per hour)
			await proxiedVault.createGrant(grantee.getAddress(), 1000, 3600);

			// check if grant token has been sent to grantee
			expect(await token.balanceOf(grantee.getAddress())).to.equal(1);

			// update grant (500 ETH per hour)
			await proxiedVault.updateGrant(1, 500, 3600);

			// check if grantee still holds the grant token
			expect(await token.balanceOf(grantee.getAddress())).to.equal(1);

			// check claimable balance as grantee
			expect(await proxiedVault.connect(grantee).getClaimableBalance()).to.equal(500000000000000000000n); // 500 ETH should be claimable
		});
		it("Update Grant amount (no unclaimed balance)", async function () {
			const {proxy, token, vault, proxiedVault} = await loadFixture(deployProxyFixture);
			const [owner, grantee] = await ethers.getSigners();
			await prepareTest(proxiedVault);

			// create grant (1000 ETH per hour)
			await proxiedVault.createGrant(grantee.getAddress(), 1000, 3600);

			// check if grant token has been sent to grantee
			expect(await token.balanceOf(grantee.getAddress())).to.equal(1);

			// claim all available balance
			await proxiedVault.connect(grantee).claim(0);

			// update grant (500 ETH per hour)
			await proxiedVault.updateGrant(1, 500, 3600);

			// check claimable balance as grantee
			expect(await proxiedVault.connect(grantee).getClaimableBalance()).to.equal(0n); // 0 ETH should be claimable
		});
		it("Update Grant amount (half unclaimed balance)", async function () {
			const {proxy, token, vault, proxiedVault} = await loadFixture(deployProxyFixture);
			const [owner, grantee] = await ethers.getSigners();
			await prepareTest(proxiedVault);

			// create grant (1000 ETH per hour)
			await proxiedVault.createGrant(grantee.getAddress(), 1000, 3600);

			// check if grant token has been sent to grantee
			expect(await token.balanceOf(grantee.getAddress())).to.equal(1);

			// claim all available balance
			await proxiedVault.connect(grantee).claim(0);

			// "wait" 30mins
			await time.increase(1800);

			// check claimable balance as grantee
			expect(await proxiedVault.connect(grantee).getClaimableBalance()).to.equal(500000000000000000000n); // 500 ETH should be claimable

			// update grant (500 ETH per hour)
			await proxiedVault.updateGrant(1, 500, 3600);

			// check claimable balance as grantee
			expect(await proxiedVault.connect(grantee).getClaimableBalance()).to.equal(250000000000000000000n); // 250 ETH should be claimable
		});
		it("Update Grant interval (max unclaimed balance)", async function () {
			const {proxy, token, vault, proxiedVault} = await loadFixture(deployProxyFixture);
			const [owner, grantee] = await ethers.getSigners();
			await prepareTest(proxiedVault);

			// create grant (1000 ETH per hour)
			await proxiedVault.createGrant(grantee.getAddress(), 1000, 3600);

			// check if grant token has been sent to grantee
			expect(await token.balanceOf(grantee.getAddress())).to.equal(1);

			// update grant (1000 ETH per 2 hours)
			await proxiedVault.updateGrant(1, 1000, 7200);

			// check claimable balance as grantee
			expect(await proxiedVault.connect(grantee).getClaimableBalance()).to.equal(500000000000000000000n); // 500 ETH should be claimable
		});
		it("Update Grant interval (no unclaimed balance)", async function () {
			const {proxy, token, vault, proxiedVault} = await loadFixture(deployProxyFixture);
			const [owner, grantee] = await ethers.getSigners();
			await prepareTest(proxiedVault);

			// create grant (1000 ETH per hour)
			await proxiedVault.createGrant(grantee.getAddress(), 1000, 3600);

			// check if grant token has been sent to grantee
			expect(await token.balanceOf(grantee.getAddress())).to.equal(1);

			// claim all available balance
			await proxiedVault.connect(grantee).claim(0);

			// update grant (1000 ETH per 2 hours)
			await proxiedVault.updateGrant(1, 1000, 7200);

			// check claimable balance as grantee
			expect(await proxiedVault.connect(grantee).getClaimableBalance()).to.equal(0n); // 0 ETH should be claimable
		});
		it("Update Grant interval (half unclaimed balance)", async function () {
			const {proxy, token, vault, proxiedVault} = await loadFixture(deployProxyFixture);
			const [owner, grantee] = await ethers.getSigners();
			await prepareTest(proxiedVault);

			// create grant (1000 ETH per hour)
			await proxiedVault.createGrant(grantee.getAddress(), 1000, 3600);

			// check if grant token has been sent to grantee
			expect(await token.balanceOf(grantee.getAddress())).to.equal(1);

			// claim all available balance
			await proxiedVault.connect(grantee).claim(0);

			// "wait" 30mins
			await time.increase(1800);

			// check claimable balance as grantee
			expect(await proxiedVault.connect(grantee).getClaimableBalance()).to.equal(500000000000000000000n); // 500 ETH should be claimable

			// update grant (1000 ETH per 2 hour)
			await proxiedVault.updateGrant(1, 1000, 7200);

			// check claimable balance as grantee
			expect(await proxiedVault.connect(grantee).getClaimableBalance()).to.equal(250000000000000000000n); // 250 ETH should be claimable
		});
		it("Delete Grant", async function () {
			const {proxy, token, vault, proxiedVault} = await loadFixture(deployProxyFixture);
			const [owner, grantee] = await ethers.getSigners();
			await prepareTest(proxiedVault);

			// create grant (1000 ETH per hour)
			await proxiedVault.createGrant(grantee.getAddress(), 1000, 3600);

			// check if grant token has been sent to grantee
			expect(await token.balanceOf(grantee.getAddress())).to.equal(1);

			// remove grant
			await proxiedVault.removeGrant(1);

			// check if grant token has been burned
			expect(await token.balanceOf(grantee.getAddress())).to.equal(0);

			// check claimable balance as grantee
			expect(await proxiedVault.connect(grantee).getClaimableBalance()).to.equal(0n); // 0 ETH should be claimable
		});
		it("Transfer Grant", async function () {
			const {proxy, token, vault, proxiedVault} = await loadFixture(deployProxyFixture);
			const [owner, grantee, grantee2] = await ethers.getSigners();
			await prepareTest(proxiedVault);

			// create grant (1000 ETH per hour)
			await proxiedVault.createGrant(grantee.getAddress(), 1000, 3600);

			// check if grant token has been sent to grantee
			expect(await token.balanceOf(grantee.getAddress())).to.equal(1);

			// transfer grant
			await proxiedVault.transferGrant(1, grantee2.getAddress());

			// check if grant token has been moved
			expect(await token.balanceOf(grantee.getAddress())).to.equal(0);
			expect(await token.balanceOf(grantee2.getAddress())).to.equal(1);

			// check claimable balance
			expect(await proxiedVault.connect(grantee).getClaimableBalance()).to.equal(0n); // 0 ETH should be claimable
			expect(await proxiedVault.connect(grantee2).getClaimableBalance()).to.equal(1000000000000000000000n); // 1000 ETH should be claimable
		});
	});
	describe("Manager Limits", function () {
		async function prepareTest(proxiedVault) {
			const [owner, manager, grantee] = await ethers.getSigners();

			// add manager role to owner
			let managerRole = await proxiedVault.GRANT_MANAGER_ROLE();
			await proxiedVault.grantRole(managerRole, owner.getAddress());

			// add manager role to manager
			await proxiedVault.grantRole(managerRole, manager.getAddress());

			// diable grant locking after creation / transfer for easier test
			// otherwise, we'd have to wait 10mins after grant creation to see some claimable balance
			await proxiedVault.setClaimTransferLockTime(0);

			// configure manager limits for tests
			await proxiedVault.setManagerGrantLimits(
				1000, // amount
				3600, // interval
				3600, // cooldown (number of seconds added to the cooldown clock when a grant worth amount/interval got managed)
				1800, // cooldownLock (lock manager when cooldown clock is above this value)
			);
			/* The manager limits above have the following effect:
			* - managers are limited to create/update/transfer grants with a maximum allowance of amount / interval ETH.
		    *   with the values above, manages are limited to manage grants worth  1000ETH/3600sec = 0.277 ETH/sec
			* - managers have their own cooldown system to avoid mass creation of grants from hijacked wallets.
			*   when a manager creates/updates/transfers a grant with the max allowance (amount / interval), `cooldown` seconds are added to the managers cooldown clock.
			*   if the cooldown clock exceeds the `cooldownLock` value, the manager cannot create/update/transfer grants anymore and needs to wait for the clock value to be lower than `cooldownLock` again.
			*   with the values above, managers can create one grant with 1000ETH/3600sec, which adds 3600secs to their cooldown clock.
			*   they are then locked for 1800secs, before they can create another grant.
			*/

			// send some funds to the vault, otherwise there is nothing to claim
			await owner.sendTransaction({
				to: await proxiedVault.getAddress(),
				value: 2000000000000000000000n, // send 2000 ETH
			});
		}

		it("Grant creation limits", async function () {
			const {proxy, token, vault, proxiedVault} = await loadFixture(deployProxyFixture);
			const [owner, manager, grantee] = await ethers.getSigners();
			await prepareTest(proxiedVault);

			// create grant (2000 ETH per hour), should fail
			var txErr = null;
			try {
				await proxiedVault.connect(manager).createGrant(grantee.getAddress(), 2000, 3600);
			} catch(ex) {
				txErr = ex;
			}
			expect(txErr?.toString()).to.match(/amount exceeds manager limits/);

			// create 2 grants (500 ETH per hour)
			await proxiedVault.connect(manager).createGrant(grantee.getAddress(), 250, 3600);
			await proxiedVault.connect(manager).createGrant(grantee.getAddress(), 750, 3600);

			// check manager cooldown
			expect(await proxiedVault.getManagerCooldown(manager.getAddress())).to.equal(3601);

			// create grant (500 ETH per hour), should fail
			var txErr = null;
			try {
				await proxiedVault.connect(manager).createGrant(grantee.getAddress(), 500, 3600);
			} catch(ex) {
				txErr = ex;
			}
			expect(txErr?.toString()).to.match(/manager cooldown/);

			// "wait" 30mins
			await time.increase(1802);

			// create grant (1000 ETH per hour)
			await proxiedVault.connect(manager).createGrant(grantee.getAddress(), 1000, 3600);

			// check manager cooldown
			expect(await proxiedVault.getManagerCooldown(manager.getAddress())).to.equal(5400);
		});
		it("Grant update limits", async function () {
			const {proxy, token, vault, proxiedVault} = await loadFixture(deployProxyFixture);
			const [owner, manager, grantee] = await ethers.getSigners();
			await prepareTest(proxiedVault);

			// create grant (1000 ETH per hour)
			await proxiedVault.createGrant(grantee.getAddress(), 1000, 3600);

			// update grant (500 ETH per hour)
			// this is a decrease of the grant allowance, so it shouldn't affect the manager cooldown
			await proxiedVault.connect(manager).updateGrant(1, 500, 3600);

			// check manager cooldown
			expect(await proxiedVault.getManagerCooldown(manager.getAddress())).to.equal(0);

			// update grant (1000 ETH per hour)
			// this is a increase of the grant allowance by 500 ETH/3600sec
			await proxiedVault.connect(manager).updateGrant(1, 1000, 3600);

			// check manager cooldown
			expect(await proxiedVault.getManagerCooldown(manager.getAddress())).to.equal(1801);
		});
		it("Grant transfer limits", async function () {
			const {proxy, token, vault, proxiedVault} = await loadFixture(deployProxyFixture);
			const [owner, manager, grantee, grantee2] = await ethers.getSigners();
			await prepareTest(proxiedVault);

			// create test grants
			await proxiedVault.createGrant(grantee.getAddress(), 100, 3600);
			await proxiedVault.createGrant(grantee.getAddress(), 100, 3600);
			await proxiedVault.createGrant(grantee.getAddress(), 1000, 3600);
			await proxiedVault.createGrant(grantee.getAddress(), 1000, 3600);

			// transfer grant (100 ETH per hour)
			await proxiedVault.connect(manager).transferGrant(1, grantee2.getAddress());

			// check manager cooldown
			expect(await proxiedVault.getManagerCooldown(manager.getAddress())).to.equal(359);

			// transfer grant (100 ETH per hour)
			await proxiedVault.connect(manager).transferGrant(2, grantee2.getAddress());

			// check manager cooldown
			expect(await proxiedVault.getManagerCooldown(manager.getAddress())).to.equal(718);

			// transfer grant (1000 ETH per hour)
			await proxiedVault.connect(manager).transferGrant(3, grantee2.getAddress());

			// check manager cooldown
			expect(await proxiedVault.getManagerCooldown(manager.getAddress())).to.equal(4318);

			// transfer grant (1000 ETH per hour), should fail
			var txErr = null;
			try {
				await proxiedVault.connect(manager).transferGrant(4, grantee2.getAddress());
			} catch(ex) {
				txErr = ex;
			}
			expect(txErr?.toString()).to.match(/manager cooldown/);

			// "wait" 2519 sec (4318 - 1800) + 1
			await time.increase(2519);

			// transfer grant (1000 ETH per hour)
			await proxiedVault.connect(manager).transferGrant(4, grantee2.getAddress());

			// check manager cooldown
			expect(await proxiedVault.getManagerCooldown(manager.getAddress())).to.equal(5399);
		});
		
	});

	describe("Request funds", function () {
		async function prepareTest(proxiedVault) {
			const [owner, grantee] = await ethers.getSigners();
			const ownerAddress = await owner.getAddress();

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
		}

		it("Request max available balance", async function () {
			const {proxy, token, vault, proxiedVault} = await loadFixture(deployProxyFixture);
			const [owner, grantee] = await ethers.getSigners();
			await prepareTest(proxiedVault);

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
		it("Request funds in multiple small claims", async function () {
			const {proxy, token, vault, proxiedVault} = await loadFixture(deployProxyFixture);
			const [owner, grantee] = await ethers.getSigners();
			await prepareTest(proxiedVault);

			// create grant (1000 ETH per hour)
			await proxiedVault.createGrant(grantee.getAddress(), 1000, 3600);

			for(let i = 0; i < 10; i++) {
				// claim 100 ETH (1/10 of granted amount)
				let oldGranteeBalance = await ethers.provider.getBalance(grantee.getAddress());
				let claimTx = await proxiedVault.connect(grantee).claim(100000000000000000000n);

				// await receipt
				let claimTxReceipt = await claimTx.wait();
				let claimTxFees = claimTxReceipt.gasUsed * claimTxReceipt.gasPrice;

				// check balance increase
				let newGranteeBalance = await ethers.provider.getBalance(grantee.getAddress());
				expect(newGranteeBalance - oldGranteeBalance).to.equal(100000000000000000000n - claimTxFees); // 100 ETH increase
			}

			// check claimable balance as grantee
			expect(await proxiedVault.connect(grantee).getClaimableBalance()).to.equal(0); // 0 ETH should be claimable
		});
		it("Request funds after half interval", async function () {
			const {proxy, token, vault, proxiedVault} = await loadFixture(deployProxyFixture);
			const [owner, grantee] = await ethers.getSigners();
			await prepareTest(proxiedVault);

			// create grant (1000 ETH per hour)
			await proxiedVault.createGrant(grantee.getAddress(), 1000, 3600);

			// claim all available balance
			await proxiedVault.connect(grantee).claim(0);

			// "wait" 30mins
			await time.increase(1800);

			// claim all available balance again
			let oldGranteeBalance = await ethers.provider.getBalance(grantee.getAddress());
			let claimTx = await proxiedVault.connect(grantee).claim(0);

			// await receipt
			let claimTxReceipt = await claimTx.wait();
			let claimTxFees = claimTxReceipt.gasUsed * claimTxReceipt.gasPrice;

			// check balance increase
			let newGranteeBalance = await ethers.provider.getBalance(grantee.getAddress());
			expect(newGranteeBalance - oldGranteeBalance).to.equal(500000000000000000000n - claimTxFees); // 500 ETH increase

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
