package embedding

import "fmt"

//Define two interfaces: Mover with a method Move(), and Sayer with a method Say().
//Create a new interface Robot that embeds both Mover and Sayer.
//Implement the Robot interface with a struct AutoBot that:
//Prints "Moving forward" for the Move() method.
//Prints "I am a robot" for the Say() method.

type Mover interface {
	Move()
}

type Sayer interface {
	Say()
}

type Robot interface {
	Mover
	Sayer
}

type AutoBot struct{}

func (a AutoBot) Move() {
	fmt.Println("Moving forward")
}
func (a AutoBot) Say() {
	fmt.Println("I am a robot")
}
