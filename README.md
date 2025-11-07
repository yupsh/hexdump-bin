# yup-hexdump

```
NAME:
   hexdump - display file contents in hexadecimal, decimal, octal, or ascii

USAGE:
   hexdump [OPTIONS] [FILE...]

      The hexdump utility displays FILE contents in various formats.
      With no FILE, or when FILE is -, read standard input.

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --length value, -l value  number of bytes to display per line (default: 0)
   --skip value, -s value    skip first BYTES bytes of each file (default: 0)
   --number value, -n value  interpret only NUM bytes of input (default: 0)
   --format value, -e value  specify format string
   --canonical, -C           canonical hex+ASCII display (default: false)
   --octal, -o               two-byte octal display (default: false)
   --decimal, -d             two-byte decimal display (default: false)
   --hex, -x                 two-byte hexadecimal display (default: false)
   --uppercase, -u           use uppercase hex letters (default: false)
   --help, -h                show help
```
