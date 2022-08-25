package main

import (
	"log"
	"net/http"
	"sync"
)

var urls = []string{
	"http_//localhost:8000?duration=3s",
	"http_//localhost:8000?duration=1s",
	"http_//localhost:8000?duration=5s",
}

func main() {
	//fetchSequential(urls)
	//fetchConcurrent(urls)
	fetchConcurrentCSP(urls)
}

func fetchSequential(urls []string) {
	for _, url := range urls {
		fetch(url)
	}
}

func fetchConcurrent(urls []string) {
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		go func(u string) {
			fetch(u)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func fetchConcurrentCSP(urls []string) {
	signal := make(chan struct{})
	for _, url := range urls {
		go func(u string) {
			fetch(u)
			signal <- struct{}{}
		}(url)
	}
	<-signal
	<-signal
	<-signal
}

func fetch(url string) {
	resp, err := http.Head(url)
	if err != nil {
		log.Fatalf("failed url: %s, err: %v", url, err)
	}
	log.Println(url, " : ", resp.StatusCode)
}
