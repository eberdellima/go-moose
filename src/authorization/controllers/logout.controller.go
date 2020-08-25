package controllers

import (
	"go-moose/database"
	"go-moose/database/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Logout removes current token pair for the
// provided user
func Logout(ctx *gin.Context) {

	token := ctx.GetHeader("Authorization")
	token = string(token[7:])

	userToken := models.UserToken{AccessToken: token}

	database.DB.Delete(&userToken)

	ctx.JSON(http.StatusNoContent, nil)
}
