package main

import (
	presentheap "day05/task02/present-heap"
	"day05/task03/knapsack"
	"fmt"
)

func main() {

	presents := []presentheap.Present{
		{Value: 150, Size: 15},
		{Value: 4000, Size: 35},
		{Value: 10, Size: 3},
		{Value: 50, Size: 4},
		{Value: 120, Size: 12},
		{Value: 100, Size: 10},
		{Value: 200, Size: 17},
	}
	capacity := 34

	bestPrices := knapsack.Knapsack2(presents, capacity, 0)

	fmt.Printf("[Capacity] -> given: %v, optimal value: %v\nPicked presents:\n", capacity, bestPrices.Value)
	for _, item := range bestPrices.Items {
		fmt.Printf("Present: %+v\n", item)
	}
}
