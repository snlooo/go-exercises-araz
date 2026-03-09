package main

import (
	"fmt"
)

func main() {
	//Write a program that counts the frequency of each word in a given slice of strings.
	//
	//Input: []string{"go", "java", "go", "python", "go", "java"}
	//Output: map[string]int{"go": 3, "java": 2, "python": 1}
	//Once the map is built, iterate through it and print each word with its count.

	words := []string{"go", "java", "go", "python", "go", "java"}
	count := make(map[string]int)

	for _, word := range words {
		count[word]++
	}

	fmt.Println(count)

}
