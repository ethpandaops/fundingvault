package fundingvault

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"regexp"
	"slices"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethpandaops/spamoor/txbuilder"
	"github.com/ethpandaops/spamoor/utils"
	"github.com/holiman/uint256"
	"github.com/sirupsen/logrus"

	"github.com/ethpandaops/fundingvault/contract"
)

type FundingVault struct {
	ctx           context.Context
	cancel        context.CancelFunc
	Config        *FundingVaultConfig
	allClients    []*txbuilder.Client
	goodClients   []*txbuilder.Client
	Logger        *logrus.Logger
	chainId       *big.Int
	txpool        *txbuilder.TxPool
	rootWallet    *txbuilder.Wallet
	vaultAddress  common.Address
	vaultContract *contract.Contract
}

func NewFundingVault(ctx context.Context, config *FundingVaultConfig, logger *logrus.Logger) *FundingVault {
	ctx, cancel := context.WithCancel(ctx)
	return &FundingVault{
		ctx:    ctx,
		cancel: cancel,
		Config: config,
		Logger: logger,
	}
}

func (fv *FundingVault) Initialize() error {
	// prepare clients
	err := fv.prepareClients()
	if err != nil {
		return err
	}

	// prepare txpool
	fv.txpool = txbuilder.NewTxPool(&txbuilder.TxPoolOptions{
		GetClientFn: func(index int, random bool) *txbuilder.Client {
			return fv.GetClient(index, random)
		},
		GetClientCountFn: func() int {
			return len(fv.goodClients)
		},
	})

	// load root wallet
	if fv.Config.Privkey == "" {
		return fmt.Errorf("privkey is not set")
	}
	rootWallet, err := txbuilder.NewWallet(fv.Config.Privkey)
	if err != nil {
		return err
	}
	fv.rootWallet = rootWallet

	client := fv.GetClient(0, true)
	err = client.UpdateWallet(fv.ctx, fv.rootWallet)
	if err != nil {
		return err
	}

	fv.Logger.Infof(
		"initialized root wallet (addr: %v balance: %v ETH, nonce: %v)",
		rootWallet.GetAddress().String(),
		utils.WeiToEther(uint256.MustFromBig(rootWallet.GetBalance())).Uint64(),
		rootWallet.GetNonce(),
	)

	// load funding vault contract
	fv.vaultAddress = common.HexToAddress(fv.Config.FundingVaultAddress)
	vaultContract, err := contract.NewContract(fv.vaultAddress, client.GetEthClient())
	if err != nil {
		return err
	}

	_, err = vaultContract.GetVaultToken(&bind.CallOpts{
		From:    fv.rootWallet.GetAddress(),
		Context: fv.ctx,
	})
	if err != nil {
		return fmt.Errorf("failed to check vault contract. Is the FundingVault deployed at %v?", fv.Config.FundingVaultAddress)
	}

	fv.vaultContract = vaultContract
	return nil
}

func (fv *FundingVault) Shutdown() {
	fv.cancel()
}

func (fv *FundingVault) GetTxPool() *txbuilder.TxPool {
	return fv.txpool
}

func (fv *FundingVault) GetRootWallet() *txbuilder.Wallet {
	return fv.rootWallet
}

func (fv *FundingVault) GetVaultContract() *contract.Contract {
	return fv.vaultContract
}

func (fv *FundingVault) GetVaultAddress() common.Address {
	return fv.vaultAddress
}

