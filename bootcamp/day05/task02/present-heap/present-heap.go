package presentheap

import (
	"container/heap"
	"errors"
)

type Present struct {
	Value int
	Size  int
}

var ErrorCountMustBeNonNegative = errors.New("Count must be a positive integer")
var ErrorCountBiggerThanPresentsLength = errors.New("Count is bigger than the length of presents")

type PresentHeap []Present

func (ph PresentHeap) Len() int {
	return len(ph)
}

func (ph PresentHeap) Less(i, j int) bool {
	if ph[i].Value == ph[j].Value {
		return ph[i].Size < ph[j].Size
	}

	return ph[i].Value > ph[j].Value
}

func (ph PresentHeap) Swap(i, j int) {
	ph[i], ph[j] = ph[j], ph[i]
}

func (ph *PresentHeap) Push(present any) {
	*ph = append(*ph, present.(Present))
}

func (ph *PresentHeap) Pop() any {
	old := *ph
	n := len(old)
	res := old[n-1]
	*ph = old[:n-1]
	return res
}

func NewPresentsHeapFromInitialPresents(initialPrsensts []Present) *PresentHeap {
	ph := PresentHeap{}
	for _, present := range initialPrsensts {
		ph.Push(present)
	}
	return &ph
}

func GetNCoolestPresents(presents []Present, count int) ([]Present, error) {
	res := []Present{}
	ph := NewPresentsHeapFromInitialPresents(presents)
	heap.Init(ph)

	if count < 0 {
		return res, ErrorCountMustBeNonNegative
	}

	if count > ph.Len() {
		return res, ErrorCountBiggerThanPresentsLength
	}

	for i := 0; i < count; i++ {
		res = append(res, heap.Pop(ph).(Present))
	}

	return res, nil
}
