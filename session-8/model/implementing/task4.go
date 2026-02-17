package implementing

import (
	"fmt"
)

//Task 4: Using Interfaces with Functions
//
//Define an interface Notifier with a method Notify(message string).
//Implement this interface for two types: EmailNotifier and SMSNotifier.
//EmailNotifier should print "Sending email notification: {message}".
//SMSNotifier should print "Sending SMS notification: {message}".
//Write a function that accepts the Notifier interface and sends a notification.

type Notifier interface {
	Notify(message string)
}

type EmailNotifier struct{}

type SMSNotifier struct{}

func (EmailNotifier) Notify(message string) {
	fmt.Println("Sending email notification:", message)
}

func (SMSNotifier) Notify(message string) {
	fmt.Println("Sending SMS notification:", message)
}
