require("dotenv").config();
const { OWNER_ADDRESS } = process.env;

const FundingVaultProxy = artifacts.require("FundingVaultProxy");
const FundingVaultToken = artifacts.require("FundingVaultToken");
const FundingVaultV1 = artifacts.require("FundingVaultV1");

module.exports = async function (deployer) {
  console.log("deploying FundingVaultProxy...");
  await deployer.deploy(FundingVaultProxy);
  var proxy = await FundingVaultProxy.deployed();

  console.log("deploying FundingVaultToken...");
  await deployer.deploy(FundingVaultToken, proxy.address);
  var token = await FundingVaultToken.deployed();

  console.log("deploying FundingVaultV1...");
  await deployer.deploy(FundingVaultV1);
  var logic = await FundingVaultV1.deployed();

  console.log("upgrading proxy and calling initialize()...");
  var initCall = logic.contract.methods.initialize(token.address).encodeABI();
  await proxy.upgradeToAndCall(logic.address, initCall);

  console.log("grant admin role to " + OWNER_ADDRESS + "...");
  var vault = await FundingVaultV1.at(proxy.address);
  var ownerRole = await logic.DEFAULT_ADMIN_ROLE.call();
  await vault.grantRole(ownerRole, OWNER_ADDRESS);

  console.log("deployment complete");
};
