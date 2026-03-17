package channels_basic_operations

import "time"

//Task 3: Sending and Receiving from an Unbuffered Channel
//
//Write a program that:
//Creates an unbuffered channel of type int.
//Launches a Goroutine that sends a value (e.g., 42) into the channel after 500ms.
//The main function should receive the value from the channel and print it.

func SendValue() chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch <- 42
	}()
	return ch
}
