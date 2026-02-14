# Homework: Session 6 - Understanding Arrays, Slices and Map in Go

## Objective
In this session - Understanding Arrays, Slices and Map in Go, homework is touching on the concepts of arrays, slices and maps. The tasks gradually build up knowledge, from basic to more complex.

## Submission
- Create a new branch named `SES-06/array-slice-map`, and submit all deliverables to this branch in your private GitHub Repository.
- Create a Pull Request (PR) and send the link to the instructor for review.
- At last your repository structure should be like below:

  ```bash
  ├───session-1
  ├───session-2
  ├───session-3
  ├───session-4
  ├───session-5
  ├───session-6
          task1
              main.go
          task2
              main.go
          task3
              main.go
          task4
              main.go
          task5
              main.go
  └───TaskManagementSystem

  ```

- The deadline for submission is **next session date - 1 day**.



## Instructions:

### Arrays

**Task 1: Array Initialization and Basic Operations**

1.  Create an integer array of length 5 and initialize it with values `[10, 20, 30, 40, 50]`.
2.  Write a program that:
    -   Prints the first and last elements of the array.
    -   Calculates and prints the sum of all elements in the array.
    -   *Algo task optional:* Reverses the array in place (without using an additional array) and prints it.

**Expected Output Example:**
```sh
-   First element: 10
-   Last element: 50
-   Sum: 150
-   Reversed array: [50, 40, 30, 20, 10]
```
### Slices

**Task 2: Slice Creation and Resizing**

1.  Create a slice of integers from an array of 10 elements: `[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]`.
2.  Perform the following operations on the slice:
    -   Extract the slice from index `2` to `6` and print it.
    -   Append values `10, 11, 12` to the slice and print the result.
    -   Show the effect of `append()` on the underlying array by printing both the slice and the original array after appending.

**Expected Output:**
```
Sub-slice: [2, 3, 4, 5]
Appended slice: [2, 3, 4, 5, 10, 11, 12]
Print original array: `The original array remains unaffected` or `modified based on capacity`.
```

**Task 3: Slice Capacity and Growth**

1.  Create an empty slice using the `make()` function with an initial length of `3` and capacity of `5`.
2.  Append values to the slice until its length exceeds `5`.
3.  At each step, print the `length` and `capacity` of the slice to show how Go manages memory allocation internally.

### Maps

**Task 4: Working with Maps**

1.  Create a map that associates countries with their capitals, such as:

```
    map[string]string{
        "Azerbaijan": "Baku",
        "USA": "Washington",
        "Germany": "Berlin",
        "Japan": "Tokyo",
    }
```

2.  Write a `function` that:
    -   Takes the country name as `input` and `returns` the corresponding capital.
    -   Returns a message if the country does not exist in the map.
3.  Add two new countries to the map and print the updated map.

**Expected Output:**
```
Capital of Azerbaijan is Baku
Capital of USA is Washington
Capital of Germany is Berlin
Capital of Japan is Tokyo
Capital of Turkey is Ankara
Capital of Italy is Rome
`Capital of France is not found`
```

**Task 5: Counting Frequency of Words**

1.  Write a program that counts the frequency of each word in a given slice of strings.

    -   Input: `[]string{"go", "java", "go", "python", "go", "java"}`
    -   Output: `map[string]int{"go": 3, "java": 2, "python": 1}`
2.  Once the map is built, iterate through it and print each word with its count.

## References

- [HashTable implementation in Go](https://medium.com/kalamsilicon/hash-tables-implementation-in-go-48c165c54553)
- [How the Go runtime implements maps efficiently](https://dave.cheney.net/2018/05/29/how-the-go-runtime-implements-maps-efficiently-without-generics)
- [Go Slice Growth logic in Golang source](https://github.com/golang/go/blob/2f507985dc24d198b763e5568ebe5c04d788894f/src/runtime/slice.go#L289)