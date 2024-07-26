package main

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

func QuickSort(arr []int) []int {
	qsDoer(arr, 0, len(arr)-1)
	return arr
}
