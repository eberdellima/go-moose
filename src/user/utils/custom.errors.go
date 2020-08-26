package utils

// EmailIsTakenError represents error
// when user provided an existsing email
type EmailIsTakenError struct{}

func (e *EmailIsTakenError) Error() string {
	return "Email is already taken"
}

// UsernameIsTakenError represents error
// when user provided an existing username
type UsernameIsTakenError struct{}

func (e *UsernameIsTakenError) Error() string {
	return "Username is already taken"
}

// IncorrectPasswordError represents error
// when user provided an incorrect password
type IncorrectPasswordError struct{}

func (e *IncorrectPasswordError) Error() string {
	return "Incorrect password provided"
}
