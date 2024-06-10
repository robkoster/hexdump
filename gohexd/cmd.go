package gohexd

import (
	"os"

	"github.com/robkoster/hexdump"
	"github.com/spf13/cobra"
)

// hexdumpCmd represents the hexdump command
var HexdumpCmd = &cobra.Command{
	Use:   "gohexd filename",
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
		opts.StartAddress, _ = cmd.Flags().GetUint64("start-address")
		opts.LimitBytes, _ = cmd.Flags().GetInt64("limit-bytes")
		err := hexdump.Dump(args[0], os.Stdout, opts)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	HexdumpCmd.Flags().BoolP("display-header", "d", false, "Show a header above the table")
	HexdumpCmd.Flags().BoolP("display-address", "a", true, "Display address value in the leftmost column")
	HexdumpCmd.Flags().BoolP("print-byte-separator", "s", true, "Print a separator between bytes")
	HexdumpCmd.Flags().IntP("bytes-per-line", "b", 16, "The number of bytes to print on a single line (minimum 4)")
	HexdumpCmd.Flags().Int("address-bytes", 7, "The number of bytes to use for displaying the address (minimum 4)")
	HexdumpCmd.Flags().Uint64("start-address", 0, "Start printing from the given address")
	HexdumpCmd.Flags().Int64P("limit-bytes", "l", -1, "Limit the number of printed bytes. -1 to print everything")
}
