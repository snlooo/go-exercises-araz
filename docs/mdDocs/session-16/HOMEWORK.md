# Session 16 Homework: Select and Sync

## Submission
- Create a new branch named `SES-16/select-sync`, and submit all deliverables to this branch in your private GitHub Repository.
- Create a Pull Request (PR) and send the link to the instructor for review.
- At last your repository structure should be like below:

```go
  ├───session-16
         task
         ───select_statement
              task1.go
              task2.go
         ───sync_waitgroup
              task3.go
              task4.go
         ───sync_mutex
              task5.go
              task6.go
         ───sync_rwmutex
              task7.go
              task8.go
         ───sync_atomic
              task9.go
         go.mod
         main.go
  └───TaskManagementSystem
```

- Run `go mod init session-16` in your `session-16` root folder
- In the `main.go` call all task's individually
- The deadline for submission is `next session date - 1 day`.

* * * * *

### Topic 1: `select` Statement

**Task 1: Basic Channel Operations with `select`**

1.  Create two channels, `ch1` and `ch2`, and launch two goroutines that:
    -   Send the string "Hello from ch1" to `ch1` after a 1-second delay.
    -   Send the string "Hello from ch2" to `ch2` after a 2-second delay.
2.  Use a `select` statement to receive from either `ch1` or `ch2` and print whichever message arrives first.

**Expected Output:**

-   The message from `ch1` will print first due to its shorter delay:

```go
Received: Hello from ch1
```

**Task 2: Timeout with `select`**

1.  Create a function that waits on a channel, but only for a certain period. Use a `select` statement with a timeout to return a "timeout" message if the data doesn't arrive within 3 seconds.
2.  Test this function by creating a channel but not sending anything to it, so it times out.

**Expected Output:**



```go
Timeout occurred: no message received
```

* * * * *

### Topic 2: `sync.WaitGroup`

**Task 3: Basic WaitGroup Usage**

1.  Create a program that launches three goroutines, each of which prints a message indicating that it's starting, waits 1 second, and then prints a message indicating it has finished.
2.  Use a `sync.WaitGroup` to wait for all goroutines to complete before the main function exits.

**Expected Output:**

```go
Goroutine 1 starting
Goroutine 2 starting
Goroutine 3 starting
Goroutine 1 finished
Goroutine 2 finished
Goroutine 3 finished
All goroutines have completed
```

**Task 4: Parallel Sum Calculation with WaitGroup**

1.  Divide a slice of integers (e.g., `[1, 2, 3, 4, 5, 6, 7, 8]`) into two parts and calculate the sum of each part in separate goroutines.
2.  Use a `WaitGroup` to wait for both sums to complete and then add them together to get the total sum.

**Expected Output:**

```go
Partial sum 1: 10
Partial sum 2: 26
Total sum: 36
```

* * * * *

### Topic 3: `sync.Mutex`

**Task 5: Protect Shared Counter with Mutex**

1.  Create a shared integer counter that is incremented by multiple goroutines.
2.  Use a `sync.Mutex` to ensure that only one goroutine can increment the counter at a time.
3.  After all increments are done, print the final value of the counter.

**Expected Output:**

```go
Final counter value: 100
```

**Task 6: Mutex for Safe Map Access**

1.  Create a map to store student grades (`map[string]int`) and launch multiple goroutines that add or update grades.
2.  Use a `sync.Mutex` to ensure safe concurrent access to the map.
3.  After all updates, print the final state of the map.

**Expected Output:**

```go
Final Grades: map[Garay:90 Ali:85 Medina:78]
```

* * * * *

### Topic 4: `sync.RWMutex`

**Task 7: RWMutex for Reading and Writing to a Map**

1.  Create a program with a map that represents a database of users and their ages.
2.  Use `sync.RWMutex` to allow multiple goroutines to read from the map concurrently, but only one to write at a time.
3.  Run multiple readers and a single writer, and print the final state of the map.

**Expected Output:**

```go
User data: map[Garay:20 Ali:25 Medina:28]
```

**Task 8: Reader-Writer Example with RWMutex**

1.  Implement a shared counter that multiple goroutines can read, but only one can increment.
2.  Use `sync.RWMutex` so that multiple readers can access the counter concurrently, while only one writer can modify it.
3.  Print the counter value at each step to see how it is being accessed and updated.

**Expected Output:**

```go
Reader accessed counter: 0
Writer updated counter: 1
Reader accessed counter: 1
Writer updated counter: 2
```

* * * * *

### Topic 5: `sync/atomic` Package

**Task 9: Atomic Counter**

1.  Use the `sync/atomic` package to implement a shared counter that is incremented by multiple goroutines without using a `Mutex`.
2.  Each goroutine should increment the counter 10 times. After all increments are complete, print the final counter value.

**Expected Output:**

```go
Final atomic counter value: 100
```


## References
- [Atomic package](https://pkg.go.dev/sync/atomic)
- [Race detector in Golang](https://go.dev/doc/articles/race_detector)