package mocking

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockNotifRepo struct {
	mock.Mock
}

func (n *MockNotifRepo) Send(message string) error {
	args := n.Called(message)
	return args.Error(0)
}

func TestNewNotifierService(t *testing.T) {

	repo := new(MockNotifRepo)
	repo.On("Send", mock.Anything).Return(nil)
	notifier := NewNotifierService(repo)
	err := notifier.SendNotification("Hello World")
	assert.Nil(t, err)

}

func TestNewNotifierServiceError(t *testing.T) {
	repo := new(MockNotifRepo)
	repo.On("Send", mock.Anything).Return(errors.New("error"))
	notifier := NewNotifierService(repo)
	err := notifier.SendNotification("Hello World")
	assert.NotNil(t, err)
}
