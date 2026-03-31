package sync_mutex

import (
	"fmt"
	"math/rand"
	"sync"
)

//Task 6: Mutex for Safe Map Access
//
//Create a map to store student grades (map[string]int) and launch multiple goroutines that add or update grades.
//Use a sync.Mutex to ensure safe concurrent access to the map.
//After all updates, print the final state of the map.

var mutex sync.Mutex

func UpdateGradeOfStudents(ch <-chan map[string]int) {

	wg := sync.WaitGroup{}
	students := <-ch

	for range 3 {
		wg.Add(1)
		go func() {
			ProcessUpdate(students, &wg)
		}()
	}

	wg.Wait()
	fmt.Printf("Final Grades: %v\n", students)
}

func ProcessUpdate(students map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	for student := range students {
		students[student] = rand.Intn(100)
	}
	mutex.Unlock()
}
