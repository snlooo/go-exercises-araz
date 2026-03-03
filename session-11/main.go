package main

import (
	"errors"
	"fmt"
	"session-11/authsystem"
	"session-11/custerr"
	"session-11/wraperr"
)

func main() {
	//Task 1
	divide, err := custerr.Divide(6, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result:", divide)
	}

	//Task 2
	divideT2, errT2 := custerr.DivideT2(10, 5)
	if errT2 != nil {
		fmt.Println(errT2)
	} else {
		fmt.Println("Result:", divideT2)
	}

	//Task 3
	errT3 := wraperr.OpenFile("nonexistent.txt")
	if errT2 != nil {
		fmt.Println("Wrapped Error:", errT2)

		originalErr := errors.Unwrap(errT2)
		fmt.Println("Original Error:", originalErr)
	}

	//Task 4
	errT4 := wraperr.OpenFileT4("nonexistent.txt")
	if errT4 != nil {
		finalErr := fmt.Errorf("%w", errT3)
		fmt.Println(finalErr)
	}

	//Task 5
	_, errT5 := authsystem.Authenticate("user78", "passworhd1")
	if errT5 != nil {
		fmt.Println(errT5)
	} else {
		fmt.Println("Login successful!")
	}
}
