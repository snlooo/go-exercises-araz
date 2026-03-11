package unit_testing

//Task 2: Unit Test with Assertions on Struct
//
//Create a function that returns a struct representing a user profile. The function should take inputs like name and age.
//
//Write unit tests that check the struct's values against expected results.
//If the age is below 18, the function should set a flag IsMinor to true.

type UserProfile struct {
	Name    string
	Age     int
	IsMinor bool
}

func NewUserProfile(name string, age int) UserProfile {

	var user UserProfile
	user.Name = name
	user.Age = age

	if user.Age < 18 {
		user.IsMinor = true
	}

	return user
}
