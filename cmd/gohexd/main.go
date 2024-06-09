package main

import (
	"os"

	"github.com/robkoster/hexdump/gohexd"
)

func Execute() {

	err := gohexd.HexdumpCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func main() {
	Execute()
}
