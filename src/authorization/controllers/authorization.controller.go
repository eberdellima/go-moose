package controllers

import (
	"go-moose/database"
	"go-moose/database/models"
	"go-moose/src/authorization/inputs"
	"go-moose/src/authorization/services"
	"go-moose/src/authorization/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Login provides used with token pair
// to access the features of the system
func Login(ctx *gin.Context) {

	user, _ := ctx.Get("user")
	assertedUser, _ := user.(models.User)

	tokenPair := services.CreateTokenPair(assertedUser)
	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  tokenPair.AccessToken,
		"refresh_token": tokenPair.RefreshToken,
	})
}

// Logout removes current token pair for the
// provided user
func Logout(ctx *gin.Context) {

	token := ctx.GetHeader("Authorization")
	token = string(token[7:])

	userToken := models.UserToken{AccessToken: token}

	database.DB.Delete(&userToken)

	ctx.JSON(http.StatusNoContent, nil)
}

// Register store user's data in the db and provides
// user with token pair
func Register(ctx *gin.Context) {

	payload, _ := ctx.Get("payload")
	assertedPayload, _ := payload.(inputs.RegisterInput)

	password, err := services.GeneratePassword(assertedPayload.Password)
	if err != nil {
		panic(err)
	}

	user := models.User{
		Email:    assertedPayload.Email,
		Password: password,
		Username: assertedPayload.Username,
	}

	database.DB.Create(&user)

	tokenPair := services.CreateTokenPair(user)
	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  tokenPair.AccessToken,
		"refresh_token": tokenPair.RefreshToken,
	})
}

// RefreshToken provides user with a new access token
// if refresh token is still valid. Otherwise genereates
// a new token pair
func RefreshToken(ctx *gin.Context) {

	userToken, _ := ctx.Get("user_token")
	assertedUserToken := userToken.(models.UserToken)

	// Generate new access token and refresh token
	if assertedUserToken.ExpiresAt.Before(time.Now()) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	// Generate only new access token
	accessToken := services.CreateAccessToken(assertedUserToken.User)

	tokenPair := utils.TokenPair{AccessToken: accessToken, RefreshToken: assertedUserToken.RefreshToken}

	services.UpdateTokenPair(assertedUserToken.User, tokenPair)

	ctx.JSON(http.StatusOK, map[string]string{
		"access_token":  accessToken,
		"refresh_token": assertedUserToken.RefreshToken,
	})
}
