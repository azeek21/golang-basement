package ex01_test

import (
	"day08/ex01"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches" json:"height" binding:"required"`
}

type TCase struct {
	Inp    interface{}
	Exp    string
	Desc   string
	ExpErr error
}

func Test(t *testing.T) {
	tCases := []TCase{
		{
			Desc: "UnknowlPlat",
			Inp: UnknownPlant{
				FlowerType: "t1",
				LeafType:   "l1",
				Color:      123,
			},
			Exp: "FlowerType:t1\nLeafType:l1\nColor(color_scheme=rgb):123\n",
		},
		{
			Desc: "AnotherUnknownPlant",
			Inp: AnotherUnknownPlant{
				FlowerColor: 432,
				LeafType:    "l2",
				Height:      3,
			},
			Exp: "FlowerColor:432\nLeafType:l2\nHeight(unit=inches json=height binding=required):3\n",
		},
	}

	for i, tcase := range tCases {
		t.Run(fmt.Sprintf("%d:%s\n", i, tcase.Desc), func(t *testing.T) {
			res := ex01.DescribeStruct(tcase.Inp)
			assert.Equal(t, tcase.Exp, res)
		})
	}
}
