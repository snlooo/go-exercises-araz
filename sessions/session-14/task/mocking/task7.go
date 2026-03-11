package mocking

//Task 7: Mocking for a Database Interaction
//
//Implement a function that interacts with a database interface to retrieve a user by ID.
//
//Mock the database interaction to test the function without requiring an actual database. The test should:
//
//Confirm that the correct user is returned for a valid ID.
//Confirm that an error is returned for an invalid ID.

type User struct {
	Id   int
	Name string
}

type Database interface {
	GetUserByID(id int) (User, error)
}

type UserService struct {
	repo Database
}

func NewUserService(repo Database) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUsername(id int) (User, error) {
	return s.repo.GetUserByID(id)
}
