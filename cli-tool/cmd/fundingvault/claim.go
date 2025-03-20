package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
)

func runClaim(cliArgs CliArgs, claimArgs ClaimArgs) {
	logrus.SetLevel(logrus.ErrorLevel)
	fv, err := startFundingVault(cliArgs)
	if err != nil {
		logrus.Fatalf("Failed to start funding vault: %v", err)
	}

	defer fv.Shutdown()

	// Parse amount as decimal ETH and convert to Wei
	amount, err := fv.ParseEthAmount(claimArgs.amount)
	if err != nil {
		logrus.Fatalf("Failed to parse amount: %v", err)
	}

	if !claimArgs.force {
		claimable, err := fv.GetClaimableBalance(context.Background())
		if err != nil {
			logrus.Fatalf("Failed to get claimable balance: %v", err)
		}
		if amount.Cmp(claimable) > 0 || claimable.Cmp(big.NewInt(0)) <= 0 {
			logrus.Fatalf("Amount exceeds claimable balance. Use --force to override.")
		}
	}

	totalClaimed, err := fv.ClaimSync(context.Background(), amount, func(tx *types.Transaction) {
		fmt.Printf("Claim transaction sent: %v\n", tx.Hash().String())
	})
	if err != nil {
		logrus.Fatalf("Failed to send claim transaction: %v", err)
	}

	fmt.Printf("Claimed %s\n", formatAmount(totalClaimed))
}
