package ex00_test

import (
	"day08/ex00"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TCase struct {
	InArr  []int
	InIdx  int
	Exp    int
	Desc   string
	ExpErr error
}

func Test(t *testing.T) {
	tcases := []TCase{
		{
			Desc:  "Getting first element",
			InArr: []int{1, 2, 3, 4, 5, 6},
			InIdx: 0,
			Exp:   1,
		},
		{
			Desc:  "Getting last element",
			InArr: []int{1, 2, 3, 4, 5, 6},
			InIdx: 5,
			Exp:   6,
		},
		{
			Desc:  "Getting random element",
			InArr: []int{1, 2, 3, 4, 5, 6},
			InIdx: 2,
			Exp:   3,
		},

		{
			Desc:  "Another random element",
			InArr: []int{1, 2, 3, 4, 5, 6},
			InIdx: 4,
			Exp:   5,
		},
		{
			Desc:   "Empty slice",
			InArr:  []int{},
			InIdx:  5,
			Exp:    0,
			ExpErr: ex00.ERR_EMPTY_SLICE,
		},
		{
			Desc:   "Out of bounds index",
			InArr:  []int{1, 2, 3},
			InIdx:  5,
			Exp:    0,
			ExpErr: ex00.ERR_OUT_OF_BOUNDS,
		},
		{
			Desc:   "Invalied index",
			InArr:  []int{1, 2, 3},
			InIdx:  -5,
			Exp:    0,
			ExpErr: ex00.ERR_INVALID_INDEX,
		},
	}

	for i, tCase := range tcases {
		t.Run(fmt.Sprintf("%d--%s\n", i, tCase.Desc), func(t *testing.T) {
			res, err := ex00.GetElement(tCase.InArr, tCase.InIdx)
			assert.Equal(t, tCase.Exp, res)
			assert.Equal(t, tCase.ExpErr, err)
		})
	}
}
