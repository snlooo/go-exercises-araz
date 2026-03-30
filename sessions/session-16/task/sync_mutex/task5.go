package sync_mutex

import (
	"fmt"
	"sync"
)

// Task 5: Protect Shared Counter with Mutex
//
// Create a shared integer counter that is incremented by multiple goroutines.
// Use a sync.Mutex to ensure that only one goroutine can increment the counter at a time.
// After all increments are done, print the final value of the counter.
var mu = sync.Mutex{}

func IncrementValue() {

	mainValue := 0
	wg := sync.WaitGroup{}

	for range 100 {
		wg.Add(1)
		go processIncrement(&mainValue, &wg)
	}
	wg.Wait()
	fmt.Printf("Final counter value: %d", mainValue)
}

func processIncrement(value *int, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	*value++
	mu.Unlock()
}
