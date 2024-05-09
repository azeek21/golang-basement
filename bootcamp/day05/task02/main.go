package main

import (
	presentheap "day05/task02/present-heap"
	"fmt"
	"log"
)

func main() {
	ps := []presentheap.Present{
		{Value: 5, Size: 1},
		{Value: 4, Size: 5},
		{Value: 3, Size: 1},
		{Value: 5, Size: 2},
	}

	res, err := presentheap.GetNCoolestPresents(ps, 2)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("[")
	for _, p := range res {
		fmt.Printf("(%v, %v)", p.Value, p.Size)
	}
	fmt.Printf("]")
}
