package selectionsort

import (
	"cmp"
)

func SortIntAsc(a, b int) int {
	return a - b
}

func SortIntDesc(a, b int) int {
	return b - a
}

func SelectionSort[T cmp.Ordered](arr []T, comparator func(a, b T) int) []T {
	arrLen := len(arr)
	if arrLen == 0 {
		return []T{}
	}

	res := make([]T, len(arr))
	copy(res, arr)
	tmp := res[0]
	for i := range res {
		condidateIdx := i
		for j := i + 1; j < arrLen; j++ {
			cmpRes := comparator(res[condidateIdx], res[j])
			if cmpRes > 0 {
				condidateIdx = j
			}
		}
		tmp = res[condidateIdx]
		res[condidateIdx] = res[i]
		res[i] = tmp
	}
	return res
}
