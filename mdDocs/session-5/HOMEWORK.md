# Homework: Session 5 - Understanding Pointers in Go

## Objective
In this session - pointers in Go, homework is touching on the concepts of "understanding pointers," "pointers vs. values," and "using pointers with functions." The tasks gradually build up their knowledge, from basic to more complex.

## Submission
- Create a new branch named `SES-05/pointers`, and submit all deliverables to this branch in your private GitHub Repository.
- Create a Pull Request (PR) and send the link to the instructor for review.
- At last your repository structure should be like below:

  ```bash
  ├───session-1
  ├───session-2
  ├───session-3
  ├───session-4
  ├───session-5
          task1.go
          task2.go
          task3.go
          task4.go
  └───TaskManagementSystem
          go.mod
          main.go
  ```

- The deadline for submission is **next session date - 1 day**.



## Instructions:

#### 1\. Understanding Pointers

-   **Task 1: Create a Go program that demonstrates how to declare, assign, and dereference a pointer.**
    -   Declare an integer variable `x` and assign it a value of 10.
    -   Declare a pointer variable that holds the memory address of `x`.
    -   Print the value of `x`, the address of `x`, and the value stored at the pointer.

**Expected Output Example:**
```
Value of x: 10
Address of x: 0x20818a220
Value at pointer: 10
```


#### 2\. Pointers vs. Values

-   **Task 2: Write a Go program that illustrates the difference between passing a value and passing a pointer to a function.**
    -   Create two functions:
        -   `incrementByValue(val int)` -- this should take an integer argument and increment it by 1.
        -   `incrementByPointer(ptr *int)` -- this should take a pointer to an integer and increment the value it points to by 1.
    -   In the `main` function:
        -   Declare an integer `x` and assign it a value of 5.
        -   Call both functions with `x` and observe the difference in behavior.

**Expected Output Example:**

```
Value of x before incrementing by value: 5
Value of x after incrementing by value: 5
Value of x before incrementing by pointer: 5
Value of x after incrementing by pointer: 6
```

#### 3\. Using Pointers with Functions

-   **Task 3: Create a function `swap` that swaps the values of two integer variables using pointers.**
    -   Define a function `swap(a *int, b *int)` that swaps the values stored at the addresses passed in.
    -   In the `main` function:
        -   Declare two integer variables `x` and `y` with values 10 and 20.
        -   Call the `swap` function and print the values of `x` and `y` before and after swapping.

**Expected Output Example:**

```
Before swapping: x = 10, y = 20
After swapping: x = 20, y = 10
```


### Bonus Task (Optional):

-   **Task 4: Implement a function `doubleVariadicElements(numbers ...*int)` that takes a variadic parameter (a list of integer pointers) and doubles the value of each element.**
    -   In the `main` function, declare a set of integer variables, pass their addresses (pointers) to the function, and print the values before and after modification.

### Task 4: Doubling Variadic Elements with Pointers

-   **Step 1**: Define a function `doubleVariadicElements(numbers ...*int)` that accepts a variadic number of pointers to integers.
-   **Step 2**: Inside the function, iterate through the pointers and modify the value they point to by doubling it.
-   **Step 3**: In the `main` function, declare a few integer variables, pass their addresses (pointers) to the `doubleVariadicElements` function, and print the values before and after modification.

**Expected Output Example:**

```
Before doubling: 1 2 3
After doubling: 2 4 6
```
