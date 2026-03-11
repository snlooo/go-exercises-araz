package benchmarking

//Task 5: Benchmark for Sorting Algorithm
//
//Implement a simple sorting algorithm, such as bubble sort, that sorts a slice of integers.
//
//Write a benchmark test to measure the performance of the sorting function with slices of increasing sizes (e.g., 10, 100, and 1,000 elements).

func BubbleSort(arr []int) []int {
	// Add sorting logic here

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	return arr
}
