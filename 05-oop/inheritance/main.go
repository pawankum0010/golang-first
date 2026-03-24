package main

type Person struct {
	Name   string
	Age    int
	Gender string
}
type Address struct {
	Person
	City    string
	Country string
}
type Employee struct {
	Address
	Company     string
	Designation string
	Department  string
}
type Salary struct {
	Employee
	Basic  float64
	HRA    float64
	DA     float64
	Tax    float64
	NetPay float64
}

func main() {
	salary := Salary{
		Employee: Employee{
			Address: Address{
				Person: Person{
					Name:   "Pawan Kumar",
					Age:    30,
					Gender: "Male",
				},
				City:    "New Delhi",
				Country: "India",
			},
			Company:     "Tech Company",
			Designation: "Software Engineer",
			Department:  "Development",
		},
		Basic:  50000,
		HRA:    20000,
		DA:     10000,
		Tax:    15000,
		NetPay: 65000,
	}

	println("Name:", salary.Name)
	println("Age:", salary.Age)
	println("Gender:", salary.Gender)
	println("City:", salary.City)
	println("Country:", salary.Country)
	println("Company:", salary.Company)
	println("Designation:", salary.Designation)
	println("Department:", salary.Department)
	println("Basic Salary:", salary.Basic)
	println("HRA:", salary.HRA)
	println("DA:", salary.DA)
	println("Tax:", salary.Tax)
	println("Net Pay:", salary.NetPay)

}
