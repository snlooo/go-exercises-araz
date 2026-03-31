package select_statement

import (
	"fmt"
	"time"
)

//Task 2: Timeout with select
//
//Create a function that waits on a channel, but only for a certain period.
//Use a select statement with a timeout to return a "timeout" message if the data doesn't arrive within 3 seconds.
//Test this function by creating a channel but not sending anything to it, so it times out.

func WaitWithTimeOut(ch <-chan string) {
	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(3 * time.Second):

		fmt.Println("Timeout occurred: no message received")
	}
}
