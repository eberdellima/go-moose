package controllers

import (
	"go-moose/database/models"
	"go-moose/src/user/inputs"
	"go-moose/src/user/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUploadedImages returns paginated list of
// images uploaded by the current user
func GetUploadedImages(ctx *gin.Context) {

	userToken, _ := ctx.Get("user_token")
	assertedUserToken := userToken.(models.UserToken)

	user := assertedUserToken.User

	paginator, _ := ctx.Get("paginator")
	assertedPaginator := paginator.(inputs.Paginator)

	ctx.JSON(http.StatusOK, services.GetUploadedImages(&user, &assertedPaginator))
}

// GetBookmarkedImages returns paginated list of
// images bookmarked by the current user
func GetBookmarkedImages(ctx *gin.Context) {

	userToken, _ := ctx.Get("user_token")
	assertedUserToken := userToken.(models.UserToken)

	user := assertedUserToken.User

	paginator, _ := ctx.Get("paginator")
	assertedPaginator := paginator.(inputs.Paginator)

	ctx.JSON(http.StatusOK, services.GetBookmarkedImages(&user, &assertedPaginator))
}
