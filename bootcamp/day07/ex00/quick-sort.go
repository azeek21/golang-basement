package ex00

// Checks if array is sorted
func IsSorted(arr []int) bool {
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

func qsDoer(arr []int, l, h int) {
	if l < h {
		p := arr[h]
		i := l
		for j := l; j < h; j++ {
			if arr[j] < p {
				arr[j], arr[i] = arr[i], arr[j]
				i++
			}
		}
		arr[i], arr[h] = arr[h], arr[i]
		qsDoer(arr, l, i-1) // left
		qsDoer(arr, i+1, h) // right
	}
}

// In place QuickSort based on divide an conquer algorithm
// NOTE: it sorts the slice in place which means arr get's modified.
// use with caution
func QuickSort(arr []int) []int {
	qsDoer(arr, 0, len(arr)-1)
	return arr
}
