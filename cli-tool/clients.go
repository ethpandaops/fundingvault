package fundingvault

import (
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethpandaops/spamoor/txbuilder"
	"golang.org/x/exp/rand"
)

func (fv *FundingVault) GetClient(input int, random bool) *txbuilder.Client {
	if len(fv.goodClients) == 0 {
		return nil
	}

	if random {
		input = rand.Intn(len(fv.goodClients))
	}

	input = input % len(fv.goodClients)
	return fv.goodClients[input]
}

func (fv *FundingVault) prepareClients() error {
	fv.allClients = make([]*txbuilder.Client, 0)
	wg := &sync.WaitGroup{}
	mtx := sync.Mutex{}

	var chainId *big.Int
	for _, rpcHost := range fv.Config.RpcHosts {
		wg.Add(1)

		go func(rpcHost string) {
			defer wg.Done()

			client, err := txbuilder.NewClient(rpcHost)
			if err != nil {
				fv.Logger.Errorf("failed creating client for '%v': %v", client.GetRPCHost(), err.Error())
				return
			}
			client.Timeout = 10 * time.Second
			cliChainId, err := client.GetChainId(fv.ctx)
			if err != nil {
				fv.Logger.Errorf("failed getting chainid from '%v': %v", client.GetRPCHost(), err.Error())
				return
			}
			if chainId == nil {
				chainId = cliChainId
			} else if cliChainId.Cmp(chainId) != 0 {
				fv.Logger.Errorf("chainid missmatch from %v (chain ids: %v, %v)", client.GetRPCHost(), cliChainId, chainId)
				return
			}
			client.Timeout = 30 * time.Second
			mtx.Lock()
			fv.allClients = append(fv.allClients, client)
			mtx.Unlock()
		}(rpcHost)
	}

	wg.Wait()
	fv.chainId = chainId
	if len(fv.allClients) == 0 {
		return fmt.Errorf("no useable clients")
	}

	err := fv.checkClientStatus()
	if err != nil {
		return err
	}

	go fv.watchClientStatusLoop()
	return nil
}

func (fv *FundingVault) watchClientStatusLoop() {
	sleepTime := 2 * time.Minute
	for {
		select {
		case <-fv.ctx.Done():
			return
		default:
			time.Sleep(sleepTime)
		}

		err := fv.checkClientStatus()
		if err != nil {
			fv.Logger.Warnf("could not check client status: %v", err)
			sleepTime = 10 * time.Second
		} else {
			sleepTime = 2 * time.Minute
		}
	}
}

func (fv *FundingVault) checkClientStatus() error {
	wg := &sync.WaitGroup{}
	mtx := sync.Mutex{}
	clientHeads := map[int]uint64{}
	highestHead := uint64(0)

	for idx, client := range fv.allClients {
		wg.Add(1)
		go func(idx int, client *txbuilder.Client) {
			defer wg.Done()

			blockHeight, err := client.GetBlockHeight(fv.ctx)
			if err != nil {
				fv.Logger.Warnf("client check failed: %v", err)
			} else {
				mtx.Lock()
				clientHeads[idx] = blockHeight
				if blockHeight > highestHead {
					highestHead = blockHeight
				}
				mtx.Unlock()
			}
		}(idx, client)
	}
	wg.Wait()

	goodClients := make([]*txbuilder.Client, 0)
	goodHead := highestHead
	if goodHead > 2 {
		goodHead -= 2
	}
	for idx, client := range fv.allClients {
		if clientHeads[idx] >= goodHead {
			goodClients = append(goodClients, client)
		}
	}
	fv.goodClients = goodClients
	fv.Logger.Infof("client check completed (%v good clients, %v bad clients)", len(goodClients), len(fv.allClients)-len(goodClients))

	return nil
}
