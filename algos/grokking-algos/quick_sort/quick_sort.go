package quicksort

import (
	"cmp"
)

// NOTE: array will be sorted in place meaning this function is not pure. If you want immutability pass a copy of your array instead
func QuickSort[T cmp.Ordered](arr []T, low, high int) {
	if low >= high {
		return
	}

	piv := arr[high]
	idx := low - 1

	for i := low; i < high; i++ {
		if arr[i] < piv {
			idx++
			arr[idx], arr[i] = arr[i], arr[idx]
		}
	}

	idx++
	arr[high], arr[idx] = arr[idx], piv

	QuickSort(arr, low, idx-1)

	QuickSort(arr, idx+1, high)
}
