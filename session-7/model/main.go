package main

import (
	"fmt"
	"math"
	"model/methods"
)

func main() {

	//In the main function, create a Circle instance and call the Area() method, printing the result.
	//Task 3
	_circle := methods.Circle{
		Radius: 7.0,
		Area: func(radius float64) float64 {
			return methods.Pi * math.Pow(radius, 2)
		},
	}

	fmt.Printf("Circle Radius: %.2f\n", _circle.Radius)
	fmt.Printf("Area: %.3f\n", _circle.Area(_circle.Radius))

	fmt.Println("-------------------------------------------------------------------------")

	//In the main function, create two Student instances and compare their average grades. Print which student has a higher average.
	//Task 4
	firstStudent := methods.Student{
		Name:   "Ali",
		Grades: []int{70, 85, 90, 95},
		Average: func(grades []int) int {
			total := 0
			for _, grade := range grades {
				total += grade
			}
			return total / len(grades)
		},
	}

	secondStudent := methods.Student{
		Name:   "Vali",
		Grades: []int{80, 100, 85, 95},
		Average: func(grades []int) int {
			total := 0
			for _, grade := range grades {
				total += grade
			}
			return total / len(grades)
		},
	}

	_firstStudentAverage := firstStudent.Average(firstStudent.Grades)
	_secondStudentAverage := secondStudent.Average(secondStudent.Grades)

	fmt.Printf("Student 1: %s, Average Grade: %d\n", firstStudent.Name, _firstStudentAverage)
	fmt.Printf("Student 1: %s, Average Grade: %d\n", secondStudent.Name, secondStudent.Average(secondStudent.Grades))

	if _firstStudentAverage > _secondStudentAverage {
		fmt.Printf("%s has a higher average grade.\n", firstStudent.Name)
	} else {
		fmt.Printf("%s has a higher average grade.\n", secondStudent.Name)
	}

}
