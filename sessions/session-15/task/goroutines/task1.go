package goroutines

import (
	"fmt"
	"time"
)

//Task 1: Basic Goroutine Creation
//
//Write a program that launches a Goroutine to print numbers from 1 to 5 with a 100ms delay between each number.
//The main function should print a message, wait for 1 second, and then end the program.

func PrintNumbers() {

	fmt.Println("Main function started")

	for i := 1; i <= 5; i++ {
		go func() {
			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)
		}()
	}

	fmt.Println("Main function ended")
	time.Sleep(1 * time.Second)
}
