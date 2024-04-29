import { type Chain } from 'viem'
import { holesky, sepolia } from "wagmi/chains";

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
	],
};

let CurrentConfig = FundingVaultConfig;

export default CurrentConfig;