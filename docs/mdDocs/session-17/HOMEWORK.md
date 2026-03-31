# Session 17 Homework: Advanced Concurrency Patterns

## Submission
- Create a new branch named `SES-17/advanced-concurrency`, and submit all deliverables to this branch in your private GitHub Repository.
- Create a Pull Request (PR) and send the link to the instructor for review.
- At last your repository structure should be like below:

```go
  ├───session-17
         task
         ───ctx_pkg
              task1.go
              task2.go
         ───rate_limit_and_throttling
              task3.go
              task4.go
         go.mod
         main.go
  └───TaskManagementSystem
```

- Run `go mod init session-17` in your `session-17` root folder
- In the `main.go` call all task's individually
- The deadline for submission is `next session date - 1 day`.

* * * * *
### Topic 1: Context Package for Cancellation

**Task 1: Cancellation with Context**

1.  Write a program that starts a long-running goroutine (e.g., a loop that prints numbers from 1 to 10, with a 1-second delay between each print).
2.  Use the `context.WithCancel` function to create a cancellable context.
3.  After 3 seconds, cancel the context and stop the goroutine using the cancellation signal.

**Expected Output:**

```go
1
2
3
(cancellation)
```

**Task 2: Context with Timeout**

1.  Create a function that simulates a task by sleeping for 5 seconds.
2.  Use `context.WithTimeout` to set a timeout of 2 seconds, so the task should be cancelled before it completes.
3.  In your program, print whether the task was completed or cancelled based on the timeout.

**Expected Output:**

```go
Task cancelled due to timeout
```

* * * * *

### Topic 2: Throttling and Rate Limiting

**Task 3: Throttling with `time.Ticker`**

1.  Create a function that prints a message, such as `"Task executed"`.
2.  Use a `time.Ticker` to call this function every 500 milliseconds, simulating throttling.
3.  After 5 executions, stop the ticker to end the program.

**Expected Output:**

```go
Task executed
Task executed
Task executed
Task executed
Task executed
(ticker stopped)
```

**Task 4: Rate Limiting with Channels**

1.  Create a function that simulates a job by printing `"Job started"` and sleeping for 1 second.
2.  Limit the execution to a maximum of 2 jobs per second by using a buffered channel with a size of 2.
3.  After 5 jobs, end the program.

**Expected Output:**

```go
Job started
Job started
Job started
Job started
Job started
(program ends)
```


## References
