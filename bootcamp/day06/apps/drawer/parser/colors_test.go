package drawer_parser_test

import (
	"bytes"
	"image/color"
	"testing"

	drawer_parser "github.com/azeek21/blog/apps/drawer/parser"
	"gotest.tools/assert"
)

type ColorTestCase struct {
	name string
	arg  []byte
	exp  any
}

func TestDrawerColors(t *testing.T) {
	colorErrorCases := []ColorTestCase{
		{
			name: "Empty string should error",
			arg:  []byte(""),
			exp:  drawer_parser.BAD_RGBA_ERROR,
		},
		{
			name: "Less than 3 colors should error",
			arg:  []byte("155,10"),
			exp:  drawer_parser.BAD_RGBA_ERROR,
		},
		{
			name: "More than 4 colors should error",
			arg:  []byte("155,10,5,10,1"),
			exp:  drawer_parser.BAD_RGBA_ERROR,
		},
		{
			name: "Bad colro values should error",
			arg:  []byte("-1,256,500"),
			exp:  drawer_parser.BAD_RGBA_ERROR,
		},
	}
	colorSuccessCases := []ColorTestCase{
		{
			name: "With only 3 colors",
			arg:  []byte("100,100,100"),
			exp:  &color.RGBA{100, 100, 100, 255},
		},
		{
			name: "With zero alpha color",
			arg:  []byte("100,100,100,0"),
			exp:  &color.RGBA{100, 100, 100, 0},
		},
		{
			name: "With alpha color",
			arg:  []byte("100,100,100,55"),
			exp:  &color.RGBA{100, 100, 100, 55},
		},
		{
			name: "All Maxed",
			arg:  []byte("255,255,255,255"),
			exp:  &color.RGBA{255, 255, 255, 255},
		},
		{
			name: "All Min",
			arg:  []byte("0,0,0,0"),
			exp:  &color.RGBA{0, 0, 0, 0},
		},
	}
	for _, errCase := range colorErrorCases {
		t.Run(errCase.name, func(t *testing.T) {
			_, err := drawer_parser.ParseRgba(errCase.arg)
			assert.Equal(t, err, errCase.exp)
		})
	}

	for _, sucCase := range colorSuccessCases {
		t.Run(sucCase.name, func(t *testing.T) {
			res, _ := drawer_parser.ParseRgba(sucCase.arg)
			assert.Equal(t, *res, *sucCase.exp.(*color.RGBA))
		})
	}
}

func TestColorMapParseAndAssign(t *testing.T) {
	ftColros := drawer_parser.NewColorMap()
	errCases := []ColorTestCase{
		{
			name: "Bad variable",
			arg:  []byte("x"),
			exp:  drawer_parser.BAD_RGBA_ERROR,
		},
		{
			name: "Bad assign",
			arg:  []byte("x-"),
			exp:  drawer_parser.BAD_RGBA_ERROR,
		},
		{
			name: "Bad variable values",
			arg:  []byte("x=1"),
			exp:  drawer_parser.BAD_RGBA_ERROR,
		},
		{
			name: "Bad variable, no alpha",
			arg:  []byte("x=255,1"),
			exp:  drawer_parser.BAD_RGBA_ERROR,
		},
		{
			name: "Bad variable, overflow (more than 4 colors)",
			arg:  []byte("x=255,1,100,10,1"),
			exp:  drawer_parser.BAD_RGBA_ERROR,
		},
		{
			name: "No variable name",
			arg:  []byte("=255,1,100,10,1"),
			exp:  drawer_parser.BAD_RGBA_ERROR,
		},
	}

	succcesCases := []ColorTestCase{
		{
			name: "With only 3 colors",
			arg:  []byte("x=100,100,100"),
			exp:  &color.RGBA{100, 100, 100, 255},
		},
		{
			name: "With zero alpha color",
			arg:  []byte("y=100,100,100,0"),
			exp:  &color.RGBA{100, 100, 100, 0},
		},
		{
			name: "With alpha color",
			arg:  []byte("z=100,100,100,55"),
			exp:  &color.RGBA{100, 100, 100, 55},
		},
		{
			name: "All Maxed",
			arg:  []byte("b=255,255,255,255"),
			exp:  &color.RGBA{255, 255, 255, 255},
		},
		{
			name: "All Min",
			arg:  []byte("m=0,0,0,0"),
			exp:  &color.RGBA{0, 0, 0, 0},
		},
	}

	for _, errCase := range errCases {
		t.Run(errCase.name, func(t *testing.T) {
			err := ftColros.ParseAndAssign(errCase.arg)
			assert.Equal(t, errCase.exp, err)
		})
	}

	for _, sucCase := range succcesCases {
		t.Run(sucCase.name, func(t *testing.T) {
			err := ftColros.ParseAndAssign(sucCase.arg)
			if err != nil {
				t.Errorf("%v: ParseAndAssign %+v returned error where it shouldn't", sucCase.name, sucCase)
			}
			kv := bytes.Split(sucCase.arg, []byte("="))
			res, _ := drawer_parser.ParseRgba(kv[1])
			assert.Equal(t, *(*ftColros)[string(kv[0])], *res)
		})
	}
}
