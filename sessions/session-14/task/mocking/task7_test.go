package mocking

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) GetUserByID(id int) (User, error) {
	args := m.Called(id)
	return args.Get(0).(User), args.Error(1)
}

func TestGetUserService_FindById(t *testing.T) {
	repo := new(MockUserRepo)
	testData := User{
		Id:   1,
		Name: "Araz",
	}

	repo.On("GetUserByID", 1).Return(testData, nil)

	service := NewUserService(repo)
	user, err := service.GetUsername(1)
	assert.Nil(t, err)
	assert.Equal(t, "Araz", user.Name)
}

func TestGetUserService_FindByIdError(t *testing.T) {

	repo := new(MockUserRepo)

	repo.On("GetUserByID", 2).Return(User{}, assert.AnError)

	service := NewUserService(repo)

	user, err := service.GetUsername(2)

	assert.NotNil(t, err)
	assert.Equal(t, User{}, user)
}
