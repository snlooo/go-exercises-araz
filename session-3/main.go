package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	_ = iota
	priorityLow
	priorityMed
	priorityHigh
)

func main() {

	// From session-1
	fmt.Println("Welcome to the Task Management System!")
	fmt.Println("Project start date is: " + time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("Project: Task Management System!")

	//Use arithmetic operators to dynamically calculate how many tasks remain.
	totalCountOfTasks := 100
	completedTasks := rand.Intn(totalCountOfTasks + 1)
	remainTasks := totalCountOfTasks - completedTasks

	fmt.Printf("Tasks remaining %d out of %d\n", remainTasks, totalCountOfTasks)

	//Use logical operators in combination with conditions to check if the project is near completion (e.g., more than 90 tasks completed).
	if completedTasks > 90 {
		fmt.Println("Current project status: IN PROGRESS")
	}

	//Add an if-else block to check if the project is completed based on the number of tasks completed. If completed, update the status.
	if completedTasks == totalCountOfTasks {
		fmt.Println("Project status : COMPLETED")
	} else {
		fmt.Println("Project status : IN PROGRESS")
	}

	//Add a switch statement that categorizes the project progress:
	//If less than 30 tasks are completed, mark the progress as "Starting phase".
	//If 30 to 60 tasks are completed, mark it as "Midway".
	//If more than 60 tasks are completed, mark it as "Final phase".
	switch {
	case completedTasks < 30:

		fmt.Println("Project is in the starting phase")
	case completedTasks > 30 && completedTasks < 60:
		fmt.Println("Project is in the midway")
	case completedTasks > 60:
		fmt.Println("Project is in the final phase")
	}

	fmt.Printf("Task priorities: %d-Low, %d-Medium, %d-High\n", priorityLow, priorityMed, priorityHigh)

	//Use a for loop to iterate tasks by number. For simplicity, create a for loop and iterate the remaining number of tasks, and output with task names (e.g., "Task 1", "Task 2", etc.).
	if completedTasks != totalCountOfTasks {
		fmt.Println("Task list:")
		for i := 0; i < remainTasks; i++ {
			fmt.Printf("Task %d\n", i+1)
		}
	}

	//Implement a simple error-checking mechanism: if the tasks completed exceeds the total number of tasks,
	//print an error message and reset the tasks completed to the maximum number.
	//Simulate an error scenario where the project cannot proceed if no tasks are completed. Use a custom error and handle it gracefully.
	if completedTasks > totalCountOfTasks {
		fmt.Println("exceeds the total number of tasks")
		completedTasks = math.MaxInt
		fmt.Printf("maximum number: %d\n", completedTasks)
	} else if completedTasks < totalCountOfTasks {
		err := errors.New("false")
		fmt.Println("Is the project completed?", err)
	} else {
		fmt.Println("Is the project completed?", true)
	}

}
