package rate_limit_and_throttling

import (
	"fmt"
	"time"
)

//Task 4: Rate Limiting with Channels
//
//Create a function that simulates a job by printing "Job started" and sleeping for 1 second.
//Limit the execution to a maximum of 2 jobs per second by using a buffered channel with a size of 2.
//After 5 jobs, end the program.

func JobStatus() {
	tick := time.NewTicker(500 * time.Millisecond)
	defer tick.Stop()

	executed := 0
	for executed < 5 {
		<-tick.C
		fmt.Println("Task executed")
		executed++
	}
	fmt.Println("(program ends)")
}
