package fundingvault

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethpandaops/spamoor/txbuilder"
	"github.com/holiman/uint256"
	"github.com/sirupsen/logrus"
)

func (fv *FundingVault) ClaimAsync(ctx context.Context, amount *big.Int, onConfirm func(amount *big.Int, err error)) (*types.Transaction, error) {
	rootWallet := fv.GetRootWallet()
	vaultContract := fv.GetVaultContract()
	tx, err := rootWallet.BuildBoundTx(&txbuilder.TxMetadata{
		GasFeeCap: uint256.MustFromBig(big.NewInt(int64(fv.Config.TxBaseFee * 1e9))),
		GasTipCap: uint256.MustFromBig(big.NewInt(int64(fv.Config.TxTipFee * 1e9))),
		Gas:       200000,
		Value:     uint256.NewInt(0),
	}, func(transactOpts *bind.TransactOpts) (*types.Transaction, error) {
		return vaultContract.Claim(transactOpts, amount)
	})
	if err != nil {
		return nil, fmt.Errorf("failed to build claim transaction: %v", err)
	}

	err = fv.sendClaimAsync(ctx, tx, onConfirm)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (fv *FundingVault) ClaimSync(ctx context.Context, amount *big.Int, onSubmit func(tx *types.Transaction)) (*big.Int, error) {
	txWg := sync.WaitGroup{}
	txWg.Add(1)

	txErr := error(nil)
	amountClaimed := new(big.Int)
	tx, err := fv.ClaimAsync(ctx, amount, func(amount *big.Int, err error) {
		defer txWg.Done()

		if err != nil {
			txErr = err
			return
		}
		amountClaimed = amount
	})
	if err != nil {
		return nil, err
	}

	if onSubmit != nil {
		onSubmit(tx)
	}

	txWg.Wait()

	return amountClaimed, txErr
}

func (fv *FundingVault) ClaimToAsync(ctx context.Context, toAddress common.Address, amount *big.Int, onConfirm func(amount *big.Int, err error)) (*types.Transaction, error) {
	rootWallet := fv.GetRootWallet()
	vaultContract := fv.GetVaultContract()
	tx, err := rootWallet.BuildBoundTx(&txbuilder.TxMetadata{
		GasFeeCap: uint256.MustFromBig(big.NewInt(int64(fv.Config.TxBaseFee * 1e9))),
		GasTipCap: uint256.MustFromBig(big.NewInt(int64(fv.Config.TxTipFee * 1e9))),
		Gas:       200000,
		Value:     uint256.NewInt(0),
	}, func(transactOpts *bind.TransactOpts) (*types.Transaction, error) {
		return vaultContract.ClaimTo0(transactOpts, amount, toAddress)
	})
	if err != nil {
		return nil, fmt.Errorf("failed to build claim transaction: %v", err)
	}

	err = fv.sendClaimAsync(ctx, tx, onConfirm)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (fv *FundingVault) ClaimToSync(ctx context.Context, toAddress common.Address, amount *big.Int, onSubmit func(tx *types.Transaction)) (*big.Int, error) {
	txWg := sync.WaitGroup{}
	txWg.Add(1)

	txErr := error(nil)
	amountClaimed := new(big.Int)
	tx, err := fv.ClaimToAsync(ctx, toAddress, amount, func(amount *big.Int, err error) {
		defer txWg.Done()

		if err != nil {
			txErr = err
			return
		}
		amountClaimed = amount
	})
	if err != nil {
		return nil, err
	}

	if onSubmit != nil {
		onSubmit(tx)
	}

	txWg.Wait()

	return amountClaimed, txErr
}

func (fv *FundingVault) sendClaimAsync(ctx context.Context, tx *types.Transaction, onConfirm func(amount *big.Int, err error)) error {
	txPool := fv.GetTxPool()
	txWg := sync.WaitGroup{}
	txWg.Add(1)
	txErr := error(nil)
	txReceipt := (*types.Receipt)(nil)
	err := txPool.SendTransaction(ctx, fv.GetRootWallet(), tx, &txbuilder.SendTransactionOptions{
		Client:              fv.GetClient(0, false),
		MaxRebroadcasts:     10,
		RebroadcastInterval: 30 * time.Second,
		OnConfirm: func(tx *types.Transaction, receipt *types.Receipt, err error) {
			defer func() {
				txWg.Done()
			}()

			txErr = err
			txReceipt = receipt
		},
		LogFn: func(client *txbuilder.Client, retry, rebroadcast int, err error) {
			if err != nil {
				fv.Logger.WithFields(logrus.Fields{
					"retry":       retry,
					"rebroadcast": rebroadcast,
					"client":      client.GetName(),
					"err":         err,
				}).Errorf("Failed to send claim transaction")
			} else {
				fv.Logger.WithFields(logrus.Fields{
					"retry":       retry,
					"rebroadcast": rebroadcast,
					"client":      client.GetName(),
				}).Infof("Claim transaction sent")
			}
		},
	})
	if err != nil {
		return fmt.Errorf("failed to send claim transaction: %v", err)
	}

	go func() {
		txWg.Wait()
		if txErr != nil {
			onConfirm(nil, fmt.Errorf("failed to send claim transaction: %v", txErr))
			return
		}

		if txReceipt == nil {
			onConfirm(nil, fmt.Errorf("claim transaction not confirmed"))
			return
		}

		if txReceipt.Status != types.ReceiptStatusSuccessful {
			onConfirm(nil, fmt.Errorf("claim transaction failed"))
			return
		}

		totalClaimed := new(big.Int)
		for _, log := range txReceipt.Logs {
			if log.Address == fv.GetVaultAddress() {
				claim, err := fv.GetVaultContract().ParseGrantClaim(*log)
				if err != nil {
					onConfirm(nil, fmt.Errorf("failed to unpack log: %v", err))
					return
				}

				totalClaimed.Add(totalClaimed, claim.Amount)
			}
		}

		onConfirm(totalClaimed, txErr)
	}()

	return nil
}
