package main

import (
	"fmt"
	"session-14/task/benchmarking"
	"session-14/task/table_driven"
	"session-14/task/unit_testing"
)

func main() {

	//Task 1
	result, errTak1 := unit_testing.Calculate(2, 3, "+")
	if errTak1 != nil {
		return
	}
	fmt.Println(result)

	//Task 2
	profile := unit_testing.NewUserProfile("Araz", 22)
	fmt.Println(profile)

	//Task 3
	reversedString := table_driven.ReverseString("civic±")
	fmt.Println(reversedString)

	//Task 4
	temperature, errTask4 := table_driven.ConvertTemperature(100, "C")
	if errTask4 != nil {
		fmt.Println(errTask4)
	}
	fmt.Println(temperature)

	//Task 5
	var arr = []int{33, 3, 25, 7, 9, 76, 89}
	benchmarking.BubbleSort(arr)

	//Task 6
	var str = []string{"Araz", "is", "a", "Gopher "}
	fmt.Println(benchmarking.ConcatenateWithJoin(str))
	fmt.Println(benchmarking.ConcatenateWithPlus(str))
}
