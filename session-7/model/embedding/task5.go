package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

type Employee struct {
	EmployeeID int
	Position   string
	Person
}

func main() {
	//Embedding Structs
	//Task 5: Embedding a Person Struct
	//
	//Define a Person struct with the following fields:
	//FirstName (string)
	//LastName (string)
	//Define an Employee struct that embeds the Person struct and adds:
	//EmployeeID (int)
	//Position (string)
	//Create an instance of Employee and print both the person's full name and the employee's details.

	_employee := Employee{

		EmployeeID: 12345,
		Position:   "Software Engineer",
		Person:     Person{"Ali", "Aliyev"},
	}

	fmt.Printf("Name: %s %s\n", _employee.FirstName, _employee.LastName)
	fmt.Printf("Employee ID: %d\n", _employee.EmployeeID)
	fmt.Printf("Position:: %s", _employee.Position)

}
