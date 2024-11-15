package main

import (
	"fmt"
	"time"
)

func main() {
	// Making goroutines to run doSomeWork() function on different threads
	go doSomeWork(1)
	go doSomeWork(2)
	go doSomeWork(3)

	// Waiting for goroutines to print something
	// If below statement is removed only "Hello, World!" gets printed
	time.Sleep(2 * time.Second)

	fmt.Println("Hello, World!")
}

func doSomeWork(num int) {
	fmt.Println(num)
}
