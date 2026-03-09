Homework: Generics

## Submission
- Create a new branch named `SES-10/generics`, and submit all deliverables to this branch in your private GitHub Repository.
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
         task
         ───go_generics
              task1.go
              task2.go
              task3.go
              task4.go
         go.mod
         main.go
  └───TaskManagementSystem
```

- Run `go mod init session-10` in your `session-10` root folder
- In the `main.go` call all task's individually
- The deadline for submission is `next session date - 1 day`.

* * * * *

#### **Topic: Generics in Go**

**Task 1: Generic Function for Sum**

1.  Write a generic function `Sum[T int | float64](slice []T) T` that:

    -   Calculates and returns the sum of elements in a slice.

**Sample Input:**

```go
    fmt.Println(Sum([]int{1, 2, 3, 4}))
    fmt.Println(Sum([]float64{1.5, 2.5, 3.5}))
```

**Expected Output:**

```go
    Sum: 10
    Sum: 7.5
```

* * * * *

**Task 2: Generic Min and Max**

1.  Write a generic function `MinMax[T int | float64](slice []T) (T, T)` that:

    -   Finds the minimum and maximum values in a slice.

**Sample Input:**

```go
    min, max := MinMax([]int{10, 20, 5, 8, 15})
    fmt.Println(min, max)
```

**Expected Output:**

```go
    Min: 5, Max: 20
```

* * * * *

**Task 3: Generic Map Function**

1.  Write a generic function `Map[T any, U any](slice []T, fn func(T) U) []U` that:

    -   Applies a given function to each element of a slice and returns a new slice with the results.

**Sample Input:**

```go
    ints := []int{1, 2, 3, 4}
    doubled := Map(ints, func(n int) int { return n * 2 })
    fmt.Println(doubled)

    words := []string{"hello", "world"}
    lengths := Map(words, func(s string) int { return len(s) })
    fmt.Println(lengths)
```

**Expected Output:**

```go
    [2 4 6 8]
    [5 5]
```

* * * * *

**Task 4: Generic Filter Function**

1.  Write a generic function `Filter[T any](slice []T, predicate func(T) bool) []T` that:

    -   Returns a new slice containing only the elements that satisfy the predicate function.

**Sample Input:**

```go
    nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    evens := Filter(nums, func(n int) bool { return n%2 == 0 })
    fmt.Println(evens)

    names := []string{"Ali", "Aysel", "Tural", "Anar"}
    aNames := Filter(names, func(s string) bool { return s[0] == 'A' })
    fmt.Println(aNames)
```

**Expected Output:**

```go
    [2 4 6 8 10]
    [Ali Aysel Anar]
```

***

## References
- [Golang Generics](https://codilime.com/blog/golang-generics/)
- [Generics implementation](https://deepsource.com/blog/go-1-18-generics-implementation)
- [When to use Generics?](https://go.dev/blog/when-generics)
