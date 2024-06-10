package hexdump

import (
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

const (
	MinimumAdressBytes  = 4
	MinimumBytesPerLine = 1
)

type HexDumpOptions struct {
	ShowHeader        bool
	ShowAddress       bool
	ShowByteSeparator bool
	BytesPerLine      int
	AddressBytes      int
	StartAddress      uint64
	LimitBytes        int64
}

var DefaultOptions = HexDumpOptions{
	ShowHeader:        false,
	ShowAddress:       true,
	ShowByteSeparator: true,
	BytesPerLine:      16,
	AddressBytes:      7,
	StartAddress:      0,
	LimitBytes:        -1,
}

// Open the file and write a Hex Dump to the Writer w.
// Returns nil when succesfully written the data
func Dump(filename string, w io.Writer, options ...HexDumpOptions) error {
	// Handle default options and basic sanity check for options
	opts := DefaultOptions
	if len(options) > 1 {
		return errors.New("can only specify one set of options")
	}
	if len(options) != 0 {
		opts = options[0]
	}
	if opts.LimitBytes == -1 {
		opts.LimitBytes = math.MaxInt64
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

	// Skip upto start address
	startAddr, err := f.Seek(int64(opts.StartAddress), 0)
	if err != nil {
		return err
	}

	// Read one line at a time and write it to the given io.Writer
	buf := make([]byte, opts.BytesPerLine)
	addr := startAddr
	for {
		n, err := f.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		// Adjust the end based on the maximum number of bytes to read
		bytesRead := addr - startAddr
		diff := int(bytesRead + int64(n) - opts.LimitBytes)
		if diff > 0 {
			n -= diff
		}
		// Nothing left, we can immediately stop
		if n <= 0 {
			break
		}

		// Print the address
		if opts.ShowAddress {
			fmt.Fprintf(w, "%0*x ", opts.AddressBytes, addr)
		}
		// Print the bytes
		for i := range n {
			fmt.Fprintf(w, "%02x%s", buf[i], byteSep)
			addr++
		}
		fmt.Fprintln(w)
	}
	return nil
}
