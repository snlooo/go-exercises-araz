package main

import "fmt"

func main() {

	//Task 1: Create a Go program that demonstrates how to declare, assign, and dereference a pointer.
	//Declare an integer variable x and assign it a value of 10.
	//Declare a pointer variable that holds the memory address of x.
	//Print the value of x, the address of x, and the value stored at the pointer.

	x := 10
	var xAddress *int
	xAddress = &x

	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Address of x: %#v\n", &x)
	fmt.Printf("Value at pointer: %d\n", *xAddress)

}
