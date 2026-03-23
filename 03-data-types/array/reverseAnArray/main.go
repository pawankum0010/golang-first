package main

import "fmt"

func reverseAnArray(arr []int) []int {
	start := 0
	end := len(arr) - 1
	for start < end {
		arr[start], arr[end] = arr[end], arr[start] //Swap element using concise syntax
		start++
		end--
	}
	return arr
}

func main() {
	num := []int{10, 20, 30, 40, 50}
	reversed := reverseAnArray(num)
	fmt.Println(reversed)
}
