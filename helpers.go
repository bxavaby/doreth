// doreth/helpers.go

package main

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	green  = "\033[32m"
	red    = "\033[1;31m"
	blue   = "\033[34m"
	yellow = "\033[1;33m"
	reset  = "\033[0m"
)

func greatSuccess(msg string) {
	fmt.Printf("%s✓ %s%s\n", green, msg, reset)
}

func ohNoNoes(msg string, err error) {
	fmt.Fprintf(os.Stderr, "%s✗ %s %s%s\n", red, msg, err, reset)
}

func ohNoes(msg string) {
	fmt.Fprintf(os.Stderr, "%s✗ %s%s\n", red, msg, reset)
}

func singleWell(msg string) {
	fmt.Printf("%s→ %s%s\n", blue, msg, reset)
}

func tripleWell(msg string) {
	fmt.Printf("%s⚠ %s%s\n", yellow, msg, reset)
}

func Wiper() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
