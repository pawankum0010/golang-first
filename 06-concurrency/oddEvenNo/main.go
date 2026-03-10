package main

import (
	"fmt"
	"sync"
)

func EvenNo(nums []int, wg *sync.WaitGroup, result *[]int) {
	defer wg.Done()
	for _, num := range nums {
		if num%2 == 0 {
			*result = append(*result, num)
		}
	}
}
func OddNo(nums []int, wg *sync.WaitGroup, result *[]int) {
	defer wg.Done()
	for _, num := range nums {
		if num%2 != 0 {
			*result = append(*result, num)
		}
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var wg sync.WaitGroup
	var evenResult []int
	var oddResult []int

	wg.Add(2)
	go EvenNo(nums, &wg, &evenResult)
	go OddNo(nums, &wg, &oddResult)
	wg.Wait()
	fmt.Println("Even Numbers:", evenResult)
	fmt.Println("Odd Numbers:", oddResult)
}
