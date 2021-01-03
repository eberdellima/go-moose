package controllers

import (
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

	user := services.RegisterUser(assertedPayload)

	ctx.JSON(http.StatusCreated, services.CreateTokenPair(user))
}
