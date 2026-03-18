package unit_testing

import (
	"errors"
	"fmt"
)

//Task 1: Basic Unit Test for a Calculator Function
//
//Implement a basic calculator function that performs addition, subtraction, multiplication, and division.
//
//Write unit tests for each operation to ensure the calculator behaves correctly.
//Consider edge cases, such as division by zero, and confirm the function returns an error for invalid inputs.

func Calculate(a, b int, operation string) (int, error) {
	// Add logic for +, -, *, and / here

	switch operation {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0.0, errors.New("Cannon devide by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unknown operation: %s", operation)
	}
}
