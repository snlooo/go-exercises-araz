package main

import "fmt"

func main() {

	//Create an empty slice using the make() function with an initial length of 3 and capacity of 5.
	//Append values to the slice until its length exceeds 5.
	//At each step, print the length and capacity of the slice to show how Go manages memory allocation internally.

	emptySlice := make([]int, 3, 5)

	for len(emptySlice) <= 5 {

		emptySlice = append(emptySlice, 2)
		fmt.Printf("len : %d  cap: %d\n", len(emptySlice), cap(emptySlice))

	}

}
