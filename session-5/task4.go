package main

import "fmt"

func main() {

	//Task 4: Implement a function doubleVariadicElements(numbers ...*int) that takes a variadic parameter (a list of integer pointers) and doubles the value of each element.
	//In the main function, declare a set of integer variables, pass their addresses (pointers) to the function, and print the values before and after modification.
	//Task 4: Doubling Variadic Elements with Pointers
	//Step 1: Define a function doubleVariadicElements(numbers ...*int) that accepts a variadic number of pointers to integers.
	//Step 2: Inside the function, iterate through the pointers and modify the value they point to by doubling it.
	//Step 3: In the main function, declare a few integer variables, pass their addresses (pointers) to the doubleVariadicElements function, and print the values before and after modification.

	x, y, z := 1, 2, 3

	fmt.Printf("Before doubling: %d %d %d\n", x, y, z)
	doubleVariadicElements(&x, &y, &z)
	fmt.Printf("After doubling: %d %d %d\n", x, y, z)

}

func doubleVariadicElements(numbers ...*int) {

	for _, n := range numbers {

		*n *= 2
	}

}
