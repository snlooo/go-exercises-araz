package main

import "fmt"

func main() {

	//3. Using Pointers with Functions
	//Task 3: Create a function swap that swaps the values of two integer variables using pointers.
	//Define a function swap(a *int, b *int) that swaps the values stored at the addresses passed in.
	//In the main function:
	//Declare two integer variables x and y with values 10 and 20.
	//Call the swap function and print the values of x and y before and after swapping.

	x := 10
	y := 20
	fmt.Printf("Before swapping: x = %d, y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("After swapping: x = %d, y = %d", x, y)

}

func swap(a, b *int) {

	c := new(int)
	*c = *a
	*a = *b
	*b = *c

}
