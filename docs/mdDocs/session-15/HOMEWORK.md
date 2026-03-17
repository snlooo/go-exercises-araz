# Homework: Concurrency in Go

## Submission
- Create a new branch named `SES-14/concurrency-in-go`, and submit all deliverables to this branch in your private GitHub Repository.
- Create a Pull Request (PR) and send the link to the instructor for review.
- At last your repository structure should be like below:

```go
  ├───session-14
         task
         ───goroutines
              task1.go
              task2.go
         ───channels_basic_operations
              task3.go
              task4.go
         ───buffered_unbuffered_channels
              task5.go
              task6.go
         go.mod
         main.go
  └───TaskManagementSystem
```

- Run `go mod init session-14` in your `session-14` root folder
- In the `main.go` call all task's individually
- The deadline for submission is `next session date - 1 day`.

* * * * *

### Topic 1: Goroutines - Creating and Managing

**Task 1: Basic Goroutine Creation**

1.  Write a program that launches a Goroutine to print numbers from 1 to 5 with a 100ms delay between each number.
2.  The main function should print a message, wait for 1 second, and then end the program.

**Expected Output:**


```go
Main function started
1
2
3
4
5
Main function ended
```

*Note: Since Goroutines run concurrently, output order may vary slightly.*

* * * * *

**Task 2: Multiple Goroutines**

1.  Create a program that launches two Goroutines:
    -   The first Goroutine prints the alphabet from `A` to `E` with a 200ms delay between each letter.
    -   The second Goroutine prints numbers from `1` to `5` with a `300ms` delay.
2.  The main function should print `Main finished` and wait for 2 seconds before ending.

**Expected Output:**

```go
A
1
B
C
2
D
3
E
4
5
Main finished
```

*Note: Order may vary slightly since the two Goroutines run concurrently.*

* * * * *

### Topic 2: Channels - Basic Operations

**Task 3: Sending and Receiving from an Unbuffered Channel**

1.  Write a program that:
    -   Creates an unbuffered channel of type `int`.
    -   Launches a Goroutine that sends a value (e.g., 42) into the channel after `500ms`.
    -   The main function should receive the value from the channel and print it.

**Expected Output:**


```go
Received value: 42
```

* * * * *

**Task 4: Channel Range Iteration**

1.  Write a program that:
    -   Creates an unbuffered channel of type `int`.
    -   Launches a Goroutine that sends values from `1` to `5` into the channel and closes it.
    -   The main function should receive and print the values from the channel using a `for` loop.

**Expected Output:**

```go
1
2
3
4
5
Channel closed
```

* * * * *

### Topic 3: Buffered vs. Unbuffered Channels

**Task 5: Buffered Channel Operations**

1.  Write a program that:
    -   Creates a buffered channel of size 3.
    -   Sends 3 values (`10, 20, 30`) into the channel from the main function without waiting for a Goroutine to receive them.
    -   Launches a Goroutine to receive and print the values from the channel.

**Expected Output:**

```go
Sent values into buffered channel
10
20
30
All values received
```

* * * * *

**Task 6: Unbuffered Channel Blocking Example**

1.  Write a program that:
    -   Creates an unbuffered channel of type `string`.
    -   In the main function, send a value (`"Hello"`) into the channel without launching a Goroutine to receive it.
2.  Observe that the program will block (deadlock) because there is no receiver.

**Expected Outcome:** 
The program will block and fail with a runtime error: 
```go
fatal error: all goroutines are asleep - deadlock!
```

## References
- [Concurrency is not Parallelism by Rob Pike](https://www.youtube.com/watch?v=oV9rvDllKEg)
- [Understanding the Go Scheduler and discovering how it works](https://medium.com/@sanilkhurana7/understanding-the-go-scheduler-and-looking-at-how-it-works-e431a6daacf)
- [Channel Axioms](https://dave.cheney.net/2014/03/19/channel-axioms)
- [Exploring the Depths of Golang Channels: A Comprehensive Guide](https://medium.com/@ravikumar19997/exploring-the-depths-of-golang-channels-a-comprehensive-guide-53e1a97cafe6)