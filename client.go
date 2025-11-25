// doreth/client.go

package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client

func InitClient(url string) {
	var err error
	client, err = ethclient.Dial(url)

	if err != nil {
		ohNoes(fmt.Sprintf("Failed to connect: %v", err))
		os.Exit(1)
	}
}
