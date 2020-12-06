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

	services.Logout(token[7:])

	ctx.JSON(http.StatusNoContent, nil)
}
