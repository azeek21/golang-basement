package binary_search

import (
	"cmp"
	"errors"
)

var ERR_NOT_FOUND = errors.New("Target not found")

func BinarySearch[T cmp.Ordered](sorterArray []T, target T) (int, error) {
	low, high := 0, len(sorterArray)-1

	for low <= high {
		mid := (low + high) / 2
		n := sorterArray[mid]

		if n == target {
			return mid, nil
		}

		if n < target {
			low = mid + 1
		}

		if n > target {
			high = mid - 1
		}
	}

	return 0, ERR_NOT_FOUND
}
