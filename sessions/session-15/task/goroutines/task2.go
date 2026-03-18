package goroutines

import (
	"fmt"
	"time"
)

//Task 2: Multiple Goroutines
//
//Create a program that launches two Goroutines:
//The first Goroutine prints the alphabet from A to E with a 200ms delay between each letter.
//The second Goroutine prints numbers from 1 to 5 with a 300ms delay.
//The main function should print Main finished and wait for 2 seconds before ending.

func PrintValues() {

	var alphabet = [5]string{"A", "B", "C", "D", "E"}

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(alphabet[i])
			time.Sleep(200 * time.Millisecond)
		}()

		go func() {
			fmt.Println(i)
			time.Sleep(300 * time.Millisecond)
		}()
	}

	fmt.Println("Main finished")
	time.Sleep(2 * time.Second)
}
