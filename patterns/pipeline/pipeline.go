package main

import "fmt"

func main() {
	// Input Stage
	// Creating a slice of integers
	nums := []int{2, 3, 4, 7, 1}

	// Stage 1
	// Convert the slice of numbers into a channel
	dataChannel := sliceToChannel(nums)

	// Stage 2
	// Square those numbers in the dataChannel
	squareChannel := square(dataChannel)

	// Stage 3
	// Print numbers from the squareChannel

	// Since the squareChannel was unbuffered
	// So the for loop will read one by one the elements
	// Which will be pushed in the output channel
	// Of square function
	for n := range squareChannel {
		fmt.Println(n)
	}
}

func sliceToChannel(nums []int) <-chan int {
	// Creating unbuffered channel
	// For synchronous operation
	output := make(chan int)

	// Create a goroutine
	// To convert slice to channel
	go func() {
		for _, n := range nums {
			output <- n
		}

		close(output)
	}()

	// Since its an unbuffered channel
	// So only one element is pushed until its recieved somewhere else
	return output
}

func square(input <-chan int) <-chan int {
	// Creating unbuffered channel
	// For synchronous operation
	output := make(chan int)

	// Create a goroutine
	// To square numbers
	go func() {
		// Since the input channel was unbuffered
		// So the for loop will read one by one the elements
		// Which will be pushed in the output channel
		// Of sliceToChannel function
		for n := range input {
			output <- n * n
		}

		close(output)
	}()

	// Since its an unbuffered channel
	// So only one element is pushed until its recieved somewhere else
	return output
}
