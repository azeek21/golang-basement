package main

import (
	"compress/gzip"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type ROTATE_OPTIONS struct {
	HasOut bool
	Out    string
}

var RotateOptions ROTATE_OPTIONS

func isFileOk(path string) bool {
	info, err := os.Lstat(path)
	if err != nil || info.IsDir() || !isFile(&info) {
		return false
	}
	return true
}

func createTargetFileName(path string) (string, error) {
	pathSlices := strings.Split(path, string(os.PathSeparator))
	fname := pathSlices[len(pathSlices)-1]
	stat, err := os.Lstat(path)
	if err != nil {
		return "", nil
	}
	ftime := stat.ModTime().Unix()
	fnameSlices := strings.Split(fname, ".")
	name, ext := fnameSlices[0], fnameSlices[len(fnameSlices)-1]
	res := fmt.Sprintf("%s_%d.%s.tar.gz", name, ftime, ext)
	return res, nil
}

func createTarget(path string) (*os.File, error) {
	tar := ""
	if RotateOptions.HasOut {
		outInfo, err := os.Lstat(RotateOptions.Out)
		if err != nil {
			return nil, errors.New("error [options]: failed to access output directory")
		}
		if !outInfo.IsDir() {
			return nil, errors.New("error [options]: output target is not a directory")
		}
		newFname, err := createTargetFileName(path)
		if err != nil {
			return nil, err
		}
		tar = RotateOptions.Out + string(os.PathSeparator) + newFname
	} else {
		newFname, err := createTargetFileName(path)
		if err != nil {
			return nil, err
		}
		pathSlices := strings.Split(path, string(os.PathSeparator))
		tar = fmt.Sprintf("%s%s%s", strings.Join(pathSlices[:len(pathSlices)-1], string(os.PathSeparator)), string(os.PathSeparator), newFname)
	}
	res, err := os.Create(tar)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func Rotate(path string, wg *sync.WaitGroup, ch chan []string) {
	defer wg.Done()
	if !isFileOk(path) {
		ch <- []string{path, "error [input]: input file invalid"}
		return
	}
	var err error
	srcFd, err := os.Open(path)
	if err != nil {
		ch <- []string{path, err.Error()}
		return
	}
	defer srcFd.Close()
	tarFd, err := createTarget(path)
	if err != nil {
		ch <- []string{path, err.Error()}
		return
	}
	defer tarFd.Close()
	archWriter := gzip.NewWriter(tarFd)
	buf := make([]byte, 16*1024)
	for {
		n, err := srcFd.Read(buf)
		if n > 0 {
			archWriter.Write(buf[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
	}
	archWriter.Close()
	if err != nil {
		ch <- []string{path, err.Error()}
		return
	}
	ch <- []string{path, "OK"}
}

func init() {
	outDir := flag.String(
		"a",
		"",
		"output directory for archived file. Default: same director as source file",
	)
	h, help := flag.Bool(
		"h",
		false,
		"show help message",
	), flag.Bool(
		"help",
		false,
		"show help message",
	)
	flag.Parse()

	if *h || *help {
		fmt.Printf("Archives passed files.\nUsage: ./rotate [-a /target/dir/] /path/to/file.txt /other/file.txt\n")
		flag.Usage()
		os.Exit(0)
	}

	RotateOptions = ROTATE_OPTIONS{Out: *outDir, HasOut: len(*outDir) > 0}
}
