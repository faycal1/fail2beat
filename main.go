package main

import (
	"os"

	"github.com/faycal1/fail2beat/cmd"
	_ "github.com/faycal1/fail2beat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
