package custerr

import (
	"errors"
	"fmt"
)

//Task 1: Create a Custom Error Type
//
//Create a custom error type DivisionError that is returned when someone tries to divide by zero.
//
//Write a function divide(a, b float64) (float64, error) that:
//
//Returns the result of a / b if b is not zero.
//Returns a DivisionError if b is zero.
//In the main() function, handle the error and print an appropriate message.

var DivisionError = errors.New("Error: Division by zero is not allowed")

func Divide(a, b float64) (float64, error) {

	if b == 0 {
		return 0, fmt.Errorf(DivisionError.Error())
	}

	return a / b, nil
}
