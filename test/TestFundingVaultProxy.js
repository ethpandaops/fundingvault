const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("FundingVaultProxyTest", function () {
	let user;
	let FundingVaultProxy;
	let FundingVaultProxyContract;

	beforeEach(async () => {
		[user] = await ethers.getSigners();
		FundingVaultProxy = await ethers.getContractFactory("FundingVaultProxy");
		FundingVaultProxyContract = await FundingVaultProxy.deploy({from: user.address});
	});

	it("Proxy default implementation is at address zero", async function () {
		expect(await FundingVaultProxyContract.implementation()).to.equal("0x0000000000000000000000000000000000000000");
	});
});