package select_statement

import (
	"fmt"
	"time"
)

//Task 1: Basic Channel Operations with select
//
//Create two channels, ch1 and ch2, and launch two goroutines that:
//Send the string "Hello from ch1" to ch1 after a 1-second delay.
//Send the string "Hello from ch2" to ch2 after a 2-second delay.
//Use a select statement to receive from either ch1 or ch2 and print whichever message arrives first.

func PrintValuesFromChannels() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Hello from ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Hello from ch2"
	}()

	select {
	case receivedValue := <-ch1:
		fmt.Printf("Received: %v \n", receivedValue)
	case receivedValue := <-ch2:
		fmt.Printf("Received: %v \n", receivedValue)
	}

}
