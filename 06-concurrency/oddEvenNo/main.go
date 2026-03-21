package main

import (
	"fmt"
	"sync"
)

func evenNo(nums []int, wg *sync.WaitGroup, result *[]int) {
	defer wg.Done()
	for _, num := range nums {
		if num%2 == 0 {
			*result = append(*result, num)
		}
	}
}

func oddNo(nums []int, wg *sync.WaitGroup, result *[]int) {
	defer wg.Done()
	for _, num := range nums {
		if num%2 != 0 {
			*result = append(*result, num)
		}
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var evenResult []int
	var oddResult []int
	var wg sync.WaitGroup

	wg.Add(2)
	go evenNo(nums, &wg, &evenResult)
	go oddNo(nums, &wg, &oddResult)
	wg.Wait()

	fmt.Println("Even numbers:", evenResult)
	fmt.Println("Odd numbers:", oddResult)
}
