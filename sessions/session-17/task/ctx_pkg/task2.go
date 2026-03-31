package ctx_pkg

import (
	"context"
	"fmt"
	"time"
)

//Task 2: Context with Timeout
//
//Create a function that simulates a task by sleeping for 5 seconds.
//Use context.WithTimeout to set a timeout of 2 seconds, so the task should be cancelled before it completes.
//In your program, print whether the task was completed or cancelled based on the timeout.

func SimulateCancellation() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	select {
	case <-ctx.Done():
		fmt.Println("Task cancelled due to timeout")
	case <-time.After(3 * time.Second):
		fmt.Println("Task completed successfully")
	}

}
