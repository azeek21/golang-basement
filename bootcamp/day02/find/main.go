package main

import (
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk(FindFlags.DIR, func(path string, info os.FileInfo, err error) error {
		return walker(path, &info, err, &FindFlags, "")
	})
}
