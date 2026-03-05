package json

import (
	"encoding/json"
	"fmt"
	"os"
)

type Students struct {
	Name  string
	Age   int
	Grade int
}

//Create a JSON file named students.json with an array of student objects. Each student should have fields like name, age, and grade.
//Write a program that reads this file, filters out students with grades below a specified threshold (e.g., grade < 60),
//and prints the names of the remaining students.

func ReadJson(path2JsonFIle string) error {

	fmt.Println("#================================#")

	inputReader, _ := os.ReadFile(path2JsonFIle)
	var students []Students

	unmarshalError := json.Unmarshal(inputReader, &students)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}

	fmt.Println("Students with passing grades: ")
	for _, student := range students {

		if student.Grade < 60 {

			fmt.Println("-", student.Name)
		}
	}

	return nil
}
