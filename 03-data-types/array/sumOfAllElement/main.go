package main

func sumOfAllElement(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func main() {
	num := []int{10, 20, 30, 40, 50}
	sum := sumOfAllElement(num)
	println("Sum of all elements in array is ", sum)
}
