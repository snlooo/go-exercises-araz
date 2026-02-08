package main

import "fmt"

func main() {

	//Pointers vs. Values
	// Task 2: Write a Go program that illustrates the difference between passing a value and passing a pointer to a function.
	//  Create two functions:
	//    incrementByValue(val int) -- this should take an integer argument and increment it by 1.
	//    incrementByPointer(ptr *int) -- this should take a pointer to an integer and increment the value it points to by 1.
	//  In the main function:
	//     Declare an integer x and assign it a value of 5.
	//     Call both functions with x and observe the difference in behavior.

	x := 5
	fmt.Printf("Value of x before incrementing by value: %d\n", x)
	incrementByValue(x)
	fmt.Printf("Value of x after incrementing by value: %d\n", x)

	fmt.Printf("Value of x before by pointer: %d\n", x)
	incrementByPointer(&x)
	fmt.Printf("Value of x after incrementing by pointer: %d\n", x)

}

func incrementByValue(val int) {
	val++
}

func incrementByPointer(ptr *int) {
	*ptr++
}
