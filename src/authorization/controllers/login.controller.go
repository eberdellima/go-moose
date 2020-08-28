package controllers

import (
	"go-moose/database/models"
	"go-moose/src/authorization/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login provides used with token pair
// to access the features of the system
func Login(ctx *gin.Context) {

	user, _ := ctx.Get("user")
	assertedUser, _ := user.(models.User)

	ctx.JSON(http.StatusOK, services.CreateTokenPair(assertedUser))
}
