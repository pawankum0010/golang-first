package main

import "fmt"

// make(map[string]string) - creates a map with string keys and string values
// make(map[string]int) - creates a map with string keys and integer values
// Map is a built-in data structure that stores data as a collection of
// unordered key-value pairs, where every key is unique and maps to exactly one value
func main() {
	studentGrades := make(map[string]int)

	studentGrades["Prince"] = 34
	studentGrades["Vinay"] = 30
	studentGrades["Bob"] = 34
	studentGrades["Cikago"] = 12
	studentGrades["Rahul"] = 90

	fmt.Println("Marks of Bob is :", studentGrades["Bob"])
	studentGrades["Bob"] = 95
	fmt.Println("New Marks of Bob is :", studentGrades["Bob"])
	delete(studentGrades, "Bob")
	fmt.Println("Marks", studentGrades)

	for index, value := range studentGrades {
		// fmt.Printf("Key is %s and marks is %s\n", index, value)
		fmt.Printf("Key is %s and marks is %d\n", index, value)
	}
}
