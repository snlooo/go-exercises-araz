package sync_rwmutex

import (
	"fmt"
	"sync"
)

//Task 8: Reader-Writer Example with RWMutex
//
//Implement a shared counter that multiple goroutines can read, but only one can increment.
//Use sync.RWMutex so that multiple readers can access the counter concurrently, while only one writer can modify it.
//Print the counter value at each step to see how it is being accessed and updated.

func IncrementValue(value int) {

	wg := sync.WaitGroup{}
	mu := sync.RWMutex{}

	//Reader
	for range 2 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.RLock()
			fmt.Printf("Reader accessed counter: %d\n", value)
			mu.RUnlock()
		}()
	}

	//Writer
	for range 2 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			value++
			fmt.Printf("Writer accessed counter: %d\n", value)
			mu.Unlock()
		}()

	}

	wg.Wait()
}