func (fv *FundingVault) GetClaimableBalance(ctx context.Context) (*big.Int, error) {
	balance, err := fv.vaultContract.GetClaimableBalance0(&bind.CallOpts{
		From:    fv.rootWallet.GetAddress(),
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return balance, nil
}

func (fv *FundingVault) ParseEthAmount(amount string) (*big.Int, error) {
	if amount == "" {
		return big.NewInt(0), nil
	}

	re := regexp.MustCompile(`^(\d+)(?:\.(\d{1,18}))?(?: *([a-zA-Z]+))?$`)
	matches := re.FindStringSubmatch(amount)
	if len(matches) != 4 {
		return nil, fmt.Errorf("invalid amount: %v", amount)
	}

	value := new(big.Int).SetUint64(0)
	switch strings.ToLower(matches[3]) {
	case "gwei":
		value.SetString(matches[1], 10)
		value.Mul(value, big.NewInt(1000000000))
		if len(matches[2]) > 0 {
			decimals := len(matches[2])
			wei := new(big.Int).SetUint64(0)
			wei.SetString(matches[2], 10)
			if decimals < 9 {
				wei.Mul(wei, big.NewInt(0).SetUint64(1000000000/uint64(math.Pow(10, float64(9-decimals)))))
			}
			value.Add(value, wei)
		}
	case "eth":
		value.SetString(matches[1], 10)
		value.Mul(value, big.NewInt(1000000000))
		value.Mul(value, big.NewInt(1000000000))
		if len(matches[2]) > 0 {
			decimals := len(matches[2])
			gwei := new(big.Int).SetUint64(0)
			gwei.SetString(matches[2], 10)
			if decimals < 9 {
				gwei.Mul(gwei, big.NewInt(0).SetUint64(1000000000/uint64(math.Pow(10, float64(9-decimals)))))
			}
			value.Add(value, gwei.Mul(gwei, big.NewInt(1000000000)))
		}
	default:
		value.SetString(matches[1], 10)
	}

	return value, nil
}

func (fv *FundingVault) CheckAndRefillWallets(ctx context.Context, onSubmit func(request FundingRequest, tx *types.Transaction)) ([]FundingRequest, error) {
	requests, err := fv.CheckWallets(ctx)
	if err != nil {
		return nil, err
	}

	err = fv.CheckRequests(ctx, requests)
	if err != nil {
		return nil, err
	}

	requests, err = fv.ExecuteFunding(ctx, requests, onSubmit)
	if err != nil {
		return nil, err
	}

	return requests, nil
}

type FundingRequest struct {
	Address   common.Address
	Priority  int
	Balance   *big.Int
	Request   *big.Int
	MinAmount *big.Int
	Error     error
}

func (fv *FundingVault) CheckWallets(ctx context.Context) ([]FundingRequest, error) {
	results := make([]FundingRequest, 0)

	client := fv.GetClient(0, true)
	if client == nil {
		return nil, fmt.Errorf("no client available")
	}

	minClaimAmount, err := fv.ParseEthAmount(fv.Config.MinClaimAmount)
	if err != nil {
		return nil, err
	}

	// Check root wallet if configured
	if fv.Config.RefillRootWallet != nil {
		rootWallet := fv.GetRootWallet()

		walletBalance := rootWallet.GetBalance()
		minBalance, err := fv.ParseEthAmount(fv.Config.RefillRootWallet.MinBalance)
		if err != nil {
			return nil, err
		}

		var refillAmount *big.Int
		if walletBalance.Cmp(minBalance) < 0 {
			refillAmount, err = fv.ParseEthAmount(fv.Config.RefillRootWallet.RefillAmount)
			if err != nil {
				return nil, err
			}
			if refillAmount.Cmp(big.NewInt(0)) == 0 {
				refillAmount = new(big.Int).Sub(minBalance, walletBalance)
			}
		} else {
			refillAmount = big.NewInt(0)
		}

		if refillAmount.Cmp(minClaimAmount) < 0 {
			refillAmount = big.NewInt(0)
		}

		results = append(results, FundingRequest{
			Address:   fv.rootWallet.GetAddress(),
			Priority:  fv.Config.RefillRootWallet.Priority,
			Balance:   walletBalance,
			Request:   refillAmount,
			MinAmount: minBalance,
		})
	}

	// Check configured addresses
	for _, refillConfig := range fv.Config.RefillAddresses {
		address := common.HexToAddress(refillConfig.Address)

		balance, err := client.GetBalanceAt(fv.ctx, address)
		if err != nil {
			results = append(results, FundingRequest{
				Address: address,
				Request: nil,
				Error:   err,
			})
			continue
		}

		minBalance, err := fv.ParseEthAmount(refillConfig.MinBalance)
		if err != nil {
			return nil, err
		}

		var refillAmount *big.Int
		if balance.Cmp(minBalance) < 0 {
			refillAmount, err = fv.ParseEthAmount(refillConfig.RefillAmount)
			if err != nil {
				return nil, err
			}
			if refillAmount.Cmp(minBalance) == 0 {
				refillAmount = new(big.Int).Sub(minBalance, balance)
			}
		} else {
			refillAmount = big.NewInt(0)
		}

		if refillAmount.Cmp(minClaimAmount) < 0 {
			refillAmount = big.NewInt(0)
		}

		results = append(results, FundingRequest{
			Address:   address,
			Priority:  refillConfig.Priority,
			Balance:   balance,
			Request:   refillAmount,
			MinAmount: minBalance,
		})
	}

	return results, nil
}

func (fv *FundingVault) CheckRequests(ctx context.Context, requests []FundingRequest) error {
	claimable, err := fv.GetClaimableBalance(ctx)
	if err != nil {
		return err
	}

	totalRequest := big.NewInt(0)
	for _, request := range requests {
		totalRequest.Add(totalRequest, request.Request)
	}

	if totalRequest.Cmp(claimable) > 0 {
		// total request is greater than claimable balance
		slices.SortFunc(requests, func(a, b FundingRequest) int {
			return a.Priority - b.Priority
		})

		availableBalance := big.NewInt(0).Set(claimable)
		priorityGroup := 0
		priorityGroupBalance := big.NewInt(0)
		priorityGroupRequests := []*FundingRequest{}
		processPriorityGroup := func() {
			defer func() {
				priorityGroupBalance = big.NewInt(0)
				priorityGroupRequests = []*FundingRequest{}
			}()

			if priorityGroupBalance.Cmp(big.NewInt(0)) == 0 || len(priorityGroupRequests) == 0 {
				return
			}

			if priorityGroupBalance.Cmp(availableBalance) > 0 {
				totalRequest := big.NewInt(0)
				for _, request := range priorityGroupRequests {
					request.Request.Div(request.Request, priorityGroupBalance)
					request.Request.Mul(request.Request, availableBalance)
					totalRequest.Add(totalRequest, request.Request)
				}

				if totalRequest.Cmp(availableBalance) > 0 {
					// rounding issue, just deduct from last request
					priorityGroupRequests[len(priorityGroupRequests)-1].Request.Sub(priorityGroupRequests[len(priorityGroupRequests)-1].Request, totalRequest)
				}

				availableBalance.SetUint64(0)
			} else {
				availableBalance.Sub(availableBalance, priorityGroupBalance)
			}

			availableBalance.Add(availableBalance, priorityGroupBalance)
			priorityGroupBalance = big.NewInt(0)
			priorityGroupRequests = []*FundingRequest{}
		}

		for _, request := range requests {
			if request.Priority != priorityGroup {
				processPriorityGroup()
				priorityGroup = request.Priority
			}

			priorityGroupRequests = append(priorityGroupRequests, &request)
			priorityGroupBalance.Add(priorityGroupBalance, request.Request)
		}

		processPriorityGroup()
	}

	return nil
}

func (fv *FundingVault) ExecuteFunding(ctx context.Context, requests []FundingRequest, onSubmit func(request FundingRequest, tx *types.Transaction)) ([]FundingRequest, error) {
	results := make([]FundingRequest, 0, len(requests))
	var wg sync.WaitGroup
	resultChan := make(chan FundingRequest, len(requests)+1)

	for _, request := range requests {
		if request.Request.Cmp(big.NewInt(0)) == 0 {
			resultChan <- request
			continue
		}

		wg.Add(1)
		go func(request FundingRequest) {
			defer wg.Done()

			if request.Error != nil {
				resultChan <- request
				return
			}

			claimed, err := fv.ClaimSync(ctx, request.Request, func(tx *types.Transaction) {
				if onSubmit != nil {
					onSubmit(request, tx)
				}
			})
			resultChan <- FundingRequest{
				Address:   request.Address,
				Balance:   request.Balance,
				Request:   claimed,
				MinAmount: request.MinAmount,
				Error:     err,
			}
		}(request)
	}

	wg.Wait()
	close(resultChan)

	for result := range resultChan {
		results = append(results, result)
	}

	return results, nil
}
