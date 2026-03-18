package xml

import (
	"encoding/xml"
	"fmt"
	"os"
)

//Task 5: XML Reader and Data Extraction
//
//Create an XML file named employees.xml containing a list of employees. Each employee should have name, position, and salary fields.
//Write a program that reads the XML file and extracts only those employees whose salary is above a specified threshold (e.g., salary > 50000).

type Employee struct {
	Name     string `xml:"name"`
	Position string `xml:"position"`
	Salary   int    `xml:"salary"`
}

type Employees struct {
	Employees []Employee `xml:"employee"`
}

func ReadXml(path string) error {

	fmt.Println("#================================#")

	inputFile, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var employees Employees
	unMarshalError := xml.Unmarshal(inputFile, &employees)
	if unMarshalError != nil {
		return err
	}

	fmt.Println("Employees with Salary above 50000:")

	for _, employee := range employees.Employees {
		if employee.Salary > 50000 {
			fmt.Printf("- %s , %s", employee.Name, employee.Position)
		}
	}

	return nil
}
