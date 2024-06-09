# hexdump
Simple hexdump library and command line utility written in Go.

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
