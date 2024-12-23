import { type Chain, defineChain } from 'viem'
// import { anvil, holesky, sepolia } from "wagmi/chains";

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

export const ata_testnet = defineChain({
	id: 1398243,
	name: 'Automata Testnet',
	nativeCurrency: { name: 'Sepolia ATA', symbol: 'SepATA', decimals: 18 },
	rpcUrls: {
		default: {
			http: ['https://1rpc.io/ata/testnet']
		}
	},
	testnet: true
});

export const ata_mainnet = defineChain({
	id: 65536,
	name: 'Automata Mainnet',
	nativeCurrency: { name: 'Automata', symbol: 'ATA', decimals: 18 },
	rpcUrls: {
		default: {
			http: ['https://1rpc.io/ata']
		}
	},
});

const FundingVaultConfig: Config = {
	AppVersion: process.env.REACT_APP_GIT_VERSION as string,
	AdminRole: "0x0000000000000000000000000000000000000000000000000000000000000000",
	ManagerRole: "0xc7386e23c63a3088d7d0389761b7b890e58c103e1a12376eb26d3a4a04e2641b",
	Chains: [
		{
			VaultContractAddr: "0x74bBc82C68fc1e83BFA98cb9A2d6ef8241F46d28",
			TokenContractAddr: "0x26540FCfd36262fbfb49Aa4eC6108B20595b796a",
            TokenName: "SepATA",
			Chain: ata_testnet,
			HumanNetworkName: "Automata Testnet",
			BlockExplorerUrl: "https://explorer-testnet.ata.network/",
		},
		{
			VaultContractAddr: "0x92B59abfE96C9E4bEe808476dD0975a9b89b6c45",
			TokenContractAddr: "0x8618395693B1e4BE231d29E338b2a1f9e645e56F",
            TokenName: "ATA",
			Chain: ata_mainnet,
			HumanNetworkName: "Automata Mainnet",
			BlockExplorerUrl: "https://explorer.ata.network/",
		}
	],
};

export var KnownChains = [
	ata_testnet,
	ata_mainnet
];

let CurrentConfig = FundingVaultConfig;

export default CurrentConfig;