package main

import "fmt"

// Array is a fixed-size sequence of elements of a single type. Its length is part of its
// 	type (e.g., [3]int is different from [4]int), meaning it cannot be resized once declared.
// Arrays are value types; assigning one to a new variable copies the entire data set.

// Slice is a dynamically-sized, flexible view into the elements of an array. It is a
// descriptor consisting of three components: a pointer to the underlying array,
// a length (current number of elements), and a capacity (maximum number of elements it
// can hold before reallocating).
func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice := array[2:3]
	fmt.Println(slice)
}
