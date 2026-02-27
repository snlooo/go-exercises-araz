package authsystem

import "fmt"

//Task 5: Refactor an Existing Project to Use Custom Errors
//Task: Refactor a User Authentication System to Handle Errors Properly
//
//You are given a simplified user authentication system that does not currently use proper error handling. Refactor the code to:
//Use custom error types where appropriate.
//Wrap and unwrap errors to provide detailed information in case of failures.

//Refactor this code to:
//Create custom error types for different failure cases (UserNotFoundError and InvalidPasswordError).
//Modify the authenticate() function to return detailed errors instead of a bool.
//Wrap errors to add context at each level, and then handle these errors in the main() function by printing detailed messages.
//Refactored Code Expectations:
//
//If the user does not exist, return a UserNotFoundError.
//If the password is incorrect, return an InvalidPasswordError.
//In main(), display the proper error messages based on the returned errors.

var UserNotFoundError = Exceptions{message: "user not found: %s"}
var InvalidPasswordError = Exceptions{message: "invalid password for user: %s"}

type Exceptions struct {
	message string
}

func (e Exceptions) Error() string {
	return e.message
}

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func Authenticate(username, password string) (bool, error) {
	if passValue, ok := users[username]; ok {
		if passValue == password {
			return true, nil
		}
	} else {
		return false, fmt.Errorf(UserNotFoundError.Error(), username)
	}
	return false, fmt.Errorf(InvalidPasswordError.Error(), username)
}
