package drawer_parser

import (
	"bufio"
	"errors"
	"io"
)

type ImageSource struct {
	Bytes [][]byte
	W     int
	H     int
}

var ERROR_SRC_IMAGE_BAD_WIDTH = errors.New("Image width must be same through all lines of the source file.")

func NewImageSource() *ImageSource {
	return &ImageSource{}
}

func (imgsrc *ImageSource) ReadTillEndAndParse(src *bufio.Reader) error {
	for {
		line, _, err := src.ReadLine()

		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if IsComment(line) {
			continue
		}
		if IsEnd(line) {
			return nil
		}

		// Initialize image width
		if imgsrc.W == 0 {
			imgsrc.W = len(line)
		}

		// Check all lines for having same width (the must)
		if imgsrc.W != len(line) {
			return ERROR_SRC_IMAGE_BAD_WIDTH
		}

		imgsrc.Bytes = append(imgsrc.Bytes, line)
		imgsrc.H++
	}
	return nil
}
