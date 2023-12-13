require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
	paths: {
		sources: "./fundingvault/contracts",
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
