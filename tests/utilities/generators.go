package utilities

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-moose/database"
	"go-moose/database/models"
	"go-moose/src/authorization/services"
	"go-moose/src/authorization/utils"
	"strconv"
	"time"
)

func CreateUser(optPassword string) models.User {

	password := "sample" + strconv.Itoa(time.Now().Nanosecond())

	if optPassword != "" {
		password = optPassword
	}

	hashedPassword, _ := services.GeneratePassword(password)

	user := models.User{
		Email:    "sample@" + strconv.Itoa(time.Now().Nanosecond()) + ".com",
		Username: "sample" + strconv.Itoa(time.Now().Nanosecond()),
		Password: hashedPassword,
	}

	database.DB.Create(&user)

	return user
}

func CreateTokenPair(user models.User) utils.TokenPair {
	return services.CreateTokenPair(user)
}

func CreateRequestBody(args interface{}) (*bytes.Reader, error) {

	if args == nil {
		return nil, errors.New("invalid args for request body")
	}

	jsonInput, err := json.Marshal(args)

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonInput), nil
}
