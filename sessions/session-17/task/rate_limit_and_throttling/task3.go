package rate_limit_and_throttling

import (
	"fmt"
	"time"
)

//Task 3: Throttling with time.Ticker
//
//Create a function that prints a message, such as "Task executed".
//Use a time.Ticker to call this function every 500 milliseconds, simulating throttling.
//After 5 executions, stop the ticker to end the program.

func PrintMessage() {

	tick := time.NewTicker(500 * time.Millisecond)
	defer func() {
		fmt.Println("(ticker stopped)")
		tick.Stop()
	}()
	for range 5 {
		<-tick.C
		fmt.Println("Task executed")
	}

}
