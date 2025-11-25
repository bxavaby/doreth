// doreth/cli.go

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func Logo() string {
	logo := `

doreth  ethereum block explorer
`
	return logo
}

func Version() string {
	version := "doreth  v0.1.0"

	return version
}

func Help() string {
	help := `
Usage: doreth [options]

Options:
  -c, --connect       Connect to the network
                      and enter the REPL
  -h, --help          Display this help message
  -v, --version       Display the version number


Once inside the REPL:

╰┈➤  [commands]

Commands:
  balance <address>   Check account balance
  block <number>      Get block details (default: latest)
  gas                 Print current gas price
  tx <hash>           Get transaction details

`
	return help
}

func HelpComm() string {
	help := `
You are currently in the doreth REPL.
That means you can:

╰┈➤  [commands]

[command]           [fetches]
balance <address>   Account balance
block <number>      Block details (default: latest)
gas                 Current gas price
tx <hash>           Transaction details

`
	return help
}

func Run() int {
	if len(os.Args) < 2 {
		fmt.Println(Help())
		return 0
	}

	if len(os.Args) > 2 {
		ohNoes("Use only one argument at a time!")
	}

	arg := strings.ToLower(os.Args[1])

	switch arg {
	case "-h", "--help", "help":
		fmt.Println(Logo())
		fmt.Println(Help())
		return 0

	case "-v", "--version", "version":
		fmt.Println(Version())
		return 0

	case "-c", "--connect", "connect":
		// With Cloudflare's public mainnet endpoint
		// rpcURL := "https://cloudflare-eth.com"
		//
		// Switched to https://mainnet.infura.io/v3/<api_key>
		// rpcURL := "https://mainnet.infura.io/v3/blablabla"
		//
		// Now reading from .env, because of API lol

		err := godotenv.Load()
		if err != nil {
			tripleWell("Warning: Error loading .env file")
		}

		rpcURL := os.Getenv("URL")
		if rpcURL == "" {
			ohNoes("URL not set. Please check your .env file.")
			return 1
		}

		fmt.Print("\n")
		singleWell("Connecting to Mainnet over HTTPS...")
		InitClient(rpcURL)

		repl()
		return 0

	default:
		Wiper()
		fmt.Printf("Unknown argument: %v\n", os.Args[1])
		fmt.Println(Help())
		return 1
	}
}
