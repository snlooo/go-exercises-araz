package defining

//Task 2: Multiple Implementations of the Same Interface
//
//Create an interface Speaker with a method Speak() string.
//Implement this interface for two different types: Dog and Person.
//Dog should return "Woof!".
//Person should return "Hello!".
//Write a function that takes a Speaker and calls the Speak() method.

type Speaker interface {
	Speak() string
}

type Dog struct{}
type Person struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

func (p Person) Speak() string {
	return "Hello!"
}
