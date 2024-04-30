import { type Chain } from 'viem'
import { holesky, sepolia } from "wagmi/chains";
import { defineChain } from "viem";

export interface ChainConfig {
	VaultContractAddr: `0x${string}`;
	TokenContractAddr: `0x${string}`;
    TokenName: string;
	HumanNetworkName: string;
	Chain: Chain;
	BlockExplorerUrl: string;
}

export interface Config {
	AppVersion: string;
	ManagerRole: string;
	Chains: ChainConfig[];
}

export const ephemery = defineChain({
	id: 39438125,
	name: 'Ephemery',
	nativeCurrency: { name: 'Ephemery Ether', symbol: 'Eph', decimals: 18 },
	rpcUrls: {
	  default: {
		http: ['https://rpc.bordel.wtf/test'],
	  },
	},
	blockExplorers: {
	  default: {
		name: 'Etherscan',
		url: 'https://explorer.ephemery.dev/',
		apiUrl: 'https://explorer.ephemery.dev/api',
	  },
	},
	contracts: {
	  multicall3: {
		address: '0x1195eDfF07CC259DF22EF34Ee8FFa7d6C5C0A128',
		blockCreated: 1,
	  },
	  ensRegistry: { address: '0x902740a7Bc8279b1A3beBDf91cf9A016235E8859' },
	  ensUniversalResolver: {
		address: '0xc8Af999e38273D658BE1b921b88A9Ddf005769cC',
		blockCreated: 1,
	  },
	},
	testnet: true,
})

const FundingVaultConfig: Config = {
	AppVersion: process.env.REACT_APP_GIT_VERSION as string,
	ManagerRole: "0xc7386e23c63a3088d7d0389761b7b890e58c103e1a12376eb26d3a4a04e2641b",
	Chains: [
		{
			VaultContractAddr: "0x610866c6089768dA95524bcc4cE7dB61eDa3931c",
			TokenContractAddr: "0x97652A83CC29043fA9Be2781cc0038EBa70de911",
            TokenName: "HolETH",
			Chain: holesky,
			HumanNetworkName: "Holesky",
			BlockExplorerUrl: "https://holesky.etherscan.io/",
		},
		{
			VaultContractAddr: "0x610866c6089768dA95524bcc4cE7dB61eDa3931c",
            TokenContractAddr: "0x97652A83CC29043fA9Be2781cc0038EBa70de911",
            TokenName: "SepETH",
			Chain: sepolia,
			HumanNetworkName: "Sepolia",
			BlockExplorerUrl: "https://sepolia.etherscan.io/",
		},
		{
			VaultContractAddr: "0x610866c6089768dA95524bcc4cE7dB61eDa3931c",
            TokenContractAddr: "0x97652A83CC29043fA9Be2781cc0038EBa70de911",
            TokenName: "EphETH",
			Chain: sepolia,
			HumanNetworkName: "Ephmery",
			BlockExplorerUrl: "https://explorer.ephemery.dev/tx/",
		},
	],
};

let CurrentConfig = FundingVaultConfig;

export default CurrentConfig;