package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	wg.Add(1)
	data := 1
	go func() {
		mu.Lock()
		data++
		mu.Unlock()
		wg.Done()
	}()

	wg.Wait()
	mu.Lock()
	fmt.Print(data)
	mu.Unlock()
}
