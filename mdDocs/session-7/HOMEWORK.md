# Homework: Session 7 - Structs and Methods in Go

## Objective
In this session - Understanding Structs and Methods in Go. The tasks gradually build up knowledge, from basic to more complex.

## Submission
- Create a new branch named `SES-07/structs-methods`, and submit all deliverables to this branch in your private GitHub Repository.
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
         model
         ───structs
              task1.go
              task2.go
         ───methods
              task3.go
              task4.go
         ───embedding
              task5.go
              task6.go
         go.mod
         main.go
  └───TaskManagementSystem

  ```

- The deadline for submission is **next session date - 1 day**.



## Instructions:

### Defining and Using Structs

**Task 1: Simple Struct for a Book**

1.  Define a `Book` struct with the following fields:

    -   `Title` (string)
    -   `Author` (string)
    -   `Pages` (int)
2.  Write a program that:

    -   Creates a `Book` instance with sample data.
    -   Prints the book's details in a readable format.

**Expected Output Example:**
```go
Title: The Go Programming Language
Author: Alan A. A. Donovan
Pages: 380
```

**Task 2: Struct for a Rectangle**

1.  Define a `Rectangle` struct with the following fields:
    -   `Width` (float64)
    -   `Height` (float64)
2.  Write a program that:
    -   Creates a `Rectangle` instance with sample data.
    -   Calculates and prints the area and perimeter of the rectangle.
    
**Expected Output:**
```go
Width: 10.5, Height: 5.0
Area: 52.5
Perimeter: 31.0
```

### Methods on Structs

**Task 3: Method to Calculate Circle Area**

1.  Define a `Circle` struct with a single field:
    -   `Radius` (float64)
2.  Write a method on the `Circle` struct called `Area()` that calculates and returns the area of the circle.

3.  In the `main` function, create a `Circle` instance and call the `Area()` method, printing the result.

(Hint: Use the formula Area = π * radius^2, where π = 3.14159)

**Expected Output:**
```go
Circle Radius: 7.0
Area: 153.938
```

**Task 4: Method to Calculate and Compare Student Grades**

1.  Define a `Student` struct with the following fields:
    -   `Name` (string)
    -   `Grades` (slice of int)
2.  Write a method `Average()` on the `Student` struct to calculate the student's average grade.

3.  In the `main` function, create two `Student` instances and compare their average grades. Print which student has a higher average.

**Expected Output:**
```go
Student 1: Ali, Average Grade: 85
Student 2: Vali, Average Grade: 90
Vali has a higher average grade.

```

### Embedding Structs

**Task 5: Embedding a Person Struct**
-   Define a `Person` struct with the following fields:
    -   `FirstName` (string)
    -   `LastName` (string)
-   Define an `Employee` struct that embeds the `Person` struct and adds:
    -   `EmployeeID` (int)
    -   `Position` (string)
-   Create an instance of `Employee` and print both the person's full name and the employee's details.

**Expected Output:**
```go
Name: Ali Aliyev
Employee ID: 12345
Position: Software Engineer
```


**Task 6: Vehicle and Car Struct Embedding**

1.  Define a `Vehicle` struct with the following fields:
    -   `Make` (string)
    -   `Model` (string)
    -   `Year` (int)
2.  Define a `Car` struct that embeds the `Vehicle` struct and adds:
    -   `FuelType` (string)
3.  Create an instance of `Car` and print all the details of the car.

**Expected Output:**
```go
Make: Kia
Model: K5
Year: 2022
Fuel Type: Gasoline

```


## References
- Next to Maps (video): [Data Structures in Golang - Hash Tables](https://www.youtube.com/watch?v=JCXLwfKMWOU)
- [Builder Pattern in Go](https://hackernoon.com/go-design-patterns-an-introduction-to-builder)
- [Composition over Inheritance in Go](https://golangbot.com/inheritance/)
- [Methods in Go](https://go101.org/article/method.html)
- [Using Methods in Go](https://www.practical-go-lessons.com/chap-14-methods)
- [Difference between Function and Methods in Golang](https://medium.com/@ravikumarray92/difference-between-function-and-methods-in-golang-986fc16b5912)
- [Golang - Struct Memory Optimizasyonu](https://medium.com/trendyol-tech/golang-struct-memory-optimizasyonu-bb79ffa556e5)
- [Optimize Your Struct with Alignment Trick](https://blog.devtrovert.com/p/struct-optimization-a-small-change)