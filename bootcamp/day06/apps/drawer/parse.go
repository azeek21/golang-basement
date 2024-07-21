package drawer

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"strconv"

	drawer_parser "github.com/azeek21/blog/apps/drawer/parser"
)

var SYNTAX_ERROR_SCALE = errors.New("Error occoured while parsing the scale variable. Please make sure to check out the docs and example.")

func openSourceImage(input string) (*os.File, error) {
	fid, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	return fid, nil
}

func (d *Drawer) ParseFile(filName string) error {
	srcFile, err := openSourceImage(filName)

	if err != nil {
		return err
	}
	defer srcFile.Close()

	src := bufio.NewReader(srcFile)
	for {
		line, _, err := src.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if bytes.HasPrefix(line, drawer_parser.KEYWORDS.COLORS_START) {
			if err := d.Colors.ReadTillEndAndParse(src); err != nil {
				return err
			}
		}

		if bytes.HasPrefix(line, drawer_parser.KEYWORDS.IMAGE_START) {
			err := d.Src.ReadTillEndAndParse(src)
			if err != nil {
				return err
			}
		}

		if bytes.HasPrefix(line, drawer_parser.KEYWORDS.SCALE) {
			scale := bytes.Split(line, drawer_parser.KEYWORDS.EQUAL)
			if len(scale) != 2 {
				return SYNTAX_ERROR_SCALE
			}
			d.Scale, err = strconv.Atoi(string(scale[1]))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
