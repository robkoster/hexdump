package main

import (
	"os"

	"github.com/robkoster/hexdump"
	"github.com/spf13/cobra"
)

// hexdumpCmd represents the hexdump command
var hexdumpCmd = &cobra.Command{
	Use:   "hexdump filename",
	Short: "Show a hexdump for the specified file",
	Long:  `Display a hexdump for the specified file. Display format can be modified by setting the corresponding flags`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		opts := hexdump.DefaultOptions
		opts.ShowHeader, _ = cmd.Flags().GetBool("display-header")
		opts.ShowAddress, _ = cmd.Flags().GetBool("display-address")
		opts.ShowByteSeparator, _ = cmd.Flags().GetBool("print-byte-separator")
		opts.BytesPerLine, _ = cmd.Flags().GetInt("bytes-per-line")
		opts.AddressBytes, _ = cmd.Flags().GetInt("address-bytes")
		err := hexdump.Dump(args[0], os.Stdout, opts)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	hexdumpCmd.Flags().BoolP("display-header", "d", false, "Show a header above the table")
	hexdumpCmd.Flags().BoolP("display-address", "a", true, "Display address value in the leftmost column")
	hexdumpCmd.Flags().BoolP("print-byte-separator", "s", true, "Print a separator between bytes")
	hexdumpCmd.Flags().IntP("bytes-per-line", "b", 16, "The number of bytes to print on a single line (minimum 4)")
	hexdumpCmd.Flags().Int("address-bytes", 7, "The number of bytes to use for displaying the address (minimum 4)")
}

func Execute() {
	err := hexdumpCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func main() {
	Execute()
}
