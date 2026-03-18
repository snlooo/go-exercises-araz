package buffered_unbuffered_channels

import (
	"fmt"
)

//Task 5: Buffered Channel Operations
//
//Write a program that:
//Creates a buffered channel of size 3.
//Sends 3 values (10, 20, 30) into the channel from the main function without waiting for a Goroutine to receive them.

func ReceiveValues(ch <-chan int) {
	for c := range ch {
		fmt.Println(c)
	}
	fmt.Println("All values received")
}
