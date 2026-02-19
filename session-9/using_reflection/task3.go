package using_reflection

import (
	"reflect"
)

//Extend the Person struct with a field City (string).
//
//Write a function SetFieldValue(value interface{}, fieldName string, newValue interface{}) that:
//
//Accepts a pointer to a struct.
//Sets the value of a specified field to a new value using reflection.

type Person struct {
	Name string
	Age  int
	City string
}

func SetFieldValue(value interface{}, fieldName string, newValue interface{}) {

	element := reflect.ValueOf(value).Elem()
	field := element.FieldByName(fieldName)
	newVal := reflect.ValueOf(newValue)

	field.Set(newVal)
}
