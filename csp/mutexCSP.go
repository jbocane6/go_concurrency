package main

import (
	"fmt"
	"time"
)

type account struct {
	name    string
	balance int
}

type bankOperation struct {
	amount int
	done   chan struct{}
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
	signal := make(chan struct{})
	transaction := make(chan *bankOperation)

	juan := account{name: "Juan", balance: 500}
	camilo := account{name: "Camilo", balance: 900}

	go func() {
		for {
			request := <-transaction
			transfer(request.amount, &juan, &camilo)
			request.done <- struct{}{}
		}
	}()

	for _, value := range []int{300, 300} {
		go func(amount int) {
			requestTransaction := bankOperation{amount: amount, done: make(chan struct{})}
			transaction <- &requestTransaction

			signal <- <-requestTransaction.done
		}(value)
	}
	<-signal
	<-signal
}
