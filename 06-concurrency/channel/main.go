package main

// ============================================================================
// CHANNELS IN GO - Communication Between Goroutines
// ============================================================================
//
// A Go channel is a built-in data structure that allows two goroutines
// to communicate and synchronize their execution.
//
// KEY CHARACTERISTICS:
// 1. It acts as a conduit for a specific data type (type-safe).
// 2. Data is passed between goroutines using send (<-) and receive operations.
// 3. Channels can be:
//    - UNBUFFERED: Blocking operation. Both sender and receiver must be ready.
//                  Sender blocks until receiver is ready; receiver blocks until data arrives.
//    - BUFFERED: Non-blocking (until buffer is full). Stores limited values.
//                Channel created with capacity: make(chan int, 5)
//
// SYNTAX:
// - Send: channel <- value          (push value into channel)
// - Receive: value := <-channel     (pull value from channel)
// - Range: for value := range channel (iterate until channel is closed)
//
// ============================================================================
// make(chan int, 0) - creates an unbuffered channel of type int
// make(chan string, 3) - creates a buffered channel of type string with capacity 3

import (
	"fmt"
	"math/rand"
	"time"
)

// ============================================================================
// EXAMPLE 1: UNBUFFERED CHANNEL WITH GOROUTINE
// ============================================================================
// This function receives numbers from a channel and processes them.
// It blocks at "for num := range numchan" until data arrives on the channel.
// The range loop continues until the channel is closed.
//
// HOW IT WORKS:
//   - Goroutine waits for data on the channel
//   - Main function sends random numbers to the channel
//   - Goroutine receives and processes each number
//   - Synchronization happens automatically: sender waits for receiver,
//     and receiver waits for sender
func processNum(numchan chan int) {
	for num := range numchan {
		fmt.Println("Processing number...", num)
		time.Sleep(300000 * time.Microsecond) // Simulate processing time
	}
	fmt.Println("Channel closed. Processing complete.")
}

// ============================================================================
// EXAMPLE 2: BUFFERED CHANNEL
// ============================================================================
// Demonstrates buffered channels that can hold multiple values
func demonstrateBufferedChannel() {
	// Create a buffered channel with capacity of 3
	messageChannel := make(chan string, 3)

	// These sends don't block because buffer has space
	messageChannel <- "First message"
	messageChannel <- "Second message"
	messageChannel <- "Third message"

	fmt.Println("Buffered Channel Examples:")
	fmt.Println(<-messageChannel) // Output: First message
	fmt.Println(<-messageChannel) // Output: Second message
	fmt.Println(<-messageChannel) // Output: Third message
}

// ============================================================================
// EXAMPLE 3: MULTIPLE GOROUTINES
// ============================================================================
func processWithMultipleWorkers() {
	// Buffered channel to distribute work
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Start 2 worker goroutines
	for w := 1; w <= 2; w++ {
		go func(workerID int) {
			for job := range jobs {
				fmt.Printf("Worker %d processing job %d\n", workerID, job)
				results <- job * 2 // Send result
				time.Sleep(100 * time.Millisecond)
			}
		}(w)
	}

	// Send jobs
	for j := 1; j <= 4; j++ {
		jobs <- j
	}
	close(jobs) // Signal workers that no more jobs

	// Collect results
	for i := 0; i < 4; i++ {
		fmt.Printf("Result: %d\n", <-results)
	}
}

// ============================================================================
// MAIN FUNCTION
// ============================================================================
func main() {
	fmt.Println("=== CHANNEL EXAMPLE 1: Unbuffered Channel ===")
	// Create an unbuffered channel of type int
	numChan := make(chan int)

	// Launch goroutine that waits to receive from channel
	go processNum(numChan)

	// Send numbers to the channel (will block until goroutine receives)
	for i := 0; i < 5; i++ {
		numChan <- rand.Intn(10)
	}

	// Close the channel to signal the goroutine to stop
	close(numChan)

	// Small delay to let goroutine finish
	time.Sleep(2 * time.Second)

	fmt.Println("\n=== CHANNEL EXAMPLE 2: Buffered Channel ===")
	demonstrateBufferedChannel()

	fmt.Println("\n=== CHANNEL EXAMPLE 3: Multiple Goroutines ===")
	processWithMultipleWorkers()

	// ========================================================================
	// WHY THE ORIGINAL CODE CAUSES ISSUES:
	// ========================================================================
	// Original deadlock example (commented out):
	//
	// messageChannel := make(chan string)
	// messageChannel <- "Hello, Channels!"  // DEADLOCK HERE!
	// msg := <-messageChannel
	// fmt.Println(msg)
	//
	// WHY IT DEADLOCKS:
	// 1. Channel is UNBUFFERED (no capacity)
	// 2. Main goroutine tries to SEND ("Hello, Channels!")
	// 3. Sender blocks because NO goroutine is ready to RECEIVE
	// 4. Only one goroutine exists (main), so nothing can receive
	// 5. Result: DEADLOCK - main goroutine waits forever for receiver
	//
	// SOLUTION:
	// - Use a goroutine to receive: go func() { <-messageChannel }()
	// - Use a buffered channel: make(chan string, 1)
	// - Use separate goroutines for send and receive
	// ========================================================================
}
