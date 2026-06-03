package main

import "fmt"

func countDigitOccurrance(input string) map[rune]int {
	count := make(map[rune]int)
	for _, char := range input {
		count[char]++
	}
	return count
}

func main() {
	input := "1234561345753"
	digitCount := countDigitOccurrance(input)
	if len(digitCount) == 0 {
		fmt.Println("No no found in string")
	}
	for digit, count := range digitCount {
		if count > 1 {
			fmt.Printf("Digit %q occurs %d times\n", digit, count)
		}
	}
}
