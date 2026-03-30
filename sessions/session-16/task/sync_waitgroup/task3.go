package sync_waitgroup

import (
	"fmt"
	"sync"
	"time"
)

//Task 3: Basic WaitGroup Usage
//
//Create a program that launches three goroutines, each of which prints a message indicating that it's starting,
//waits 1 second, and then prints a message indicating it has finished.
//Use a sync.WaitGroup to wait for all goroutines to complete before the main function exits.

func InformStarting() {

	wg := sync.WaitGroup{}

	for i := range 3 {
		wg.Add(1)
		go func() {
			fmt.Printf("Goroutine %d starting\n", i+1)
			wg.Done()
		}()
	}

	wg.Wait()
	time.Sleep(time.Second)
	fmt.Printf("All goroutines have completed\n")

}
