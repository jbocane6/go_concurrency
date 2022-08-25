package main

import (
	"fmt"
	"sync"
	"time"
)

type account struct {
	name    string
	balance int
}

func transfer(amount int, source, dest *account) {
	if source.balance < amount {
		fmt.Printf("X: %s\n", fmt.Sprintf("%s %s", source, dest))
		return
	}

	time.Sleep(time.Second)

	dest.balance += amount
	source.balance -= amount

	fmt.Printf("ok: %s\n", fmt.Sprintf("%s %s", source, dest))
}

func main() {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	wg.Add(2)
	juan := account{name: "Juan", balance: 500}
	camilo := account{name: "Camilo", balance: 900}

	for _, value := range []int{300, 300} {
		go func(amount int) {
			mu.Lock()
			transfer(amount, &juan, &camilo)
			mu.Unlock()
			wg.Done()
		}(value)
	}
	wg.Wait()
}
