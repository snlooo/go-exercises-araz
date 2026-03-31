package ctx_pkg

import (
	"context"
	"fmt"
	"time"
)

//Topic 1: Context Package for Cancellation
//Task 1: Cancellation with Context
//
//Write a program that starts a long-running goroutine (e.g., a loop that prints numbers from 1 to 10, with a 1-second delay between each print).
//Use the context.WithCancel function to create a cancellable context.
//After 3 seconds, cancel the context and stop the goroutine using the cancellation signal.

func PrintNumbers() {

	_, cancel := context.WithCancel(context.Background())

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Second * 3)
	cancel()
	fmt.Println("(cancellation)")

}
