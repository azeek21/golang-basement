package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	args := flag.Args()
	wg := sync.WaitGroup{}
	ch := make(chan []string)
	for _, name := range args {
		wg.Add(1)
		go Rotate(name, &wg, ch)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for res := range ch {
		fmt.Printf("%s\t%s\n", res[1], res[0])
	}
}