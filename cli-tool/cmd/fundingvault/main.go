package main

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ethpandaops/fundingvault"
)

type CliArgs struct {
	verbose             bool
	config              string
	rpchosts            []string
	privkey             string
	fundingvaultAddress string
	baseFee             uint64
	tipFee              uint64
}

type MainArgs struct {
	amount string
	daemon bool
}

type ClaimArgs struct {
	amount string
	force  bool
}

func main() {
	cliArgs := CliArgs{}
	mainArgs := MainArgs{}
	rootCmd := &cobra.Command{
		Use:   "fundingvault",
		Short: "Funding vault CLI tool",
		Run: func(cmd *cobra.Command, args []string) {
			runMain(cliArgs, mainArgs)
		},
	}

	rootCmd.PersistentFlags().BoolVarP(&cliArgs.verbose, "verbose", "v", false, "Run the script with verbose output")
	rootCmd.PersistentFlags().StringVarP(&cliArgs.config, "config", "c", "", "The config file to use.")
	rootCmd.PersistentFlags().StringVarP(&cliArgs.privkey, "privkey", "p", "", "The private key of the wallet to send funds from.")
	rootCmd.PersistentFlags().StringSliceVarP(&cliArgs.rpchosts, "rpchost", "r", []string{}, "The RPC host to send transactions to.")
	rootCmd.PersistentFlags().StringVar(&cliArgs.fundingvaultAddress, "fundingvault", "", "The address of the funding vault.")

	rootCmd.Flags().Uint64Var(&cliArgs.baseFee, "basefee", 20, "Max fee per gas to use in claim transaction (in gwei)")
	rootCmd.Flags().Uint64Var(&cliArgs.tipFee, "tipfee", 2, "Max tip per gas to use in claim transaction (in gwei)")
	rootCmd.Flags().BoolVarP(&mainArgs.daemon, "daemon", "d", false, "Run the script in daemon mode")
	rootCmd.Flags().StringVar(&mainArgs.amount, "amount", "", "The amount of funds to claim for the root wallet")

	rootCmd.AddCommand(&cobra.Command{
		Use:   "balance",
		Short: "Show wallet & claimable balance",
		Run: func(cmd *cobra.Command, args []string) {
			runBalance(cliArgs)
		},
	})

	claimArgs := ClaimArgs{}
	claimCmd := &cobra.Command{
		Use:   "claim [amount]",
		Short: "Claim funds from the vault",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			claimArgs.amount = args[0]
			runClaim(cliArgs, claimArgs)
		},
	}
	claimCmd.Flags().Uint64Var(&cliArgs.baseFee, "basefee", 20, "Max fee per gas to use in claim transaction (in gwei)")
	claimCmd.Flags().Uint64Var(&cliArgs.tipFee, "tipfee", 2, "Max tip per gas to use in claim transaction (in gwei)")
	claimCmd.Flags().BoolVarP(&claimArgs.force, "force", "f", false, "Force claim even if amount exceeds claimable balance")
	rootCmd.AddCommand(claimCmd)

	rootCmd.Execute()
}

func startFundingVault(cliArgs CliArgs) (*fundingvault.FundingVault, error) {
	if cliArgs.verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}

	var config *fundingvault.FundingVaultConfig
	if cliArgs.config != "" {
		var err error
		config, err = fundingvault.LoadConfig(cliArgs.config)
		if err != nil {
			panic(err)
		}
	} else {
		config = fundingvault.NewConfig()
	}

	for _, rpcHost := range strings.Split(strings.Join(cliArgs.rpchosts, ","), ",") {
		if rpcHost != "" {
			config.RpcHosts = append(config.RpcHosts, rpcHost)
		}
	}

	if cliArgs.privkey != "" {
		config.Privkey = cliArgs.privkey
	}

	if cliArgs.fundingvaultAddress != "" {
		config.FundingVaultAddress = cliArgs.fundingvaultAddress
	}

	if cliArgs.baseFee > 0 {
		config.TxBaseFee = cliArgs.baseFee
	}

	if cliArgs.tipFee > 0 {
		config.TxTipFee = cliArgs.tipFee
	}

	if config.RefillRootWallet == nil {
		config.RefillRootWallet = &fundingvault.FundingVaultRefillConfig{}
	}

	logger := logrus.New()
	logger.SetLevel(logrus.GetLevel())

	fv := fundingvault.NewFundingVault(context.Background(), config, logger)
	if err := fv.Initialize(); err != nil {
		fv.Shutdown()
		return nil, err
	}

	return fv, nil
}

