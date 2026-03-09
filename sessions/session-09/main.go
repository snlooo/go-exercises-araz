package main

import (
	"fmt"
	"session-9/task"
	"session-9/using_reflection"
)

func main() {

	//Task 1
	task.IdentifyTypeAndKind(42)
	task.IdentifyTypeAndKind("Hello")
	task.IdentifyTypeAndKind([]int{1, 2, 3})

	//Task 2
	p := task.Person{Name: "Ali", Age: 30}
	task.InspectStruct(p)

	//Task 3

	p2 := &using_reflection.Person{Name: "Ali", Age: 30, City: "Baku"}
	using_reflection.SetFieldValue(p2, "City", "Ganja")
	fmt.Printf("%+v\n", p2)

	//Task 4
	p3 := using_reflection.Person1{Name: "Ali", Age: 30}
	using_reflection.InvokeMethod(p3, "Greet")

}
