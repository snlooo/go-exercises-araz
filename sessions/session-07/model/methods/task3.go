package methods

// Task 3: Method to Calculate Circle Area
//
// Define a Circle struct with a single field:
//
// Radius (float64)
// Write a method on the Circle struct called Area() that calculates and returns the area of the circle.
//
// In the main function, create a Circle instance and call the Area() method, printing the result.
//
// (Hint: Use the formula Area = π * radius^2, where π = 3.14159)
type calculateArea func(float64) float64

type Circle struct {
	Radius float64
	Area   calculateArea
}

const Pi = 3.14159265358979323846
