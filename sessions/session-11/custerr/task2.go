package custerr

import "fmt"

//Task 2: Enhanced Error Information
//
//Enhance the DivisionError from Task 1 to include the values of a and b that caused the error.
//Modify the divide(a, b float64) function to use this enhanced DivisionError.
//Print the detailed error message in the main() function when an error occurs.

var DivisionErrorT2 = Exception{message: "Cannot divide %.f by %.f"}

type Exception struct {
	message string
}

func (e Exception) Error() string {
	return e.message
}

func DivideT2(a, b float64) (float64, error) {

	if b == 0 {
		return 0, fmt.Errorf(DivisionErrorT2.message, a, b)
	}

	return a / b, nil
}
