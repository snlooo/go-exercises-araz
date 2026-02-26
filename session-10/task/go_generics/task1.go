package go_generics

//Task 1: Generic Function for Sum
//
//Write a generic function Sum[T int | float64](slice []T) T that:
//
//Calculates and returns the sum of elements in a slice.

func Sum[T int | float64](slice []T) T {

	sum := T(0)
	for _, x := range slice {

		sum += x
	}

	return sum
}
