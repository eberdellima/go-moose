package services

import (
	"go-moose/database"
	"go-moose/database/models"
	"go-moose/src/authorization/utils"
	"strconv"
)

// UpdateTokenPair update the token pairs for a user
func UpdateTokenPair(user models.User, tokenPair utils.TokenPair) {

	var userToken models.UserToken

	database.DB.Where(&models.UserToken{UserID: user.ID, RefreshToken: tokenPair.RefreshToken}).First(&userToken)

	if (&userToken) == nil {
		panic("Token pair not exists for provided user with id: " + strconv.Itoa(int(user.ID)))
	}

	userToken.AccessToken = CreateAccessToken(user)

	database.DB.Save(&userToken)
}
