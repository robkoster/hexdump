package main

import (
	"os"

	"github.com/robkoster/hexdump/gohexd"
)

func Execute() {
	cmd := gohexd.InitializeHexDumpCmd()
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func main() {
	Execute()
}
