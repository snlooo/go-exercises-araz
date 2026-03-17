package main

import (
	"fmt"
	"session-15/task/buffered_unbuffered_channels"
	"session-15/task/channels_basic_operations"
	"session-15/task/goroutines"
)

func main() {
	//Task 1
	goroutines.PrintNumbers()

	//Task 2
	goroutines.PrintValues()

	//Task 3
	receivedValue := <-channels_basic_operations.SendValue()
	fmt.Printf("Received value: %d \n", receivedValue)

	//Task 4
	ch := make(chan int)
	for i := 1; i <= 5; i++ {
		receivedValueFromChanel := <-channels_basic_operations.SendValueAndClose(ch)
		fmt.Println(receivedValueFromChanel)
	}
	fmt.Println("Channel closed")
	close(ch)

	//Task 5
	ch5 := make(chan int, 3)

	ch5 <- 10
	ch5 <- 20
	ch5 <- 30

	fmt.Println("Sent values into buffered channel")

	close(ch5)

	buffered_unbuffered_channels.ReceiveValues(ch5)

	//Task 6
	ch6 := make(chan string)

	ch6 <- "Hello"
	buffered_unbuffered_channels.NoReverseChanel(ch6)
}
