package controllers

import (
	"go-moose/database/models"
	"go-moose/src/user/inputs"
	"go-moose/src/user/services"
	"go-moose/src/user/utils"
	"net/http"

	AuthorizationServices "go-moose/src/authorization/services"

	"github.com/gin-gonic/gin"
)

// GetProfile returns basic user related info
// abiding to ProfileInfo struct
func GetProfile(ctx *gin.Context) {

	userToken, _ := ctx.Get("user_token")
	assertedUserToken := userToken.(models.UserToken)

	user := assertedUserToken.User

	ctx.JSON(http.StatusOK, utils.ProfileInfo{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	})
}

// UpdateProfile modifies user profile in accordance
// with the data provided
func UpdateProfile(ctx *gin.Context) {

	userToken, _ := ctx.Get("user_token")
	assertedUserToken := userToken.(models.UserToken)

	payload, _ := ctx.Get("payload")
	assertedPayload, _ := payload.(inputs.ProfileInput)

	err := services.UpdateUserProfile(&assertedUserToken.User, assertedPayload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenPair := AuthorizationServices.CreateTokenPair(assertedUserToken.User)
	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  tokenPair.AccessToken,
		"refresh_token": tokenPair.RefreshToken,
	})
}
