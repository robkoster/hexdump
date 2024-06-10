package gohexd

import (
	"os"

	"github.com/robkoster/hexdump"
	"github.com/spf13/cobra"
)

func InitializeHexDumpCmd() *cobra.Command {
	var hexdumpCmd = &cobra.Command{
		Use:     "gohexd filename",
		Version: "v0.2.2",
		Short:   "Show a hexdump for the specified file",
		Long:    `Display a hexdump for the specified file. Display format can be modified by setting the corresponding flags`,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts := hexdump.DefaultOptions
			opts.ShowHeader, _ = cmd.Flags().GetBool("display-header")
			opts.ShowAddress, _ = cmd.Flags().GetBool("display-address")
			opts.ShowByteSeparator, _ = cmd.Flags().GetBool("print-byte-separator")
			opts.BytesPerLine, _ = cmd.Flags().GetInt("bytes-per-line")
			opts.AddressBytes, _ = cmd.Flags().GetInt("address-bytes")
			opts.StartAddress, _ = cmd.Flags().GetUint64("start-address")
			opts.LimitBytes, _ = cmd.Flags().GetInt64("limit-bytes")

			if err := hexdump.Dump(args[0], os.Stdout, opts); err != nil {
				return err
			}
			return nil
		},
	}

	hexdumpCmd.Flags().BoolP("display-header", "d", false, "Show a header above the table")
	hexdumpCmd.Flags().BoolP("display-address", "a", true, "Display address value in the leftmost column")
	hexdumpCmd.Flags().BoolP("print-byte-separator", "s", true, "Print a separator between bytes")
	hexdumpCmd.Flags().IntP("bytes-per-line", "b", 16, "The number of bytes to print on a single line (minimum 4)")
	hexdumpCmd.Flags().Int("address-bytes", 7, "The number of bytes to use for displaying the address (minimum 4)")
	hexdumpCmd.Flags().Uint64("start-address", 0, "Start printing from the given address")
	hexdumpCmd.Flags().Int64P("limit-bytes", "l", -1, "Limit the number of printed bytes. -1 to print everything")

	return hexdumpCmd
}
