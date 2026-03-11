package mocking

//Task 8: Mocking an External Service
//
//Create a function that sends notifications through an interface Notifier, which has a method Send(message string) error.
//
//Mock the Notifier interface to simulate the sending of notifications. Write tests to verify that:
//
//A notification is sent successfully.
//The function returns an error if sending fails.

type Notification struct {
	message string
}

type Notifier interface {
	Send(message string) error
}

type NotifierService struct {
	repo Notifier
}

func (n *NotifierService) SendNotification(message string) error {
	return n.repo.Send(message)
}

func NewNotifierService(repo Notifier) *NotifierService {
	return &NotifierService{repo: repo}
}
