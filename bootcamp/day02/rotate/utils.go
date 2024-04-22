package main

import (
	"os"
	"strings"
)

func isFile(info *os.FileInfo) bool {
	inftoString := (*info).Mode().String()
	inftoString = strings.ToLower(inftoString)
	return strings.HasPrefix(inftoString, "-")
}
