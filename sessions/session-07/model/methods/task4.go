package methods

type Average func([]int) int

// Student Task 4: Method to Calculate and Compare Student Grades
//
// Define a Student struct with the following fields:
//
// Name (string)
// Grades (slice of int)
// Write a method Average() on the Student struct to calculate the student's average grade.
//
// In the main function, create two Student instances and compare their average grades. Print which student has a higher average.
type Student struct {
	Name    string
	Grades  []int
	Average Average
}
