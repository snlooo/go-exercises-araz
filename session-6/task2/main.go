package main

import "fmt"

func main() {

	//Create a slice of integers from an array of 10 elements: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9].
	//Perform the following operations on the slice:
	//Extract the slice from index 2 to 6 and print it.
	//Append values 10, 11, 12 to the slice and print the result.
	//Show the effect of append() on the underlying array by printing both the slice and the original array after appending.

	sliceOfInt := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("Sub-slice:", sliceOfInt[2:6])

	appendedSlice := append(sliceOfInt, 10, 11, 12)
	fmt.Println("Appended slice:", appendedSlice)
	//
	fmt.Printf("Before append len : %d  cap: %d\nOriginal slice: %v\n", cap(sliceOfInt), len(sliceOfInt), sliceOfInt)
	sliceOfInt = append(sliceOfInt, 10)
	fmt.Printf("After append len : %d  cap: %d\nModified slice: %d", cap(sliceOfInt), len(sliceOfInt), sliceOfInt)

}
