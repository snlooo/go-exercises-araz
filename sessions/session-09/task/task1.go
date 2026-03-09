package task

import (
	"fmt"
	"reflect"
)

// IdentifyTypeAndKind Task 1: Type and Kind Identification
//
//Write a function IdentifyTypeAndKind(value interface{}s) that:
//
//Uses reflection to determine the type and kind of the given value.
//Prints the type and kind.
//Handles basic types (int, string, bool) and a slice of integers.

func IdentifyTypeAndKind(value interface{}) {
	fmt.Printf("Value: %v, Type: %v, Kind: %v\n", value, reflect.TypeOf(value), reflect.TypeOf(value).Kind())
}
