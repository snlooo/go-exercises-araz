package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	// From session-1
	fmt.Println("Welcome to the Task Management System!")
	fmt.Println("Project start date is: " + time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("Project: Task Management System!")

	//The current project status (e.g., "active", "in progress").
	statusInProgress := "IN PROGRESS"
	//statusIsActive := "ACTIVE"
	// The number of tasks created so far (integer).
	taskCount := 25
	// A flag to indicate whether the project is completed (boolean).
	var flagComplete bool

	// Declare a constant for the total number of available tasks (e.g., 100).
	const availableTasks int = 100

	//Use iota to declare constants representing different task priorities (e.g., low, medium, high).
	const (
		priorityLow = iota
		priorityMedium
		priorityHigh
	)

	// Project status.
	// Current project status: IN PROGRESS
	fmt.Printf("Current task status: %v\n", statusInProgress)

	//Tasks completed: 25 out of 100
	fmt.Printf("Tasks completed: %d out of %d\n", taskCount, availableTasks)

	//Task priorities: 0-Low, 1-Medium, 2-High
	fmt.Printf("Task priorities: %d-Low, %d-Medium, %d-High\n", priorityLow, priorityMedium, priorityHigh)

	//  Is the project completed? false
	//Type conversion and string formatting:
	//Convert the boolean value (project completed) into a string and display it in the output message.
	fmt.Printf("Is the project completed? " + strconv.FormatBool(flagComplete))

}
