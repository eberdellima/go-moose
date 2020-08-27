package services

import (
	"go-moose/database"
	"go-moose/database/models"
)

// Logout removes access with the received token
// by deleting it
func Logout(accessToken string) {

	userToken := models.UserToken{AccessToken: accessToken}
	database.DB.Delete(&userToken)
}
