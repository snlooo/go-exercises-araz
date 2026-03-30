package sync_waitgroup

import (
	"fmt"
	"sync"
)

//Task 4: Parallel Sum Calculation with WaitGroup
//
//Divide a slice of integers (e.g., [1, 2, 3, 4, 5, 6, 7, 8]) into two parts and calculate the sum of each part in separate goroutines.
//Use a WaitGroup to wait for both sums to complete and then add them together to get the total sum.

func CalculateSum(ch <-chan []int) {

	numbers := <-ch
	mid := len(numbers) / 2
	firstPart := numbers[:mid]
	secondPart := numbers[mid:]
	sumFirst := 0
	sumSecond := 0

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		sumFirst = CalculatePartOfSlices(firstPart)
		fmt.Printf("Partial sum 1: %d\n", sumFirst)
		wg.Done()
	}()

	go func() {
		sumSecond = CalculatePartOfSlices(secondPart)
		fmt.Printf("Partial sum 2: %d\n", sumSecond)
		wg.Done()
	}()

	wg.Wait()
	fmt.Printf("TotalSum: %d\n", sumFirst+sumSecond)

}

func CalculatePartOfSlices(part []int) int {
	sum := 0
	for _, number := range part {
		sum += number
	}

	return sum
}
