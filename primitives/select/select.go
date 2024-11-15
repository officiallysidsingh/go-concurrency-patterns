package main

import (
	"fmt"
)

func main() {
	// Make two channels
	channel1 := make(chan string)
	channel2 := make(chan string)

	// Make to goroutines using anonymous function
	// Send data to both channels in seperate goroutines
	go func() {
		channel1 <- "Hello from channel 1"
	}()

	go func() {
		channel2 <- "Hello from channel 2"
	}()

	// Use select statement to get which channel gets the data first
	// This also determines which goroutine finishes execution first
	select {
	case msg := <-channel1:
		fmt.Println(msg)
	case msg := <-channel2:
		fmt.Println(msg)
	}
}
