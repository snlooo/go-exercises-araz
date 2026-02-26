package go_generics

//Task 3: Generic Map Function
//
//Write a generic function Map[T any, U any](slice []T, fn func(T) U) []U that:
//
//Applies a given function to each element of a slice and returns a new slice with the results.

func Map[T any, U any](slice []T, fn func(T) U) []U {

	var newSlice []U
	for _, value := range slice {
		newSlice = append(newSlice, fn(value))
	}

	return newSlice
}
