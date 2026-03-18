package main

import (
	"fmt"
	"session-13/json"
	"session-13/reading_writing"
	"session-13/xml"
)

func main() {

	//Task 1
	errTask1 := reading_writing.ReadFile("./reading_writing/data.csv")
	if errTask1 != nil {
		fmt.Println(errTask1)
	}

	//Task 2
	errTask2 := reading_writing.ReadStory("./reading_writing/story.txt")
	if errTask2 != nil {
		fmt.Println(errTask2)
	}

	//Task 3
	errTask3 := json.ReadJson("./json/students.json")
	if errTask3 != nil {
		fmt.Println(errTask3)
	}

	//Task 4
	errTask4 := json.ReadConfig("./json/config.json")
	if errTask4 != nil {
		fmt.Println(errTask4)
	}

	//Task 5
	errTask5 := xml.ReadXml("./xml/employees.xml")
	if errTask5 != nil {
		fmt.Println(errTask5)
	}

	//Task 6
	errTask6 := xml.ReadConfigXml("mydb", "admin", "secret", true, 100)
	if errTask6 != nil {
		fmt.Println(errTask6)
	}

}
