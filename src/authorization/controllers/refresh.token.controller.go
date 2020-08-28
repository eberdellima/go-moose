package controllers

import (
	"go-moose/database/models"
	"go-moose/src/authorization/services"
	"go-moose/src/authorization/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

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

	ctx.JSON(http.StatusOK, tokenPair)
}
