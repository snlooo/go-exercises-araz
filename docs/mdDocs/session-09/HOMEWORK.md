# Session 16 Homework: Reflection and Generics

## Submission
- Create a new branch named `SES-09/reflection-and-generics`, and submit all deliverables to this branch in your private GitHub Repository.
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
         task
         ───reflection
              task1.go
              task2.go
         ───using_reflection
              task3.go
              task4.go
         go.mod
         main.go
  └───TaskManagementSystem
```

- Run `go mod init session-09` in your `session-09` root folder
- In the `main.go` call all task's individually
- The deadline for submission is `next session date - 1 day`.

* * * * *

#### **Topic 1: Understanding Reflection in Go**

**Task 1: Type and Kind Identification**

1.  Write a function `IdentifyTypeAndKind(value interface{})` that:

    -   Uses reflection to determine the type and kind of the given value.
    -   Prints the type and kind.
    -   Handles basic types (`int`, `string`, `bool`) and a slice of integers.

**Sample Input:**

```go
    IdentifyTypeAndKind(42)
    IdentifyTypeAndKind("Hello")
    IdentifyTypeAndKind([]int{1, 2, 3})
```
**Output:**

```go
    Value: 42, Type: int, Kind: int
    Value: Hello, Type: string, Kind: string
    Value: [1 2 3], Type: []int, Kind: slice
```

* * * * *

**Task 2: Inspecting Struct Fields**

1.  Define a struct `Person` with fields `Name` (string) and `Age` (int).

2.  Write a function `InspectStruct(value interface{})` that:

    -   Accepts any struct.
    -   Uses reflection to print the names, types, and values of all the struct's fields.

**Sample Input:**

```go
type Person struct {
    Name string
    Age  int
    }
    p := Person{Name: "Ali", Age: 30}
    InspectStruct(p)
```

**Expected Output:**

```go
    Field Name: Name, Type: string, Value: Alice
    Field Name: Age, Type: int, Value: 30
```

* * * * *

#### **Topic 2: Using reflect Package**

**Task 3: Modifying Struct Values**

1.  Extend the `Person` struct with a field `City` (string).

2.  Write a function `SetFieldValue(value interface{}, fieldName string, newValue interface{})` that:

    -   Accepts a pointer to a struct.
    -   Sets the value of a specified field to a new value using reflection.

**Sample Input:**

```go
    p := &Person{Name: "Ali", Age: 30, City: "Baku"}
    SetFieldValue(p, "City", "Ganja")
    fmt.Println(p)
```    

**Expected Output:**

```go
    &{Name: Ali, Age: 30, City: Ganja}
```    

* * * * *

**Task 4: Accessing Method Details**

1.  Extend the `Person` struct with a method `Greet() string` that returns a greeting message like "Hello, I am Alice."

2.  Write a function `InvokeMethod(value interface{}, methodName string)` that:

    -   Uses reflection to invoke the specified method on the struct.
    -   Prints the returned value.

**Sample Input:**

```go
    p := Person{Name: "Ali", Age: 30}
    InvokeMethod(p, "Greet")`
```

**Expected Output:**

```go
    Invoked method: Greet, Result: Hello, I am Alice.
```

***

## References
- [Reflect package in Go](https://blog.stackademic.com/reflect-package-in-go-e42eeab85bad)