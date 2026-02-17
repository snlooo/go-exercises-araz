package implementing

import (
	"fmt"
)

// Task 3: Implementing Interfaces with Struct Methods
//
// Define an interface PaymentProcessor with a method ProcessPayment(amount float64).
// Implement this interface for two different payment types: CreditCard and PayPal.
// For CreditCard, the method should print "Processing credit card payment of {amount}".
// For PayPal, the method should print "Processing PayPal payment of {amount}".
// Write a program that processes both types of payments using the same interface method.

type PaymentProcessor interface {
	ProcessPayment(amount float64)
}

type CreditCard struct{}

type Paypal struct{}

func (c CreditCard) ProcessPayment(amount float64) {
	fmt.Printf("Processing credit card payment of %v\n", amount)
}

func (p Paypal) ProcessPayment(amount float64) {
	fmt.Printf("Processing PayPal payment of %v\n", amount)
}
