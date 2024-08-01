package ex02

import (
	"day09/ex00"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Multiplex[T interface{}](chans ...chan T) chan T {
	res := make(chan T)
	wg := sync.WaitGroup{}

	go func() {
		for _, c := range chans {
			wg.Add(1)
			go func(ch chan T) {
				defer wg.Done()
				for val := range ch {
					res <- val
				}
			}(c)
		}
		wg.Wait()
		close(res)
	}()

	return res
}

func shuffle(a []int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

func ShowCaseMultiplex() {
	arr := []int{1, 2, 3, 4}
	mergedOutput := Multiplex(
		ex00.SleepSort(arr, time.Second),
		ex00.SleepSort(arr, time.Millisecond),
		ex00.SleepSort(arr, time.Second),
	)
	for output := range mergedOutput {
		fmt.Println(output)
	}
}
