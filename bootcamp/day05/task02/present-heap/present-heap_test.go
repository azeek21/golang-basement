package presentheap_test

import (
	presentheap "day05/task02/present-heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	name   string
	arg1   []presentheap.Present
	arg2   int
	expect any
}

func TestGetNCooestPresents(t *testing.T) {
	testSuccessCases := []TestCase{
		{
			name:   "empty presents in and 0 count, must return empty slice",
			arg1:   []presentheap.Present{},
			arg2:   0,
			expect: []presentheap.Present{},
		},
		{
			name: "Should return coolest elements when no value overlap",
			arg1: []presentheap.Present{
				{Value: 9, Size: 1},
				{Value: 2, Size: 5},
				{Value: 3, Size: 1},
				{Value: 6, Size: 2},
				{Value: 5, Size: 1},
				{Value: 4, Size: 10},
			},
			arg2: 4,
			expect: []presentheap.Present{
				{Value: 9, Size: 1},
				{Value: 6, Size: 2},
				{Value: 5, Size: 1},
				{Value: 4, Size: 10},
			},
		},
		{
			name: "Should return coolest elements value overlab, should sort by size",
			arg1: []presentheap.Present{
				{Value: 9, Size: 1},
				{Value: 5, Size: 2},
				{Value: 9, Size: 3},
				{Value: 4, Size: 10},
				{Value: 9, Size: 2},
				{Value: 5, Size: 1},
			},
			arg2: 5,
			expect: []presentheap.Present{
				{Value: 9, Size: 1},
				{Value: 9, Size: 2},
				{Value: 9, Size: 3},
				{Value: 5, Size: 1},
				{Value: 5, Size: 2},
			},
		},
	}

	testErrorCases := []TestCase{
		{
			name:   "Given a count bigger than the length of presents should return error",
			arg1:   []presentheap.Present{},
			arg2:   1,
			expect: presentheap.ErrorCountBiggerThanPresentsLength,
		},
		{
			name:   "Given a count bigger than the length of presents should return error",
			arg1:   []presentheap.Present{},
			arg2:   -1,
			expect: presentheap.ErrorCountMustBeNonNegative,
		},
		{
			name: "Given correct answers, error should be nil",
			arg1: []presentheap.Present{
				{Value: 9, Size: 3},
				{Value: 4, Size: 10},
				{Value: 9, Size: 2},
			},
			arg2:   2,
			expect: nil,
		},
	}

	for _, testCase := range testSuccessCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, _ := presentheap.GetNCoolestPresents(testCase.arg1, testCase.arg2)
			assert.Equal(t, testCase.expect, res)
		})
	}
	for _, testCase := range testErrorCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := presentheap.GetNCoolestPresents(testCase.arg1, testCase.arg2)
			assert.Equal(t, testCase.expect, err)
		})
	}
}
