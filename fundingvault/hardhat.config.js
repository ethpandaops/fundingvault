require("@nomicfoundation/hardhat-toolbox");

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
		}
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
