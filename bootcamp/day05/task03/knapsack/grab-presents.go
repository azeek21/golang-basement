package knapsack

import (
	presentheap "day05/task02/present-heap"
)

type Prize struct {
	Value int
	Size  int
}

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func Knapsack(values, weights []int, capacity, i int) int {
	if i >= len(values) {
		return 0
	} else if weights[i] > capacity {
		return Knapsack(values, weights, capacity, i+1)
	} else if capacity < 0 {
		return MinInt
	} else {
		v1 := values[i] + Knapsack(values, weights, capacity-weights[i], i+1)
		v2 := Knapsack(values, weights, capacity, i+1)
		return max(v1, v2)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Result struct {
	Value int
	Items []presentheap.Present
}

func Knapsack2(presents []presentheap.Present, capacity, i int) Result {
	if i >= len(presents) {
		return Result{Value: 0, Items: []presentheap.Present{}}
	} else if presents[i].Size > capacity {
		return Knapsack2(presents, capacity, i+1)
	} else if capacity < 0 {
		return Result{Value: MinInt, Items: []presentheap.Present{}}
	} else {
		v1 := Knapsack2(presents, capacity-presents[i].Size, i+1)
		v1.Value += presents[i].Value
		v1.Items = append(v1.Items, presents[i])

		v2 := Knapsack2(presents, capacity, i+1)

		if v1.Value > v2.Value {
			return v1
		}
		return v2
	}
}

func GrabPresents(presents []presentheap.Present, capacity int) []presentheap.Present {
	res := []presentheap.Present{}

	return res
}
