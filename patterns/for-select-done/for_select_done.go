package main

import (
	"fmt"
	"time"
)

func main() {
	// Making a buffered channel
	// Communication between goroutines asynchronous
	done := make(chan bool, 1)

	// Creating a goroutine
	go doWork(done)

	// Closing the unbuffered channel
	// After 3 seconds
	time.Sleep(time.Second * 3)
	close(done)
}

func doWork(done <-chan bool) {
	// Keep on printing "Doing Work"
	// Until the done channel is closed
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("DOING WORK")
		}
	}
}
