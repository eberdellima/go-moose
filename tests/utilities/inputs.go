package utilities

import (
	"go-moose/database/models"
	"go-moose/src/authorization/inputs"
	"strconv"
	"time"
)

func CreateRegistrationInputs() inputs.RegisterInput {

	return inputs.RegisterInput{
		Email:    "sample@" + strconv.Itoa(time.Now().Nanosecond()) + ".com",
		Username: "sample" + strconv.Itoa(time.Now().Nanosecond()),
		Password: "sample" + strconv.Itoa(time.Now().Nanosecond()),
	}
}

func CreateLoginInput(email string, password string) inputs.LoginInput {

	return inputs.LoginInput{
		Email:    email,
		Password: password,
	}
}

func CreateRefreshTokenInput(user models.User) inputs.RefreshTokenInput {

	tokenPair := CreateTokenPair(user)

	return inputs.RefreshTokenInput{
		RefreshToken: tokenPair.RefreshToken,
		UserID:       user.ID,
	}
}
