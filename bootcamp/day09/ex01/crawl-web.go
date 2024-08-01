package ex01

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

func crawlAndSendOver(url string, resReceiver chan *string, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	var errMsg string

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		errMsg = fmt.Sprintf("create request for %s %s", url, err.Error())
		resReceiver <- &errMsg
		return
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		errMsg = fmt.Sprintf("do GET request for %s %s", url, err.Error())
		resReceiver <- &errMsg
		return
	}

	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		errMsg = fmt.Sprintf("read response body of %s %s", url, err.Error())
		resReceiver <- &errMsg
		return
	}

	bodyAsString := string(resBody)

	resReceiver <- &bodyAsString
}

func CrawlWeb(urls chan string, ctx context.Context) chan *string {
	res := make(chan *string)

	wg := sync.WaitGroup{}

	go func() {
		for url := range urls {
			timeoutCtx, _ := context.WithTimeout(ctx, time.Second*10)
			wg.Add(1)
			func(url string) {
				go crawlAndSendOver(url, res, timeoutCtx, &wg)
			}(url)
		}

		wg.Wait()
		close(res)
	}()

	return res
}

func CrawlWebShowcase(urls []string) {
	urlInChan := make(chan string)
	bodyOutChan := make(chan *string)
	ctx, cancelShowcase := context.WithCancel(context.Background())
	sigKillChannel := make(chan os.Signal, 1)
	signal.Notify(sigKillChannel, os.Interrupt)
	bodyOutChan = CrawlWeb(urlInChan, ctx)

	go func() {
		fmt.Println("Listening for interrupt user interrupt")
		for signal := range sigKillChannel {
			fmt.Println("Interrupted. Terminating program...", signal)
			cancelShowcase()
			close(urlInChan)
			close(sigKillChannel)
			time.Sleep(time.Second)
			os.Exit(0)
		}
	}()

	go func() {
		fmt.Println("Starting to send over urls to crawl")
		for _, url := range urls {
			fmt.Println("Sent ", url)
			urlInChan <- url
		}
		fmt.Println("Finished sending all urls.")
		close(urlInChan)
	}()

	for body := range bodyOutChan {
		fmt.Println("RES: ", *body)
	}

	close(sigKillChannel)

	fmt.Println("Finished")
}
