// doreth/cmds.go

package main

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func getBalance(addressStr string) {
	if !common.IsHexAddress(addressStr) {
		ohNoes("Invalid Ethereum address hex string")
		fmt.Print("\n")
		return
	}
	account := common.HexToAddress(addressStr)

	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		ohNoes(fmt.Sprintf("Error fetching balance: %v", err))
		fmt.Print("\n")
		return
	}

	// Wei (int) to Ether (float)
	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))

	greatSuccess(fmt.Sprintf("Balance: %f ETH", ethValue))
	fmt.Print("\n")
}

func getBlock(blockNumStr string) {
	// Default to latest
	var blockNumber *big.Int

	if blockNumStr != "latest" {
		i := new(big.Int)
		if _, success := i.SetString(blockNumStr, 10); !success {
			ohNoes("Invalid block number")
			fmt.Print("\n")
			return
		}
		blockNumber = i
	}

	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		ohNoes(fmt.Sprintf("Error fetching block: %v", err))
		fmt.Print("\n")
		return
	}

	fmt.Print("\n")
	fmt.Printf("Hash:    %s\n", block.Hash().Hex())
	fmt.Printf("Time:    %s\n", time.Unix(int64(block.Time()), 0))
	fmt.Printf("Txs:     %d\n", len(block.Transactions()))
	fmt.Print("\n")
}

func getGas() {
	price, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		ohNoes(fmt.Sprintf("Error fetching gas price: %v", err))
		fmt.Print("\n")
		return
	}

	// Wei to Gwei
	fPrice := new(big.Float).SetInt(price)
	gwei := new(big.Float).Quo(fPrice, big.NewFloat(math.Pow10(9)))

	greatSuccess(fmt.Sprintf("Current price: %f Gwei", gwei))
	fmt.Print("\n")
}

func getTransaction(hashStr string) {
	txHash := common.HexToHash(hashStr)
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		ohNoes(fmt.Sprintf("Error fetching transaction: %v", err))
		fmt.Print("\n")
		return
	}

	fmt.Print("\n")
	fmt.Printf("Hash:      %s\n", tx.Hash().Hex())
	fmt.Printf("Value:     %s Wei\n", tx.Value().String())
	fmt.Printf("Gas Limit: %d\n", tx.Gas())
	fmt.Printf("Gas Price: %s Wei\n", tx.GasPrice().String())
	fmt.Printf("Pending:   %v\n", isPending)
	fmt.Print("\n")
}
