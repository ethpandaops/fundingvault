require("@nomicfoundation/hardhat-toolbox");
require("@nomicfoundation/hardhat-ignition");

// import tasks
require("./scripts/tasks/manager");
require('./scripts/tasks/owner');
require('./scripts/tasks/grantee');

const DEPLOYER_PRIVATE_KEY = vars.has("FUNDINGVAULT_OWNER_PRIVATE_KEY") ? [ vars.get("FUNDINGVAULT_OWNER_PRIVATE_KEY") ] : undefined;

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
	paths: {
		sources: "./contracts",
	},
	networks: {
		hardhat: {
		  	chainId: 1337,
		  	allowBlocksWithSameTimestamp: true,
		  	gas: 8000000,
		 	mining: {
				auto: true,
				interval: 0
		  	}
		},
		anvil: {
			chainId: 31337,
			url: 'http://127.0.0.1:8545',
			accounts: DEPLOYER_PRIVATE_KEY
		},
		sepolia: {
			chainId: 11155111,
			url: `https://rpc.sepolia.ethpandaops.io/`,
      		accounts: DEPLOYER_PRIVATE_KEY,
		},
		holesky: {
			chainId: 17000,
			url: `https://rpc.holesky.ethpandaops.io/`,
      		accounts: DEPLOYER_PRIVATE_KEY,
		},
		ephemery: {
			url: `https://otter.bordel.wtf/erigon`,
      		accounts: DEPLOYER_PRIVATE_KEY,
		},
	},
	solidity: {
		version: "0.8.21",
		settings: {
			optimizer: {
				enabled: true,
				runs: 2000,
			},
		},
	},
};
