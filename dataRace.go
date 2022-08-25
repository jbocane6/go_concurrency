package main

import "fmt"

func main() {
	data := 1
	go func() {
		data++
	}()

	fmt.Print(data)
}
