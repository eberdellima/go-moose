package services

import (
	"go-moose/database"
	"go-moose/database/models"
	"go-moose/src/authorization/inputs"
)

func RegisterUser(userInput inputs.RegisterInput) models.User {

	password, err := GeneratePassword(userInput.Password)
	if err != nil {
		panic(err)
	}

	user := models.User{
		Email:    userInput.Email,
		Password: password,
		Username: userInput.Username,
	}

	database.DB.Create(&user)

	return user
}
