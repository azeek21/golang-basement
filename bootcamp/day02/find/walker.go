package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func handleSymlink(path string, prefix string) error {
	res, err := os.Readlink(path)
	newPath := makePathForSymlink(path, res)
	err = filepath.Walk(newPath, func(inner string, info os.FileInfo, err error) error {
		return walker(inner, &info, err, &FindFlags, prefix+path+" -> ")
	})

	if err != nil && FindFlags.SL && os.IsNotExist(err) {
		fmt.Printf("%s -> [broken]\n", path)
		return err
	}

	return nil
}

func walker(path string, info *os.FileInfo, err error, options *SFindFlags, prefix string) error {
	if err != nil {
		return err
	}

	symlink := isSymlink(info)
	file := isFile(info)
	dir := (*info).IsDir()

	shouldPrint := (dir && options.D) || (symlink && options.SL) ||
		(file && options.F && isExtensionMatch(path, options.EXT))

	if shouldPrint {
		fmt.Printf("%s%s\n", prefix, path)
	}

	if isSymlink(info) {
		handleSymlink(path, prefix)
	}

	return nil
}
