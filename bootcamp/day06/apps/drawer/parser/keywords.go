package drawer_parser

import (
	"bytes"
)

type keywords struct {
	END          []byte
	COLORS_START []byte
	IMAGE_START  []byte
	VERSION      []byte
	COMMENT      []byte
	SCALE        []byte
	EQUAL        []byte
}

var KEYWORDS = keywords{
	END:          []byte("end"),
	COLORS_START: []byte("colors"),
	IMAGE_START:  []byte("image"),
	VERSION:      []byte("version"),
	COMMENT:      []byte("#"),
	SCALE:        []byte("scale"),
	EQUAL:        []byte("="),
}

func IsComment(line []byte) bool {
	return bytes.HasPrefix(line, KEYWORDS.COMMENT)
}

func IsEnd(line []byte) bool {
	return bytes.HasPrefix(line, KEYWORDS.END)
}
