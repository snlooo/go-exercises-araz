package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Use keywords for quick search in project
// - constKeyword
// - taskCountKeyword
// - currProjectStatusKeyword
// - iotaKeyword
// - formattingKeyword
// - conversionKeyword

// FlagComplete Keyword: constKeyword
const FlagComplete bool = false

// AvailableTaskCount Declare a constant for the total number of available tasks (e.g., 100).
// Keyword: constKeyword
const AvailableTaskCount int = 100

// Use iota to declare constants representing different task priorities (e.g., low, medium, high).
// Keyword : iotaKeyword
const (
	PriorityLow = iota
	PriorityMedium
	PriorityHigh
)

func main() {

	//From Session 1
	fmt.Println("Welcome to the Task Management System!")
	fmt.Println("Project start date is: " + time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(" Project: Task Management System")
	fmt.Println()
	// ----------------------------------------------------------------------------------------------

	pathWithProjects, _ := os.ReadDir(getDir()) //os.ReadDir - get content of directory as an array
	getCountOfDir(pathWithProjects)

	// ----------------------------------------------------------------------------------------------

	priority := getTaskStatus()
	getTaskPriority(priority)

	// ----------------------------------------------------------------------------------------------

	//Keyword: formattingKeyword
	fmt.Printf("Is the project completed? , %t", FlagComplete)
	fmt.Println()
	// ----------------------------------------------------------------------------------------------

	// Convert the boolean value
	// Keyword: conversionKeyword
	fmt.Printf("Converted bool : %q ", strconv.FormatBool(FlagComplete))

}

// Function - for change path , that contains sessions
func getDir() string {

	err := os.Chdir("..")
	if err != nil {
		return err.Error()
	}

	dir, _ := os.Getwd()
	return dir
}

// The number of tasks created so far (integer).
// Keyword: taskCountKeyword
func getCountOfDir(dir []os.DirEntry) {

	count := 0
	for _, entry := range dir {
		// Check if the entry is a directory and contains "session" keyword
		if strings.Contains(entry.Name(), "session") && entry.IsDir() {
			count++
		}
	}
	//Keyword: formattingKeyword
	fmt.Printf("Tasks completed: %d of %d \n", count, AvailableTaskCount)
}

// The current project status (e.g., "active", "in progress").
// A flag to indicate whether the project is completed (boolean).
// Keyword : currProjectStatusKeyword ,  constKeyword
func getTaskStatus() int {

	var statusInProgress = "IN PROGRESS"
	statusComplete := "COMPLETED"
	var taskPriority int

	if FlagComplete {
		fmt.Println("Current project status:" + statusComplete)
		taskPriority = 0
	} else {
		taskPriority = 1
		fmt.Println("Current project status:" + statusInProgress)

	}

	return taskPriority
}

// iota to declare constants representing different task priorities (e.g., low, medium, high).
// Keyword: iotaKeyword
func getTaskPriority(priority int) {

	//Keyword: formattingKeyword
	fmt.Printf("Task priorities: %d - Low, %d - Medium , %d - High \n", PriorityLow, PriorityMedium, PriorityHigh)

	switch priority {
	case PriorityLow:
		fmt.Println("Current task priorities: low")
	case PriorityMedium:
		fmt.Println("Current task priorities: medium")
	case PriorityHigh:
		fmt.Println("Current task priorities: high")

	}
}
