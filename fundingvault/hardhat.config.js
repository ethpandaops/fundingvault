require("@nomicfoundation/hardhat-toolbox");
require("@nomicfoundation/hardhat-ignition");
require("@nomicfoundation/hardhat-verify");

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
		testnet: {
			chainId: 1398243,
			accounts: DEPLOYER_PRIVATE_KEY,
			url: 'https://1rpc.io/ata/testnet'
		},
		mainnet: {
			chainId: 65536,
			accounts: DEPLOYER_PRIVATE_KEY,
			url: 'https://1rpc.io/ata'
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
	etherscan: {
		apiKey: {
		  mainnet: "verifyContract", // apiKey is not required, just set a placeholder
		},
		customChains: [
		  {
			network: "mainnet",
			chainId: 65536,
			urls: {
			  apiURL: "https://api.routescan.io/v2/network/mainnet/evm/65536_2/etherscan",
			  browserURL: "https://explorer.ata.network"
			}
		  }
		]
	  }
};
