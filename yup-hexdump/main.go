package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/hexdump"
)

const (
	flagBytesPerLine = "length"
	flagSkipBytes    = "skip"
	flagReadBytes    = "number"
	flagFormat       = "format"
	flagCanonical    = "canonical"
	flagOctal        = "octal"
	flagDecimal      = "decimal"
	flagHex          = "hex"
	flagUppercase    = "uppercase"
)

func main() {
	app := &cli.App{
		Name:  "hexdump",
		Usage: "display file contents in hexadecimal, decimal, octal, or ascii",
		UsageText: `hexdump [OPTIONS] [FILE...]

   The hexdump utility displays FILE contents in various formats.
   With no FILE, or when FILE is -, read standard input.`,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    flagBytesPerLine,
				Aliases: []string{"l"},
				Usage:   "number of bytes to display per line",
			},
			&cli.IntFlag{
				Name:    flagSkipBytes,
				Aliases: []string{"s"},
				Usage:   "skip first BYTES bytes of each file",
			},
			&cli.IntFlag{
				Name:    flagReadBytes,
				Aliases: []string{"n"},
				Usage:   "interpret only NUM bytes of input",
			},
			&cli.StringFlag{
				Name:    flagFormat,
				Aliases: []string{"e"},
				Usage:   "specify format string",
			},
			&cli.BoolFlag{
				Name:    flagCanonical,
				Aliases: []string{"C"},
				Usage:   "canonical hex+ASCII display",
			},
			&cli.BoolFlag{
				Name:    flagOctal,
				Aliases: []string{"o"},
				Usage:   "two-byte octal display",
			},
			&cli.BoolFlag{
				Name:    flagDecimal,
				Aliases: []string{"d"},
				Usage:   "two-byte decimal display",
			},
			&cli.BoolFlag{
				Name:    flagHex,
				Aliases: []string{"x"},
				Usage:   "two-byte hexadecimal display",
			},
			&cli.BoolFlag{
				Name:    flagUppercase,
				Aliases: []string{"u"},
				Usage:   "use uppercase hex letters",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "hexdump: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add file arguments (or none for stdin)
	for i := 0; i < c.NArg(); i++ {
		params = append(params, gloo.File(c.Args().Get(i)))
	}

	// Add flags based on CLI options
	if c.IsSet(flagBytesPerLine) {
		params = append(params, BytesPerLine(c.Int(flagBytesPerLine)))
	}
	if c.IsSet(flagSkipBytes) {
		params = append(params, SkipBytes(c.Int(flagSkipBytes)))
	}
	if c.IsSet(flagReadBytes) {
		params = append(params, ReadBytes(c.Int(flagReadBytes)))
	}
	if c.IsSet(flagFormat) {
		params = append(params, Format(c.String(flagFormat)))
	}
	if c.Bool(flagCanonical) {
		params = append(params, Canonical)
	}
	if c.Bool(flagOctal) {
		params = append(params, Octal)
	}
	if c.Bool(flagDecimal) {
		params = append(params, Decimal)
	}
	if c.Bool(flagHex) {
		params = append(params, Hex)
	}
	if c.Bool(flagUppercase) {
		params = append(params, Uppercase)
	}

	// Create and execute the hexdump command
	cmd := Hexdump(params...)
	return gloo.Run(cmd)
}
