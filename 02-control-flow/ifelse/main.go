package main

func main() {
	age := 19

	if age <= 18 {
		println("You are a minor")
	} else {
		println("You are an adult")
	}

	// if num := 11; num%2 == 0 {
	// 	println("Even number")
	// } else {
	// 	println("Odd number")
	// }
	if num := 15; num%2 == 0 {
		println("Even number")
	} else {
		println("Odd number")
	}

	var role string = "admin"
	var hasPermission bool = false

	if role == "admin" && hasPermission {
		println("You can access admin panel")
	} else if role == "user" && hasPermission {
		println("You can access user panel")
	} else {
		println("Access denied")
	}

}
