package quicksort

import (
	"fmt"
	"grogos/shared"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TTestItem struct {
	desc       string
	arg        []int
	comparator func(a, b int) int
	exp        []int
}

func TestQuickSort(t *testing.T) {
	cases := []TTestItem{
		{
			desc:       "already sorted array (1,2,3) asc",
			comparator: shared.SortIntAsc,
			arg:        []int{1, 2, 3},
			exp:        []int{1, 2, 3},
		},
		{
			desc:       "random (5,0,2,1,3,4) asc",
			comparator: shared.SortIntAsc,
			arg:        []int{5, 0, 2, 1, 3, 4},
			exp:        []int{0, 1, 2, 3, 4, 5},
		},
		{
			desc:       "random (2, 4, 3, -1, -5, 0) asc with negative",
			comparator: shared.SortIntAsc,
			arg:        []int{2, 4, 3, -1, -5, 0},
			exp:        []int{-5, -1, 0, 2, 3, 4},
		},
		{
			desc:       "random (3,1,2,2,5,2,0,4,2) repeated asc",
			comparator: shared.SortIntAsc,
			arg:        []int{3, 1, 2, 2, 5, 2, 0, 4, 2},
			exp:        []int{0, 1, 2, 2, 2, 2, 3, 4, 5},
		},
	}

	for i, tcase := range cases {
		t.Run(fmt.Sprintf("%v: %s\n", i, tcase.desc), func(t *testing.T) {
			QuickSort(tcase.arg, 0, len(tcase.arg)-1)
			assert.Equal(t, tcase.exp, tcase.arg)
		})
	}
}
