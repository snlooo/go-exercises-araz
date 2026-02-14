package main

import "fmt"

func main() {
	//Create an integer array of length 5 and initialize it with values [10, 20, 30, 40, 50].
	//Write a program that:
	//Prints the first and last elements of the array.
	//Calculates and prints the sum of all elements in the array.
	//Algo task optional: Reverses the array in place (without using an additional array) and prints it.
	arrInteger := [5]int{10, 20, 30, 40, 50}
	processArr(arrInteger)
	reverseArr(arrInteger)
}

func processArr(numbers [5]int) {

	fmt.Printf("First element: %d\n", numbers[0])
	fmt.Printf("Last element: %d\n", numbers[4])

	sum := 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Printf("Sum: %d\n", sum)
}

func reverseArr(numbers [5]int) {

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] < numbers[j] {
				temp := numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = temp
			}
		}
	}

	fmt.Printf("Reversed array: %v", numbers)

}
