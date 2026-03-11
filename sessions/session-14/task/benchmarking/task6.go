package benchmarking

import "strings"

//Write a function that concatenates a slice of strings using a + operator and another that uses strings.Join().
//
//Create benchmarks to compare the performance of both concatenation methods with slices of different sizes (e.g., 10, 100, 1,000).

func ConcatenateWithPlus(strs []string) string {

	var completeWord string
	for i := 0; i < len(strs); i++ {
		completeWord += strs[i] + " "
	}
	return completeWord
}

func ConcatenateWithJoin(strs []string) string {
	// Add logic to concatenate with strings.Join()

	completeWordJoin := strings.Join(strs, " ")

	return completeWordJoin
}
