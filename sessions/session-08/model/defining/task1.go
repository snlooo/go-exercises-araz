package defining

import "math"

//Task 1: Defining and Using Interfaces
//
//Define an interface Shape with the method Area() float64. Then, implement this interface for two structs: Circle and Rectangle.
//The Circle struct should have a field for radius.
//The Rectangle struct should have fields for width and height.
//Write a function that takes a Shape and prints the area of that shape.

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.GetRadius() * c.GetRadius()
}

// GetRadius Circle section
func (c Circle) GetRadius() float64 {
	return c.radius
}

func (c *Circle) SetRadius(radius float64) {
	c.radius = radius
}

// GetHeightWidth Rectangle section
func (r Rectangle) GetHeightWidth() (float64, float64) {
	return r.width, r.height
}

func (r *Rectangle) SetHeightWidth(width float64, height float64) {
	r.width = width
	r.height = height
}
