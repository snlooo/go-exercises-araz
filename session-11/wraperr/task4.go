package wraperr

import (
	"fmt"
	"os"
)

//Task 4: Chaining Multiple Wrapped Errors
//
//Modify the previous task so that:
//openFile(filename string) calls another function readFile(f *os.File) error which returns a wrapped error if reading the file fails.
//Both functions should wrap errors with additional context using fmt.Errorf(), and the error should propagate back to main().
//In main(), print the entire chain of wrapped errors using fmt.Errorf("%w").

func OpenFileT4(filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		errRead := ReadFile(file)
		if errRead != nil {
			errRead = fmt.Errorf("%w", errRead)
		}
		wrappedError := fmt.Errorf("Error: failed to open file: %w %w", errRead, err)
		return wrappedError
	}

	return nil
}

func ReadFile(f *os.File) error {

	_, err := f.Read(make([]byte, 1))
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	return nil
}
