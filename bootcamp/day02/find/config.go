package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	SL   = "sl"
	F    = "f"
	D    = "d"
	EXT  = "ext"
	HELP = "help"
	H    = "h"
)

type SFindFlags struct {
	SL  bool
	F   bool
	D   bool
	EXT string
	DIR string
}

var FindFlags SFindFlags

func init() {
	sl := flag.Bool(SL, false, "show only symlinks\ne.g: ./find ./foo -sm")
	f := flag.Bool(F, false, "show only files\ne.g: ./find ./foo -f")
	d := flag.Bool(D, false, "show only directories\ne.g: ./find ./foo -d")
	ext := flag.String(
		EXT,
		"",
		"find files with specified extension\ne.g: ./find /foo -ext go  #shows all *.go files in /foo directory",
	)
	help := flag.Bool(HELP, false, "show help message")
	h := flag.Bool(H, false, "show help message")
	flag.Parse()
	args := flag.Args()

	if *h || *help {
		fmt.Printf("Usage: command path -optionsn\ne.g: ./find /some/path\n")
		flag.Usage()
		os.Exit(0)
	}

	if len(args) != 1 {
		fmt.Printf("error: missing path argument")
		os.Exit(1)
	}

	if *sl || *d || *f {
		FindFlags = SFindFlags{SL: *sl, F: *f, D: *d}
	} else {
		FindFlags = SFindFlags{true, true, true, "", ""}
	}

	FindFlags.DIR = args[0]

	if len(*ext) > 0 {
		if FindFlags.F && !FindFlags.D && !FindFlags.SL {
			FindFlags.EXT = *ext
		} else {
			fmt.Printf("error [options]: -ext can only be used with -f option to print files with a specific extension\n")
			os.Exit(1)
		}
	}
}
