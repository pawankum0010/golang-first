package main

import "fmt"

const age int = 30
const name string = "Pawan Kumar"

func main() {
	const (
		city      string = "Bangalore"
		country   string = "India"
		education string = "B.Tech"
	)

	fmt.Println("This is my age: ", age)
	fmt.Println("This is my name: ", name)
	fmt.Println("This is my city: ", city)
	fmt.Println("This is my country: ", country)
	fmt.Println("This is my education: ", education)
}
