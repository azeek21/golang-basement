package main

import (
	"flag"
)

func init() {
	flag.Parse()
}

func main() {
	args := flag.Args()
	Xargs(args)
}
