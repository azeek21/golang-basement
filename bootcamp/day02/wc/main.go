package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	args := flag.Args()
	wg := sync.WaitGroup{}
	ch := make(chan [2]string)

	for _, name := range args {
		wg.Add(1)
		go Wc(name, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	for res := range ch {
		fmt.Printf("%s\t%s\n", res[0], res[1])
	}
}
