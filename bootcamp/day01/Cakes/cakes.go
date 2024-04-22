package cakes

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type CakeIngredient struct {
	Name  string `json:"ingredient_name"  xml:"itemname"`
	Count string `json:"ingredient_count" xml:"itemcount"`
	Unit  string `json:"ingredient_unit"  xml:"itemunit"`
}

type Cake struct {
	Name        string           `json:"name"        xml:"name"`
	Time        string           `json:"time"        xml:"stovetime"`
	Ingredients []CakeIngredient `json:"ingredients" xml:"ingredients>item"`
}

type Cakes struct {
	xmlName          xml.Name `xml:"recipes"`
	Items            []Cake   `xml:"cake"    json:"cake"`
	fileName         string
	fileFormat       string
	supportedFormats []string
}

type dbReader interface {
	Load(fileName string) error
	Print()
	parseFileName(fileName string) error
	isFileValid(supportedFormats []string) bool
	loadFromXml() error
	loadFromJson() error
}

var SupportedFormats = []string{"json", "xml"}

func (cakes *Cakes) isFileValid() bool {
	isSupported := false

	for _, format := range cakes.supportedFormats {
		if format == cakes.fileFormat {
			isSupported = true
			break
		}
	}
	return isSupported
}

func (cakes *Cakes) parseFileName(fileName string) error {
	fileNameSlices := strings.Split(fileName, ".")
	if len(fileNameSlices) < 2 {
		return errors.New("error [input]: file doens't have an extension")
	}
	cakes.fileName = fileName
	cakes.fileFormat = fileNameSlices[len(fileNameSlices)-1]
	return nil
}

func (cakes *Cakes) loadFromJson() error {
	file, err := os.Open(cakes.fileName)
	if err != nil {
		return err
	}

	defer file.Close()
	byteArr, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(byteArr, &cakes); err != nil {
		return err
	}

	return nil
}

func (cakes *Cakes) loadFromXml() error {
	file, err := os.Open(cakes.fileName)
	if err != nil {
		return err
	}

	defer file.Close()
	byteArr, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if err = xml.Unmarshal(byteArr, &cakes); err != nil {
		return err
	}

	return nil
}

func (cakes *Cakes) Load(inputFile string) error {
	cakes.supportedFormats = SupportedFormats

	if len(inputFile) < 5 {
		return errors.New(
			"error [input]: filename not provided or bad filename.\nuse -f flag to provide a file",
		)
	}

	err := cakes.parseFileName(inputFile)
	if err != nil {
		return err
	}

	if !cakes.isFileValid() {
		return errors.New("error [input]: not supported format")
	}

	if cakes.fileFormat == "json" {
		err = cakes.loadFromJson()
	} else {
		err = cakes.loadFromXml()
	}

	if err != nil {
		return err
	}

	return nil
}

func (cakes *Cakes) Print() {
	var res []byte
	var err error
	if cakes.fileFormat == "json" {
		res, err = xml.MarshalIndent(cakes.Items, "", "    ")
	} else {
		res, err = json.MarshalIndent(cakes.Items, "", "    ")
	}

	if err != nil {
		fmt.Printf("error [output]: failed to format data")
	} else {
		fmt.Printf("%v\n", string(res))
	}
}
