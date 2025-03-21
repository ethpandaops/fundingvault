package fundingvault

import (
	"os"

	"gopkg.in/yaml.v3"
)

type FundingVaultConfig struct {
	Privkey  string   `yaml:"privkey"`
	RpcHosts []string `yaml:"rpchosts"`

	FundingVaultAddress string `yaml:"fundingvaultAddress"`
	TxBaseFee           uint64 `yaml:"txBaseFee"`
	TxTipFee            uint64 `yaml:"txTipFee"`
	MinClaimAmount      string `yaml:"minClaimAmount"`

	CheckInterval uint64 `yaml:"checkInterval"`

	RefillRootWallet *FundingVaultRefillConfig   `yaml:"refillRootWallet"`
	RefillAddresses  []*FundingVaultRefillConfig `yaml:"refillAddresses"`
}

type FundingVaultRefillConfig struct {
	Address      string `yaml:"address"`
	Priority     int    `yaml:"priority"`
	MinBalance   string `yaml:"minBalance"`
	RefillAmount string `yaml:"refillAmount"`
	RefillDelay  uint64 `yaml:"refillDelay"`
}

func NewConfig() *FundingVaultConfig {
	return &FundingVaultConfig{
		FundingVaultAddress: "0x610866c6089768dA95524bcc4cE7dB61eDa3931c",
		MinClaimAmount:      "1ETH",
	}
}

func LoadConfig(configFile string) (*FundingVaultConfig, error) {
	config := NewConfig()
	yamlFile, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
