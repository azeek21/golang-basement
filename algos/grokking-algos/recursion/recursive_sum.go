package recursion

import (
	"cmp"
	"errors"
	"fmt"
	"grogos/shared"
)

var ERR_NO_EMTPY_ARRAY = errors.New("Array can't be empty")
var ERR_TARGET_NOT_FOUND = errors.New("Item not found")

// EX 4.1 recursive sum
func RecursiveSum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return arr[0]
	}

	return arr[0] + RecursiveSum(arr[1:])
}

// Ex 4.2 recursive count
func RecursiveCount[T comparable](arr []T, tar T) int {
	length := len(arr)

	if length == 0 {
		return 0
	}

	if length == 1 {
		return shared.B2i(arr[0] == tar)
	}

	mid := length / 2
	return RecursiveCount(arr[mid:], tar) + RecursiveCount(arr[:mid], tar)
}

// Ex 4.3
func RecursiveMax[T cmp.Ordered](arr []T) (error, T) {
	l := len(arr)
	if l == 0 {
		return ERR_NO_EMTPY_ARRAY, T(0)
	}
	if l == 1 {
		return nil, arr[0]
	}
	mid := l / 2

	le, left := RecursiveMax(arr[:mid])
	if le != nil {
		return le, left
	}

	re, right := RecursiveMax(arr[mid:])
	if re != nil {
		return re, right
	}

	if left > right {
		return nil, left
	}
	return nil, right
}

// ex 4.4
func RecursiveBinarySearch[T cmp.Ordered](arr []T, tar T, low, high int) int {
	fmt.Printf("l: %v, h: %v, tar: %v\n", low, high, tar)
	if low > high {
		return -1
	}

	mid := low + ((high - low) / 2)
	fmt.Printf("mid: %v\n", mid)
	if arr[mid] == tar {
		return mid
	}

	if tar > arr[mid] {
		return RecursiveBinarySearch(arr, tar, mid+1, high)
	}
	return RecursiveBinarySearch(arr, tar, low, mid-1)
}
