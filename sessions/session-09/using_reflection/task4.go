package using_reflection

import (
	"fmt"
	"reflect"
)

//Task 4: Accessing Method Details
//
//Extend the Person struct with a method Greet() string that returns a greeting message like "Hello, I am Alice."
//
//Write a function InvokeMethod(value interface{}, methodName string) that:
//
//Uses reflection to invoke the specified method on the struct.
//Prints the returned value.

type Person1 struct {
	Name string
	Age  int
}

func (p Person1) Greet() {

	fmt.Println("Hello, I am Alice.")
}

func InvokeMethod(value interface{}, methodName string) {

	method := reflect.ValueOf(value).MethodByName(methodName)
	method.Call(nil)

}
