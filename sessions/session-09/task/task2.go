package task

import (
	"fmt"
	"reflect"
)

//Task 2: Inspecting Struct Fields
//
//Define a struct Person with fields Name (string) and Age (int).
//
//Write a function InspectStruct(value interface{}) that:
//
//Accepts any struct.
//Uses reflection to print the names, types, and values of all the struct's fields.

type Person struct {
	Name string
	Age  int
}

func InspectStruct(value interface{}) {

	typeOfStruct := reflect.TypeOf(value)
	valueOfFields := reflect.ValueOf(value)

	for i := 0; i < typeOfStruct.NumField(); i++ {
		fmt.Printf("Field Name: %v, Type: %v, Value: %v\n", typeOfStruct.Field(i).Name, typeOfStruct.Field(i).Type, valueOfFields.Field(i))
	}

}
