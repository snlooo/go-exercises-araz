package sync_atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Task 9: Atomic Counter
//
//Use the sync/atomic package to implement a shared counter that is incremented by multiple goroutines without using a Mutex.
//Each goroutine should increment the counter 10 times. After all increments are complete, print the final counter value.

func IncrementWithAtomic(value int32) {

	wg := sync.WaitGroup{}

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&value, 1)
		}()
	}

	wg.Wait()
	fmt.Printf("Final atomic counter value: %d\n", value)
}
