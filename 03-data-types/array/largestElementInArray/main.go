package main

import (
	"fmt"
	"math"
)

func findLargestElement(arr []int) int {
	max := math.MinInt64
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	return max
}

func main() {
	num := []int{0, -10, 5, 3, 99, 27, -45}
	if len(num) == 0 {
		println("Array is empty")
		return
	}
	largest := findLargestElement(num[:])
	fmt.Println("Largest element in the array is:", largest)
}
