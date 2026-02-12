package main

import "fmt"

// Rectangle Task 2: Struct for a Rectangle
//
// Define a Rectangle struct with the following fields:
// Width (float64)
// Height (float64)
// Write a program that:
// Creates a Rectangle instance with sample data.
// Calculates and prints the area and perimeter of the rectangle.
type Rectangle struct {
	Width  float64
	Height float64
}

func main() {
	_rectangle := Rectangle{Width: 5.0, Height: 10.5}
	area := _rectangle.Height * _rectangle.Width
	perimeter := 2 * (_rectangle.Height + _rectangle.Width)

	fmt.Printf("Height: %.2f Width: %.2f\n", _rectangle.Height, _rectangle.Width)
	fmt.Printf("Area: %.2f\n", area)
	fmt.Printf("Perimeter: %.2f", perimeter)

}
