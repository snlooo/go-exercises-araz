# Homework: Advanced Error Handling in Go

## Objective
In this session - Advanced Error Handling in Go. The tasks gradually build up knowledge, from basic to more complex.

## Submission
- Create a new branch named `SES-11/advanced-error-handling`, and submit all deliverables to this branch in your private GitHub Repository.
- Create a Pull Request (PR) and send the link to the instructor for review.
- At last your repository structure should be like below:

```go
  ├───session-11
         task
         ───custerr
              task1.go
              task2.go
         ───wraperr
              task3.go
              task4.go
         ───authsystem
              task5.go
         go.mod
         main.go
  └───TaskManagementSystem
```

- Run `go mod init session-11` in your `session-11` root folder
- In the `main.go` call all task's individually
- The deadline for submission is `next session date - 1 day`.



### **Homework 1: Custom Error Types**

**Task 1: Create a Custom Error Type**

1.  Create a custom error type `DivisionError` that is returned when someone tries to divide by zero.

2.  Write a function `divide(a, b float64) (float64, error)` that:

    -   Returns the result of `a / b` if `b` is not zero.
    -   Returns a `DivisionError` if `b` is zero.
3.  In the `main()` function, handle the error and print an appropriate message.

**Expected Output:**

```go
Error: Division by zero is not allowed.
```

If `b` is non-zero, the output should be:


```go
Result: 5.0
```

* * * * *

**Task 2: Enhanced Error Information**

1.  Enhance the `DivisionError` from Task 1 to include the values of `a` and `b` that caused the error.
2.  Modify the `divide(a, b float64)` function to use this enhanced `DivisionError`.
3.  Print the detailed error message in the `main()` function when an error occurs.

**Expected Output (for divide(10, 0)):**

```go
Error: Cannot divide 10 by 0.
```

For a valid division, the output should be:

```go
Result: 5.0
```

* * * * *

### **Homework 2: Wrap Errors**

**Task 3: Wrapping and Unwrapping Errors**

1.  Write a function `openFile(filename string) error` that tries to open a file using `os.Open()`.
2.  If the file does not exist, return a wrapped error using `fmt.Errorf()` that includes the message `"file not found"`, wrapping the original error.
3.  In `main()`, call `openFile("nonexistent.txt")` and handle the error by:
    -   Printing the wrapped error.
    -   Using `errors.Unwrap()` to print the original error.

Import Go builtin `os` package for `os.Open()`

**Expected Output:**

```go
Wrapped Error: file not found: open nonexistent.txt: no such file or directory
Original Error: open nonexistent.txt: no such file or directory
```

* * * * *

**Task 4: Chaining Multiple Wrapped Errors**

1.  Modify the previous task so that:
    -   `openFile(filename string)` calls another function `readFile(f *os.File) error` which returns a wrapped error if reading the file fails.
2.  Both functions should wrap errors with additional context using `fmt.Errorf()`, and the error should propagate back to `main()`.
3.  In `main()`, print the entire chain of wrapped errors using `fmt.Errorf("%w")`.

**Expected Output (when file does not exist):**

```go
Error: failed to open file: failed to read file: open nonexistent.txt: no such file or directory
```
* * * * *

### **Task 5: Refactor an Existing Project to Use Custom Errors**

**Task: Refactor a User Authentication System to Handle Errors Properly**

1.  You are given a simplified user authentication system that does not currently use proper error handling. Refactor the code to:
    -   Use custom error types where appropriate.
    -   Wrap and unwrap errors to provide detailed information in case of failures.

Here is the original code:

```go
package main

import (
	"fmt"
)

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func authenticate(username, password string) bool {
	if passValue, ok := users[username]; ok {
		if passValue == password {
			return true
		}
	}
	return false
}

func main() {
	if authenticate("user1", "password1") {
		fmt.Println("Login successful!")
	} else {
		fmt.Println("Invalid username or password")
	}
}
```

1.  Refactor this code to:
    -   Create custom error types for different failure cases (`UserNotFoundError` and `InvalidPasswordError`).
    -   Modify the `authenticate()` function to return detailed errors instead of a `bool`.
    -   Wrap errors to add context at each level, and then handle these errors in the `main()` function by printing detailed messages.

**Refactored Code Expectations:**

-   If the user does not exist, return a `UserNotFoundError`.
-   If the password is incorrect, return an `InvalidPasswordError`.
-   In `main()`, display the proper error messages based on the returned errors.

**Expected Output (for wrong username):**

```go
Error: user not found: "user3"
```

**Expected Output (for wrong password):**

```go
Error: invalid password for user: "user1"
```

**Expected Output (for correct login):**

```go
Login successful!
```

## References
- [Golang Errors](https://www.programiz.com/golang/errors)
- [Golang Panic, Refer, Defer](https://www.programiz.com/golang/defer-panic-recover)
- [Go wrapping multiple error](https://medium.com/@vajahatkareem/golang-error-wrapping-multierror-759d04bdbfaf)
- [Handle errors in Go with errors.Is() and errors.As()](https://gosamples.dev/check-error-type/)