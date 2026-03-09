package main

import (
	"fmt"
	"strconv"
)

func main() {

	totalTasks := 100
	tasksCompleted := 95

	remain, nearCompletion := getCountOfRemainTasks(totalTasks, tasksCompleted)
	status := checkTaskStatus(totalTasks, tasksCompleted)
	listOfTasks := taskListing(totalTasks, tasksCompleted, remain)
	countDownRemainTasks := countDown(remain)

	printDetailsOfTasks("Task1 priority is 1", "Task1 priority is 2", "Task1 priority is 3")
	fmt.Printf("Status check completed.\n%s", listOfTasks)
	fmt.Printf("Recursive countdown of remaining tasks:\n%s", countDownRemainTasks)
	fmt.Printf("Tasks remaining: %d\n", remain)
	fmt.Printf("Is Near Completion : %v\n", nearCompletion)
	fmt.Printf("Task status : %s\n", status)
}

// Move the logic for task completion checks and task remaining calculations into separate functions.
func checkTaskStatus(total, completed int) string {
	status := ""

	if completed == total {
		status = "Project status : COMPLETED"
	} else {
		status = "Project status : IN PROGRESS"
	}

	return status
}

// Move the logic for task completion checks and task remaining calculations into separate functions.
func taskListing(completedTasks, totalCountOfTasks, remainTasks int) string {
	listing := ""
	if completedTasks != totalCountOfTasks {
		for i := 0; i < remainTasks; i++ {
			listing += "Task" + strconv.Itoa(i+1) + "\n"
		}
	}
	return listing
}

// Create a function that returns both the tasksRemaining and a boolean indicating if the project is near completion (e.g., if more than 90 tasks are completed).
// Create a scenario where you use defer to clean up after checking the project status.
// Simulate a panic when the tasksCompleted exceeds the total tasks, and use recover to handle this case gracefully.
func getCountOfRemainTasks(totalTasks, completed int) (remainTasks int, nearCompletion bool) {

	defer func() {
		if a := recover(); a != nil {
			fmt.Println("Recovered from panic: Error: tasksCompleted exceeds total tasks", a)
		}
	}()

	remainTasks = totalTasks - completed

	if completed > totalTasks {
		panic("")
	} else if completed >= 90 {
		nearCompletion = true
	}

	return
}

// Implement a recursive function to simulate
// a countdown that shows how many tasks are left to be completed, reducing the count one by one.
func countDown(remain int) string {

	listing := ""
	for remain > 0 {
		listing += "Task" + strconv.Itoa(remain) + "\n"
		remain--
	}

	return listing
}

// Implement a variadic function to accept and print the details of any number of tasks (e.g., task names).
func printDetailsOfTasks(tasks ...string) {
	fmt.Println(tasks)
}
