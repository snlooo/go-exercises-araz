package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {

	//v1
	fmt.Printf("Welcome to the Task Management System!\nProject start date is: 2024-09-18 00:00:00\nProject: Task Management System\n")

	fmt.Println("---------------------------------------------------------------")

	//v2
	fmt.Println("Welcome to the Task Management System!")
	fmt.Println("Project start date is: " + time.Now().Format("2006-01-02 15:04:05"))
	//Get Project Name
	dir, _ := os.Getwd()
	fmt.Println("Project: " + filepath.Base(dir))
}
