# Session 14 Homework: Testing in Go

## Submission
- Create a new branch named `SES-14/testing-in-go`, and submit all deliverables to this branch in your private GitHub Repository.
- Create a Pull Request (PR) and send the link to the instructor for review.
- At last your repository structure should be like below:

```go
  ├───session-1
  ├───session-2
  ├───session-3
  ├───session-4
  ├───session-5
  ├───session-6
  ├───session-7
  ├───session-8  
  ├───session-9
  ├───session-10
  ├───session-11
  ├───session-12
  ├───session-13
  ├───session-14
         task
         ───unit_testing
              task1.go
              task2.go
         ───table_driven
              task3.go
              task4.go
         ───benchmarking
              task5.go
              task6.go
         ───mocking
              task7.go
              task8.go
         go.mod
         main.go
  └───TaskManagementSystem
```

- Run `go mod init session-14` in your `session-14` root folder
- In the `main.go` call all task's individually
- The deadline for submission is `next session date - 1 day`.

* * * * *

### Topic 1: Writing Unit Tests

**Task 1: Basic Unit Test for a Calculator Function**

1.  Implement a basic calculator function that performs addition, subtraction, multiplication, and division.

2.  Write unit tests for each operation to ensure the calculator behaves correctly. Consider edge cases, such as division by zero, and confirm the function returns an error for invalid inputs.

```go
func Calculate(a, b int, operation string) (int, error) {
    // Add logic for +, -, *, and / here
}
```

**Expected Output:** The tests should cover:

-   Basic arithmetic operations (e.g., `Calculate(2, 3, "+") == 5`).
-   Division by zero returns an error.
-   Unsupported operations return an error.

* * * * *

**Task 2: Unit Test with Assertions on Struct**

1.  Create a function that returns a struct representing a user profile. The function should take inputs like `name` and `age`.

2.  Write unit tests that check the struct's values against expected results. If the age is below 18, the function should set a flag `IsMinor` to `true`.

```go
type UserProfile struct {
    Name    string
    Age     int
    IsMinor bool
}

func NewUserProfile(name string, age int) UserProfile {
    // Add logic to create UserProfile struct
}
```

**Expected Output:**

-   Verify that `Name` and `Age` match input.
-   Confirm that `IsMinor` is set to `true` if `Age` < 18.

* * * * *

### Topic 2: Table-Driven Tests

**Task 3: Table-Driven Test for a String Reversal Function**

1.  Implement a function that reverses a string.

2.  Write a table-driven test that tests multiple cases, including:

    -   An empty string.
    -   A palindrome.
    -   A string with special characters.


```go
    func ReverseString(s string) string {
    // Add logic to reverse string
    }
   ```

**Expected Output:**

-   The tests should verify correct reversal in each scenario.

**Task 4: Table-Driven Test for Temperature Conversion**

1.  Implement a function that converts temperatures between Celsius and Fahrenheit.

2.  Write a table-driven test that checks for various cases, such as:

    -   `0°C` to `32°F`.
    -   `100°C` to `212°F`.
    -   `-40°C` to `-40°F`.
```go
func ConvertTemperature(value float64, scale string) (float64, error) {
    // Add logic for Celsius to Fahrenheit and vice versa
}
```

**Expected Output:**

-   Verify conversion accuracy and ensure unsupported scales return an error.

* * * * *

### Topic 3: Benchmarking

**Task 5: Benchmark for Sorting Algorithm**

1.  Implement a simple sorting algorithm, such as bubble sort, that sorts a slice of integers.

2.  Write a benchmark test to measure the performance of the sorting function with slices of increasing sizes (e.g., 10, 100, and 1,000 elements).

```go
func BubbleSort(arr []int) []int {
    // Add sorting logic here
}
```

**Expected Output:**

-   Measure and analyze the performance for each slice size.

**Task 6: Benchmark for String Concatenation Methods**

1.  Write a function that concatenates a slice of strings using a `+` operator and another that uses `strings.Join()`.

2.  Create benchmarks to compare the performance of both concatenation methods with slices of different sizes (e.g., 10, 100, 1,000).


```go
    func ConcatenateWithPlus(strs []string) string {
    // Add logic to concatenate with +
    }

    func ConcatenateWithJoin(strs []string) string {
    // Add logic to concatenate with strings.Join()
    }
```

**Expected Output:**

-   Compare the performance results and identify the more efficient method.

* * * * *

### Topic 4: Mocking Dependencies

**Task 7: Mocking for a Database Interaction**

1.  Implement a function that interacts with a database interface to retrieve a user by ID.

2.  Mock the database interaction to test the function without requiring an actual database. The test should:

    -   Confirm that the correct user is returned for a valid ID.
    -   Confirm that an error is returned for an invalid ID.

```go
    type Database interface {
    GetUserByID(id int) (User, error)
    }

    func GetUser(db Database, id int) (User, error) {
    // Add logic to retrieve user by ID
    }
```    

**Expected Output:**

-   Ensure the function handles both valid and invalid ID cases correctly.

**Task 8: Mocking an External Service**

1.  Create a function that sends notifications through an interface `Notifier`, which has a method `Send(message string) error`.

2.  Mock the `Notifier` interface to simulate the sending of notifications. Write tests to verify that:

    -   A notification is sent successfully.
    -   The function returns an error if sending fails.

```go   
    type Notifier interface {
    Send(message string) error
    }

    func NotifyUser(n Notifier, message string) error {
    // Add logic to send notification
    }
```    

**Expected Output:**

-   Tests should confirm successful notifications and proper handling of errors when sending fails.

## References
- Task related: [Solving the ‘Palindrome Number Problem’ on LeetCode — Go Solutions Walkthrough](https://medium.com/@AlexanderObregon/solving-the-palindrome-number-on-leetcode-go-solutions-walkthrough-7951276b04e1)
- [Testify Official](https://github.com/stretchr/testify)
- Task related: [Benchmark testing in Go](https://dev.to/stefanalfbo/benchmark-testing-in-go-17dc)
- [Benchmarking golang](https://blog.logrocket.com/benchmarking-golang-improve-function-performance/)