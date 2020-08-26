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
// abidding to ProfileInfo struct
func GetProfile(ctx *gin.Context) {

	userToken, _ := ctx.Get("user_token")
	assertedUserToken := userToken.(models.UserToken)

	userInfo := utils.ProfileInfo{
		ID:        assertedUserToken.User.ID,
		Username:  assertedUserToken.User.Username,
		Email:     assertedUserToken.User.Email,
		CreatedAt: assertedUserToken.User.CreatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":         userInfo.ID,
		"username":   userInfo.Username,
		"email":      userInfo.Email,
		"created_at": userInfo.CreatedAt,
	})
}

// UpdateProfile modifies user profile in accordance
// with the data provided
func UpdateProfile(ctx *gin.Context) {

	userToken, _ := ctx.Get("user_token")
	assertedUserToken := userToken.(models.UserToken)

	user := &assertedUserToken.User

	payload, _ := ctx.Get("payload")
	assertedPayload, _ := payload.(inputs.ProfileInput)

	err := services.UpdateUserProfile(user, assertedPayload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenPair := AuthorizationServices.CreateTokenPair(*user)
	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  tokenPair.AccessToken,
		"refresh_token": tokenPair.RefreshToken,
	})
}
