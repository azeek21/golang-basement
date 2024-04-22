package main

import (
	"os"
	"strings"
)

func isSymlink(info *os.FileInfo) bool {
	inftoString := (*info).Mode().String()
	inftoString = strings.ToLower(inftoString)
	return strings.HasPrefix(inftoString, "l")
}

func isFile(info *os.FileInfo) bool {
	inftoString := (*info).Mode().String()
	inftoString = strings.ToLower(inftoString)
	return strings.HasPrefix(inftoString, "-")
}

func isExtensionMatch(fileName, extension string) bool {
	fnameSlices := strings.Split(fileName, ".")

	if len(fnameSlices) < 2 {
		return false
	}

	return strings.HasSuffix(
		strings.ToLower(fnameSlices[len(fnameSlices)-1]),
		extension,
	)
}

func makePathForSymlink(path, target string) string {
	separator := string(os.PathSeparator)
	pathSlices := strings.Split(path, separator)
	return strings.Join(pathSlices[:len(pathSlices)-1], separator) + separator + target
}
