package controllers

import (
	"go-moose/database"
	"go-moose/database/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteUser soft deletes a user. Since a user should have
// a field CreatedAt it actually sets this field to the current time
// instead of deletting the user record. The "deleted" user is not considered
// when querying for data
func DeleteUser(ctx *gin.Context) {

	userToken, _ := ctx.Get("user_token")
	assertedUserToken := userToken.(models.UserToken)

	user := &assertedUserToken.User

	database.DB.Delete(user)

	ctx.JSON(http.StatusNoContent, nil)
}
