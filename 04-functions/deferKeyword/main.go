package main

import "fmt"

// In Go, the defer keyword schedules a function call to be executed immediately
// After the surrounding function returns, following a Last-In, First-Out (LIFO) order.
func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Starting of program")
	data := add(6, 8)
	defer fmt.Println("Data is", data)
	defer fmt.Println("Middle of program")
	fmt.Println("End of program")
}
