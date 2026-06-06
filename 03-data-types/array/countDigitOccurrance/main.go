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
	input := "123544562134541534152358246739569"
	digitCount := countDigitOccurrance(input)
	if len(digitCount) == 0 {
		fmt.Println("No no found in this string")
	}
	for digit, count := range digitCount {
		if count > 4 {
			fmt.Printf("Digit %q has occurance %d\n", digit, count)
		}
	}
}
