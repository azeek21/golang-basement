package ex00

import (
	"fmt"
	"sync"
	"time"
)

func sleeAndSendOver(outChannel chan int, n int, unit time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(unit * time.Duration(n))
	outChannel <- n
}

func SleepSort(arr []int, dur time.Duration) chan int {
	ch := make(chan int, 1)
	wg := sync.WaitGroup{}
	go func() {
		defer close(ch)
		for _, n := range arr {
			wg.Add(1)
			go sleeAndSendOver(ch, n, dur, &wg)
		}
		wg.Wait()
	}()
	return ch
}

func Showcase() {
	arr := []int{9, 3, 2, 4, 1, 3, 6, 0, 5, 10, 1000}
	ch := SleepSort(arr, time.Millisecond)
	for n := range ch {
		fmt.Println(n)
	}
}