func formatAmount(amount *big.Int) string {
	value := new(big.Int).Set(amount)
	gwei, _ := value.Div(value, big.NewInt(1e9)).Float64()
	if gwei > 0 && gwei < 100000 {
		return fmt.Sprintf("%.0f gwei", gwei)
	}

	decimalVal := fmt.Sprintf("%.4f", gwei/1e9)
	for i := len(decimalVal) - 1; i >= 0; i-- {
		if decimalVal[i] == '0' {
			decimalVal = decimalVal[:i]
		} else {
			break
		}
	}

	if decimalVal[len(decimalVal)-1] == '.' {
		decimalVal = decimalVal[:len(decimalVal)-1]
	}

	return fmt.Sprintf("%s ETH", decimalVal)
}

func runMain(cliArgs CliArgs, mainArgs MainArgs) {
	if !mainArgs.daemon {
		logrus.SetLevel(logrus.ErrorLevel)
	}

	fv, err := startFundingVault(cliArgs)
	if err != nil {
		logrus.Fatalf("Failed to start funding vault: %v", err)
	}

	defer fv.Shutdown()

	if mainArgs.amount != "" {
		fv.Config.RefillRootWallet.MinBalance = mainArgs.amount
	} else if fv.Config.RefillRootWallet.MinBalance == "" {
		fv.Config.RefillRootWallet.MinBalance = "1ETH"
	}

	if mainArgs.daemon {
		for {
			requests, err := fv.CheckAndRefillWallets(context.Background(), func(request fundingvault.FundingRequest, tx *types.Transaction) {
				logrus.Infof("Sending funding request for %v: %v (tx: %v)", request.Address, formatAmount(request.Request), tx.Hash())
			})
			if err != nil {
				fv.Logger.Errorf("Failed to check and refill wallets: %v", err)
			} else {
				funded := 0
				checked := 0
				if len(requests) > 0 {
					for _, request := range requests {
						if request.Error != nil {
							fv.Logger.Errorf("request for %v failed: %v", request.Address, request.Error)
						} else if request.Request.Cmp(big.NewInt(0)) > 0 {
							fv.Logger.Infof("request for %v: %v", request.Address, formatAmount(request.Request))
							funded++
						}
						checked++
					}
				}

				fv.Logger.Infof("checked %v wallets, funded %v", checked, funded)
			}

			sleepTime := time.Duration(fv.Config.CheckInterval) * time.Second
			if sleepTime < time.Second*30 {
				sleepTime = time.Second * 30
			}
			time.Sleep(sleepTime)
		}
	} else {
		fmt.Println("FundingVault cli-tool")
		fmt.Println("=======================")

		balance, err := fv.GetClaimableBalance(context.Background())
		if err != nil {
			logrus.Fatalf("Failed to get claimable balance: %v", err)
		}

		rootWallet := fv.GetRootWallet()
		fmt.Printf("Wallet address:    %v\n", rootWallet.GetAddress().String())
		fmt.Printf("Vault contract:    %v\n", fv.GetVaultAddress().String())
		fmt.Printf("Wallet balance:    %s\n", formatAmount(rootWallet.GetBalance()))
		fmt.Printf("Claimable balance: %s\n", formatAmount(balance))
		fmt.Println("")
		fmt.Println("Checking wallets...")

		requests, err := fv.CheckAndRefillWallets(context.Background(), func(request fundingvault.FundingRequest, tx *types.Transaction) {
			fmt.Printf("Sending funding request for %v: %v (tx: %v)\n", request.Address, formatAmount(request.Request), tx.Hash())
		})
		if err != nil {
			fmt.Printf("Failed to check and refill wallets: %v", err)
			return
		}

		fmt.Println("Checked wallets:")
		for _, request := range requests {
			if request.Error != nil {
				fmt.Printf("  %v (balance: %v, min: %v): ERROR: %v\n", request.Address, formatAmount(request.Balance), formatAmount(request.MinAmount), request.Error)
			} else if request.Request.Cmp(big.NewInt(0)) == 0 {
				fmt.Printf("  %v (balance: %v, min: %v): no funding needed\n", request.Address, formatAmount(request.Balance), formatAmount(request.MinAmount))
			} else {
				fmt.Printf("  %v (balance: %v, min: %v): %v\n", request.Address, formatAmount(request.Balance), formatAmount(request.MinAmount), formatAmount(request.Request))
			}
		}
	}
}
