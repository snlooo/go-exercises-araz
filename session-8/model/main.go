package main

import (
	"fmt"
	"model/assertion"
	"model/defining"
	"model/embedding"
	"model/implementing"
	"model/switching"
)

func main() {

	//Task 1
	var circle defining.Circle
	var rectangle defining.Rectangle

	circle.SetRadius(5)
	rectangle.SetHeightWidth(5, 10)

	fmt.Println("Circle area:", defining.Shape(circle).Area())
	fmt.Println("Rectangle area:", defining.Shape(rectangle).Area())

	//------------------------------------------------------------------------//

	//Task 2
	var r defining.Speaker = defining.Dog{}

	if value, ok := r.(defining.Dog); ok {
		fmt.Println("Dog says:", value.Speak())
	} else if value1, ok1 := r.(defining.Person); ok1 {
		fmt.Println("Person says:", value1.Speak())
	}

	//-------------------------------------------------------------------------//

	//Task 3
	var PayPal implementing.Paypal
	var creditCard implementing.CreditCard
	implementing.PaymentProcessor(creditCard).ProcessPayment(100)
	implementing.PaymentProcessor(PayPal).ProcessPayment(75.5)

	//--------------------------------------------------------------------------//

	//Task 4
	SendNotification(implementing.Notifier(implementing.SMSNotifier{}))

	//--------------------------------------------------------------------------//

	//Task 5

	var intType assertion.EmptyInterfaceAny = 42
	var stringType assertion.EmptyInterfaceAny = "GoLang"
	var floatType assertion.EmptyInterfaceAny = 3.1415

	value, _ := intType.(int)
	fmt.Println("Value is of type int:", value)
	value1, _ := stringType.(string)
	fmt.Println("Value is of type string:", value1)
	value2, _ := floatType.(float64)
	fmt.Println("Value is of type int:", value2)

	//--------------------------------------------------------------------------//

	//Task 6

	var _emptyInterface switching.EmptyInterface = 1000 //100 //Hi, Gophers //true //struct {}

	switch v := _emptyInterface.(type) {
	case int:
		fmt.Println("Type is int: ", v)
	case string:
		fmt.Println("Type is string: ", v)
	case bool:
		fmt.Println("Type is bool: ", v)
	default:
		fmt.Println("Unknown type")
	}

	//--------------------------------------------------------------------------//

	//Task 7

	var bot embedding.AutoBot

	embedding.Robot(bot).Move()
	embedding.Robot(bot).Say()

	//--------------------------------------------------------------------------//
}

func SendNotification(n implementing.Notifier) {

	if value, ok := n.(implementing.SMSNotifier); ok {
		implementing.Notifier(value).Notify("Your item has shipped")
	} else if value1, ok1 := n.(implementing.EmailNotifier); ok1 {
		implementing.Notifier(value1).Notify("Your item has shipped")
	}
}
