# hexdump
Simple hexdump library and command line utility __gohexd__ written in Go.

## Library usage
Use the Dump function to write a hexdump to the provided io.Writer. Optionally pass in an __HexDumpOptions__ object with the following options:
|Option|Type|Meaning|
|-|-|-|
|ShowHeader|bool|Print a header above the table|
|ShowAddress|bool|Print hexadecimal addresses in the leftmost column|
|ShowByteSeparator|bool|Print a space between bytes|
|BytesPerLine|int|Number of bytes to print per line|
|AddressBytes|int|Number of bytes to use for displaying the address|

## CLI Usage
Install the latest version of __gohexd__ with:

```go install github.com/robkoster/hexdump/cmd/gohexd@latest```

Run the following command to print a hexdump for any file

```gohexd <filename>```

Optionally you can provide the following options:
|Option|Shortcut|Type|Meaning|
|-|-|-|-|
|--display-header|-d|flag|Print a header above the table|
|--display-address|-a|flag|Print hexadecimal addresses in the leftmost column|
|--print-byte-separator|-s|flag|Print a space between bytes|
|--bytes-per-line|-b|int|Number of bytes to print per line|
|--address-bytes||int|Number of bytes to use for displaying the address|