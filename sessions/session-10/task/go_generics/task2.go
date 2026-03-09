package go_generics

//Write a generic function MinMax[T int | float64](slice []T) (T, T) that:
//
//Finds the minimum and maximum values in a slice.

func MinMax[T int | float64](slice []T) (T, T) {

	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[i] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}

	return slice[0], slice[len(slice)-1]
}
