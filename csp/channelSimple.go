package main

import "fmt"

func main() {
	message := make(chan string)

	go func() {
		message <- "World!"
	}()
	go func() {
		message <- "Hello "
	}()

	fmt.Println(<-message, <-message)
}
