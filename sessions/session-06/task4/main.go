package main

import (
	"fmt"
)

var Countries = map[string]string{
	"Azerbaijan": "Baku",
	"USA":        "Washington",
	"Germany":    "Berlin",
	"Japan":      "Tokyo",
}

func main() {
	//Create a map that associates countries with their capitals, such as:
	//    map[string]string{
	//        "Azerbaijan": "Baku",
	//        "USA": "Washington",
	//        "Germany": "Berlin",
	//        "Japan": "Tokyo",
	//    }
	//Write a function that:
	//Takes the country name as input and returns the corresponding capital.
	//Returns a message if the country does not exist in the map.
	//Add two new countries to the map and print the updated map.

	for countries := range Countries {
		checkCity(countries)
	}

	Countries["Georgia"] = "Tbilisi"
	Countries["Turkey"] = "Istanbul"

	fmt.Println("Updated Map:", Countries)

}

func checkCity(countryName string) {

	value := Countries[countryName]

	if _, ok := Countries[countryName]; !ok {
		value = "is not found"
	}

	fmt.Printf("Capital of %s is %s\n", countryName, value)

}
