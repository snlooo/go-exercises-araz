package main

import "fmt"

type Vehicle struct {
	Make  string
	Model string
	Year  int
}

type Car struct {
	Vehicle
	FuelType string
}

func main() {
	//Task 6: Vehicle and Car Struct Embedding
	//
	//Define a Vehicle struct with the following fields:
	//Make (string)
	//Model (string)
	//Year (int)
	//Define a Car struct that embeds the Vehicle struct and adds:
	//FuelType (string)
	//Create an instance of Car and print all the details of the car.

	_car := Car{
		Vehicle: Vehicle{
			Make:  "Kia",
			Model: "Model",
			Year:  2022,
		},
		FuelType: "Gasoline",
	}

	fmt.Printf("Make: %s\n", _car.Make)
	fmt.Printf("Model: %s\n", _car.Model)
	fmt.Printf("Year : %d\n", _car.Year)
	fmt.Printf("FuelType: %s\n", _car.FuelType)

}
