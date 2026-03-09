package go_generics

//Task 4: Generic Filter Function
//
//Write a generic function Filter[T any](slice []T, predicate func(T) bool) []T that:
//
//Returns a new slice containing only the elements that satisfy the predicate function.

func Filter[T any](slice []T, predicate func(T) bool) []T {

	var newSlice []T
	for _, v := range slice {
		if predicate(v) {
			newSlice = append(newSlice, v)
		}
	}

	return newSlice
}
