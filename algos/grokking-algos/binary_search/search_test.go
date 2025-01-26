package binary_search

import (
	"testing"
)

type SearchRes struct {
	res int
	err error
}

type SearchArgs struct {
	arr []int
	tar int
}

type Case struct {
	args     []SearchArgs
	expected []SearchRes
}

func TestBinarySearch(t *testing.T) {
	binaryCases := Case{
		args: []SearchArgs{
			{
				arr: []int{1, 2, 3, 4, 5, 6},
				tar: 4,
			},
			{
				arr: []int{3, 4, 5, 6, 7},
				tar: 4,
			},
			{
				arr: []int{-1, 0, 9, 99, 999, 9999},
				tar: 999,
			},
			{
				arr: []int{1, 2, 3, 5},
				tar: 999,
			},
			{
				arr: []int{0, 1, 5, 78, 99},
				tar: 999,
			},
		},
		expected: []SearchRes{
			{
				res: 3,
			},
			{
				res: 1,
			},
			{
				res: 4,
			},
			{
				res: 0,
				err: ERR_NOT_FOUND,
			},
			{
				res: 0,
				err: ERR_NOT_FOUND,
			},
		},
	}

	for i, arg := range binaryCases.args {
		funcRes, funcErr := BinarySearch(arg.arr, arg.tar)
		expected := binaryCases.expected[i]

		if funcErr != nil {
			if funcErr != expected.err {
				t.Fatalf("Binary search returned unexpected error")
			}
		}

		if funcRes != expected.res {
			t.Fatalf("Binary search returned unexpected result. expected: %d\tactually returned: %d", expected.res, funcRes)
		}

		if expected.err != nil && funcErr == nil {
			t.Fatalf("Binary search was expected to return error but didn't")
		}
	}
}
