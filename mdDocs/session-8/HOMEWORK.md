# Homework: Session 8: Introduction to Interfaces in Go

## Objective
In this session - Understanding Interfaces in Go. The tasks gradually build up knowledge, from basic to more complex.

## Submission
- Create a new branch named `SES-08/interfaces`, and submit all deliverables to this branch in your private GitHub Repository.
- Create a Pull Request (PR) and send the link to the instructor for review.
- At last your repository structure should be like below:

  ```bash
  ├───session-1
  ├───session-2
  ├───session-3
  ├───session-4
  ├───session-5
  ├───session-6
  ├───session-7
  ├───session-8  
         model
         ───defining
              task1.go
              task2.go
         ───implementing
              task3.go
              task4.go
         ───assertion
              task5.go
         ───switching
              task6.go
         ───embedding
              task7.go
         go.mod
         main.go
  └───TaskManagementSystem

  
  ```
  
  1. Run `go mod init session-8` in your root folder
  2. In the `main.go` call all task's individually
  

- The deadline for submission is **next session date - 1 day**.



### Homework 1: Understanding Interfaces in Go

**Task 1: Defining and Using Interfaces**

1.  Define an interface `Shape` with the method `Area() float64`. Then, implement this interface for two structs: `Circle` and `Rectangle`.
    -   The `Circle` struct should have a field for `radius`.
    -   The `Rectangle` struct should have fields for `width` and `height`.
    -   Write a function that takes a `Shape` and prints the area of that shape.

**Expected Output:**

```go
Circle area: 78.53981633974483
Rectangle area: 50
```


**Task 2: Multiple Implementations of the Same Interface**

1.  Create an interface `Speaker` with a method `Speak() string`.
2.  Implement this interface for two different types: `Dog` and `Person`.
    -   `Dog` should return `"Woof!"`.
    -   `Person` should return `"Hello!"`.
3.  Write a function that takes a `Speaker` and calls the `Speak()` method.

**Expected Output:**

```go
Dog says: Woof!
Person says: Hello!
```


### Homework 2: Implementing Interfaces

**Task 3: Implementing Interfaces with Struct Methods**

1.  Define an interface `PaymentProcessor` with a method `ProcessPayment(amount float64)`.
2.  Implement this interface for two different payment types: `CreditCard` and `PayPal`.
    -   For `CreditCard`, the method should print `"Processing credit card payment of {amount}"`.
    -   For `PayPal`, the method should print `"Processing PayPal payment of {amount}"`.
3.  Write a program that processes both types of payments using the same interface method.

**Expected Output:**

```go
Processing credit card payment of 100
Processing PayPal payment of 75.5
```

**Task 4: Using Interfaces with Functions**

1.  Define an interface `Notifier` with a method `Notify(message string)`.
2.  Implement this interface for two types: `EmailNotifier` and `SMSNotifier`.
    -   `EmailNotifier` should print `"Sending email notification: {message}"`.
    -   `SMSNotifier` should print `"Sending SMS notification: {message}"`.
3.  Write a function that accepts the `Notifier` interface and sends a notification.

**Expected Output:**

```go
Sending email notification: Your item has shipped
Sending SMS notification: Your item has shipped
```

* * * * *

### Homework 3: Type Assertions and Type Switches

**Task 5: Type Assertions**

1.  Write a program where you define an empty interface `any` and assign values of different types (e.g., `int`, `string`, `float64`) to it.
2.  Use type assertions to detect the underlying types and print different messages depending on the type.

**Expected Output:**

```go
Value is of type int: 42
Value is of type string: GoLang
Value is of type float64: 3.1415
```

* * * * *

### Homework 4: Type Switches

**Task 6: Type Switch**

1.  Write a program that takes an `interface{}` and uses a type switch to determine whether the underlying value is an `int`, `string`, or `bool`.
    -   Print a specific message for each type.
    -   If the type is not recognized, print `"Unknown type"`.

**Expected Output:**

```go
Type is int: 100
Type is string: Hi, Gophers
Type is bool: true
Unknown type
```

* * * * *

### Homework 5: Embedding Interfaces

**Task 7: Interface Embedding**

1.  Define two interfaces: `Mover` with a method `Move()`, and `Sayer` with a method `Say()`.
2.  Create a new interface `Robot` that embeds both `Mover` and `Sayer`.
3.  Implement the `Robot` interface with a struct `AutoBot` that:
    -   Prints `"Moving forward"` for the `Move()` method.
    -   Prints `"I am a robot"` for the `Say()` method.

**Expected Output:**

```go
Moving forward
I am a robot
```

## References
- [Gopherfest 2015 | Go Proverbs with Rob Pike](https://www.youtube.com/watch?v=PAAkCSZUG1c)
- [Краш-курс по интерфейсам в Go](https://habr.com/ru/articles/276981/)
- [Интерфейсы в Go — как красиво выстрелить себе в ногу](https://habr.com/ru/articles/597461/)