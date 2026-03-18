package table_driven

import "errors"

//Task 4: Table-Driven Test for Temperature Conversion
//
//Implement a function that converts temperatures between Celsius and Fahrenheit.
//
//Write a table-driven test that checks for various cases, such as:
//
//0°C to 32°F.
//100°C to 212°F.
//-40°C to -40°F.

func ConvertTemperature(value float64, scale string) (float64, error) {
	// Add logic for Celsius to Fahrenheit and vice versa

	switch scale {
	case "C":
		return (value * 9 / 5) + 32, nil
	case "F":
		return (value - 32) * 5 / 9, nil
	default:
		return 0, errors.New("invalid temperature scale")
	}

}
