// 'doreth connect' enters the ethereum network REPL
// '' becomes the prompt indicator
//
// Inside the REPL, users can run the following:
//
//  balance <address>
// to check account balance
//
//  block <number>
// to fetch block details (defaults to latest)
//
//  tx <hash>
// to get transaction details by hash
//
//  gas
// to print current gas price

package main

import (
	"os"
)

func main() {
	os.Exit(Run())
}
