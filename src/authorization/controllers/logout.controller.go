package controllers

import (
	"go-moose/src/authorization/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Logout removes current token pair for the
// provided user
func Logout(ctx *gin.Context) {

	token := ctx.GetHeader("Authorization")
	token = string(token[7:])

	services.Logout(token)

	ctx.JSON(http.StatusNoContent, nil)
}
