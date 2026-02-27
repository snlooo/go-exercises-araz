package wraperr

import (
	"fmt"
	"os"
)

//Task 3: Wrapping and Unwrapping Errors
//
//Write a function openFile(filename string) error that tries to open a file using os.Open().
//If the file does not exist, return a wrapped error using fmt.Errorf() that includes the message "file not found", wrapping the original error.
//In main(), call openFile("nonexistent.txt") and handle the error by:
//Printing the wrapped error.
//Using errors.Unwrap() to print the original error.
//Import Go builtin os package for os.Open()

func OpenFile(filename string) error {
	_, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("file not found: %w", err)
	}
	return nil
}
