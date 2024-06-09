package hexdump

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const MinimumAdressBytes = 4
const MinimumBytesPerLine = 1

type HexDumpOptions struct {
	ShowHeader        bool
	ShowAddress       bool
	ShowByteSeparator bool
	BytesPerLine      int
	AddressBytes      int
}

var DefaultHexDumpOptions = HexDumpOptions{
	ShowHeader:        false,
	ShowAddress:       true,
	ShowByteSeparator: true,
	BytesPerLine:      16,
	AddressBytes:      7,
}

// Open the file and write a Hex Dump to the Writer w.
// Returns nil when succesfully written the data
func Dump(filename string, w io.Writer, options ...HexDumpOptions) error {
	// Handle default options and basic sanity check for options
	opts := DefaultHexDumpOptions
	if len(options) > 1 {
		return errors.New("can only specify one set of options")
	}
	if len(options) != 0 {
		opts = options[0]
	}
	if opts.AddressBytes < MinimumAdressBytes {
		return fmt.Errorf("address bytes should be at least %d, given %d", MinimumAdressBytes, opts.AddressBytes)
	}
	if opts.BytesPerLine < MinimumBytesPerLine {
		return fmt.Errorf("address bytes should be at least %d, given %d", MinimumBytesPerLine, opts.BytesPerLine)
	}

	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Show optional Header
	if opts.ShowHeader {
		if opts.ShowAddress {
			fmt.Fprintf(w, "%-*s ", opts.AddressBytes, "Addr")
		}
		fmt.Fprintf(w, "%s\n", "Bytes")

		if opts.ShowAddress {
			addrSeparator := strings.Repeat("=", opts.AddressBytes)
			fmt.Fprintf(w, "%s ", addrSeparator)
		}

		byteChars := opts.BytesPerLine * 2
		if opts.ShowByteSeparator {
			byteChars += opts.BytesPerLine - 1
		}
		bytesSeparator := strings.Repeat("=", byteChars)
		fmt.Fprintf(w, "%s\n", bytesSeparator)
	}

	byteSep := ""
	if opts.ShowByteSeparator {
		byteSep = " "
	}
	// Read one line at a time and write it to the given io.Writer
	buf := make([]byte, opts.BytesPerLine)
	addr := 0
	for {
		n, err := f.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		// Print the address
		if opts.ShowAddress {
			fmt.Fprintf(w, "%0*x ", opts.AddressBytes, addr)
		}
		// Print the bytes
		for i := range n {
			fmt.Fprintf(w, "%02x%s", buf[i], byteSep)
		}
		addr += opts.BytesPerLine
		fmt.Fprintln(w)
	}
	return nil
}
