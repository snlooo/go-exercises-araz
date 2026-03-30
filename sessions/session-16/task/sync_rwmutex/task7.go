package sync_rwmutex

import (
	"fmt"
	"math/rand"
	"sync"
)

//Task 7: RWMutex for Reading and Writing to a Map
//
//Create a program with a map that represents a database of users and their ages.
//Use sync.RWMutex to allow multiple goroutines to read from the map concurrently, but only one to write at a time.
//Run multiple readers and a single writer, and print the final state of the map.

func UpdateAgeOfUsers(ch <-chan map[string]int) {

	users := <-ch

	mu := sync.RWMutex{}
	wg := sync.WaitGroup{}

	//READ
	for range 3 {
		wg.Add(1)
		go func() {
			mu.RLock()
			fmt.Printf("Students: %v\n", users)
			mu.RUnlock()
			wg.Done()
		}()
	}

	//Write
	wg.Go(func() {
		mu.Lock()
		for user := range users {
			users[user] = rand.Intn(30)
		}
		mu.Unlock()
	})

	wg.Wait()
	fmt.Printf("User data:: %v\n", users)
}
