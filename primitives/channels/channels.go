package main

import "fmt"

func main() {
	// Make a channel
	channel := make(chan string)

	// Make a goroutine with an anonymous function
	// Use that channel in the goroutine
	go func() {
		channel <- "Hello from anonymous function"
	}()

	// Make a gouroutine with a normal function
	// Pass and use that channel in the goroutine
	go doSomeWork(channel)

	// Get and print all messages from the channel
	for i := 0; i < 2; i++ {
		msg := <-channel
		fmt.Println(msg)
	}
}

func doSomeWork(channel chan<- string) {
	channel <- "Hello from doSomeWork function"
}
