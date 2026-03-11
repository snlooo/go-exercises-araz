package table_driven

//Task 3: Table-Driven Test for a String Reversal Function
//
//Implement a function that reverses a string.
//
//Write a table-driven test that tests multiple cases, including:
//
//An empty string.
//A palindrome.
//A string with special characters.

func ReverseString(s string) string {
	runes := []rune(s)

	var reverse []rune
	for i := len(runes) - 1; i >= 0; i-- {
		reverse = append(reverse, runes[i])
	}

	return string(reverse)
}
