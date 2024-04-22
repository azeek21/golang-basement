package snapshots

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type Snapshot struct {
	fileName         string
	fileType         string
	Items            []string
	supportedFormats []string
}

func (snap *Snapshot) parseFileName(fileName string) error {
	nameSlices := strings.Split(fileName, ".")
	if len(nameSlices) < 1 {
		return errors.New("error [input] no file extension")
	}
	snap.fileType = nameSlices[len(nameSlices)-1]
	snap.fileName = fileName
	return nil
}

func (snap *Snapshot) isValid() bool {
	for _, format := range snap.supportedFormats {
		if format == snap.fileType {
			return true
		}
	}
	return false
}

func (snap *Snapshot) Load(fileName string) error {
	snap.supportedFormats = []string{"txt"}
	if err := snap.parseFileName(fileName); err != nil {
		return err
	}
	if !snap.isValid() {
		return errors.New("error [input] invalid input")
	}

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		snap.Items = append(snap.Items, scanner.Text())
	}
	return scanner.Err()
}
