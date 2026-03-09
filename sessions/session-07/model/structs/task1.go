package main

import "fmt"

// Book Task 1: Simple Struct for a Book
//
// Define a Book struct with the following fields:
//
// Title (string)
// Author (string)
// Pages (int)
// Write a program that:
//
// Creates a Book instance with sample data.
// Prints the book's details in a readable format.
type Book struct {
	Title  string
	Author string
	Pages  int
}

func main() {

	_book := Book{
		Title:  "The Go Programming Language",
		Author: "Alan A. A. Donovan",
		Pages:  380,
	}

	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d", _book.Title, _book.Author, _book.Pages)

}
