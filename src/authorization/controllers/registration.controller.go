package controllers

import (
	"go-moose/database"
	"go-moose/database/models"
	"go-moose/src/authorization/inputs"
	"go-moose/src/authorization/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
