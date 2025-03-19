import { type Chain } from 'viem'
import { holesky, sepolia } from "wagmi/chains";
import { defineChain } from 'viem';

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
	AdminRole: string;
	ManagerRole: string;
	Chains: ChainConfig[];
}

const sepoliaWithCustomRPC = Object.assign({}, sepolia, {
	rpcUrls: {
		default: {
			http: ['https://eth-sepolia.g.alchemy.com/v2/74gAuwdkOHanwiWJEl1sYb1rt-5XN3M0'],
		},
	},
});

const holeskyWithCustomRPC = Object.assign({}, holesky, {
	rpcUrls: {
		default: {
			http: ['https://eth-holesky.g.alchemy.com/v2/74gAuwdkOHanwiWJEl1sYb1rt-5XN3M0'],
		},
	},
});

/*
const now = Math.floor((new Date()).getTime() / 1000);
const iteration = Math.floor(((now - 1638471600) / 604800));
export const ephemery = defineChain({
	id: 39438000 + iteration,
	name: 'Ephemery',
	nativeCurrency: { name: 'Ephemery Ether', symbol: 'Eph', decimals: 18 },
	rpcUrls: {
	  default: {
		http: ['https://rpc.bordel.wtf/test', 'https://otter.bordel.wtf/erigon'],
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
*/

export const hoodiWithCustomRPC = /*#__PURE__*/ defineChain({
	id: 560048,
	name: 'Hoodi',
	nativeCurrency: { name: 'Hoodi Ether', symbol: 'ETH', decimals: 18 },
	rpcUrls: {
	  default: {
		http: [
			'https://rpc.hoodi.ethpandaops.io'
		],
	  },
	},
	blockExplorers: {
	  default: {
		name: 'Etherscan',
		url: 'https://holesky.etherscan.io',
	  },
	},
	
	contracts: {
	  /*
	  multicall3: {
		address: '0xca11bde05977b3631167028862be2a173976ca11',
		blockCreated: 77,
	  },
	  ensRegistry: {
		address: '0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e',
		blockCreated: 801613,
	  },
	  ensUniversalResolver: {
		address: '0xa6AC935D4971E3CD133b950aE053bECD16fE7f3b',
		blockCreated: 973484,
	  },
	  */
	},
	
	testnet: true,
  })

const FundingVaultConfig: Config = {
	AppVersion: process.env.REACT_APP_GIT_VERSION as string,
	AdminRole: "0x0000000000000000000000000000000000000000000000000000000000000000",
	ManagerRole: "0xc7386e23c63a3088d7d0389761b7b890e58c103e1a12376eb26d3a4a04e2641b",
	Chains: [
		{
			VaultContractAddr: "0x610866c6089768dA95524bcc4cE7dB61eDa3931c",
			TokenContractAddr: "0x97652A83CC29043fA9Be2781cc0038EBa70de911",
            TokenName: "HooETH",
			Chain: hoodiWithCustomRPC,
			HumanNetworkName: "Hoodi",
			BlockExplorerUrl: "https://hoodi.cloud.blockscout.com/",
		},
		{
			VaultContractAddr: "0x610866c6089768dA95524bcc4cE7dB61eDa3931c",
			TokenContractAddr: "0x97652A83CC29043fA9Be2781cc0038EBa70de911",
            TokenName: "HolETH",
			Chain: holeskyWithCustomRPC,
			HumanNetworkName: "Holesky",
			BlockExplorerUrl: "https://holesky.etherscan.io/",
		},
		{
			VaultContractAddr: "0x610866c6089768dA95524bcc4cE7dB61eDa3931c",
            TokenContractAddr: "0x97652A83CC29043fA9Be2781cc0038EBa70de911",
            TokenName: "SepETH",
			Chain: sepoliaWithCustomRPC,
			HumanNetworkName: "Sepolia",
			BlockExplorerUrl: "https://sepolia.etherscan.io/",
		},
		/*
		{
			VaultContractAddr: "0x610866c6089768dA95524bcc4cE7dB61eDa3931c",
            TokenContractAddr: "0x97652A83CC29043fA9Be2781cc0038EBa70de911",
            TokenName: "EphETH",
			Chain: ephemery,
			HumanNetworkName: "Ephmery",
			BlockExplorerUrl: "https://explorer.ephemery.dev/",
		},
		*/
	],
};

export var KnownChains = [
	holeskyWithCustomRPC,
	hoodiWithCustomRPC,
    sepoliaWithCustomRPC,
    //ephemery,
];

let CurrentConfig = FundingVaultConfig;

export default CurrentConfig;