package main

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

func runBalance(cliArgs CliArgs) {
	logrus.SetLevel(logrus.ErrorLevel)
	fv, err := startFundingVault(cliArgs)
	if err != nil {
		logrus.Fatalf("Failed to start funding vault: %v", err)
	}

	defer fv.Shutdown()

	balance, err := fv.GetClaimableBalance(context.Background())
	if err != nil {
		logrus.Fatalf("Failed to get claimable balance: %v", err)
	}

	rootWallet := fv.GetRootWallet()
	fmt.Printf("Wallet address:    %v\n", rootWallet.GetAddress().String())
	fmt.Printf("Wallet balance:    %s\n", formatAmount(rootWallet.GetBalance()))
	fmt.Printf("Claimable balance: %s\n", formatAmount(balance))
}
