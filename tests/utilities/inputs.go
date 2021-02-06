package utilities

import (
	"go-moose/database/models"
	"go-moose/src/authorization/inputs"
	"strconv"
	"time"
)

// CreateRegistrationInput creates an inputs.RegisterInput object
// Used for testing purposes only
func CreateRegistrationInput() inputs.RegisterInput {

	return inputs.RegisterInput{
		Email:    "sample@" + strconv.Itoa(time.Now().Nanosecond()) + ".com",
		Username: "sample" + strconv.Itoa(time.Now().Nanosecond()),
		Password: "sample" + strconv.Itoa(time.Now().Nanosecond()),
	}
}

// CreateLoginInput creates an inputs.LoginInput object
// Used for testing purposes only
func CreateLoginInput(email, password string) inputs.LoginInput {

	return inputs.LoginInput{
		Email:    email,
		Password: password,
	}
}

// CreateRefreshTokenInput creates an inputs.RefreshTokenInput object
// Used for testing purposes only
func CreateRefreshTokenInput(user models.User) inputs.RefreshTokenInput {

	tokenPair := CreateTokenPair(user)

	return inputs.RefreshTokenInput{
		RefreshToken: tokenPair.RefreshToken,
		UserID:       user.ID,
	}
}
