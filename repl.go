// doreth/repl.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
)

func repl() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	go func() {
		<-sigChan
		fmt.Print("\n")
		singleWell("Exiting...")
		os.Exit(0)
	}()

	fmt.Print("\n")
	greatSuccess("Connected to the Ethereum network.")
	singleWell("Type 'help' for commands, 'exit' to quit.")
	fmt.Print("\n")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" ")

		if !scanner.Scan() {
			fmt.Print("\n")
			singleWell("Goodbye!")
			break
		}

		line := strings.TrimSpace(scanner.Text())
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := strings.ToLower(args[0])

		switch cmd {
		case "help":
			fmt.Println(HelpComm())

		case "exit", "quit":
			singleWell("Goodbye!")
			return

		case "balance":
			if len(args) < 2 {
				tripleWell("Usage: balance <address>")
				fmt.Print("\n")
				continue
			}
			// getBalance(args[1])

		case "block":
			if len(args) < 2 {
				tripleWell("Usage: block <number>")
				fmt.Print("\n")
				continue
			}
			// getBlock(args[1])

		case "gas":
			if len(args) > 1 {
				tripleWell("Usage: gas requires no arguments")
				fmt.Print("\n")
				continue
			}
			// getGas(args[1])

		case "tx":
			if len(args) < 2 {
				tripleWell("Usage: tx <hash>")
				fmt.Print("\n")
				continue
			}
			// getTransaction(args[1])

		default:
			msg := "Unknown command: " + cmd
			ohNoes(msg)
			fmt.Print("\n")
		}
	}
}
