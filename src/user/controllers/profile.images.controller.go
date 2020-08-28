package controllers

import (
	"go-moose/database"
	"go-moose/database/models"
	"go-moose/src/user/inputs"
	"go-moose/src/user/services"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"

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

// RemoveUploadedImage deletes an image
// that was uploaded by the current user
func RemoveUploadedImage(ctx *gin.Context) {

	imageID, _ := strconv.Atoi(ctx.Param("id"))

	database.DB.Delete(&models.Image{Model: gorm.Model{ID: uint(imageID)}})

	ctx.JSON(http.StatusNoContent, nil)
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
