package controllers

import (
	"go-moose/database"
	"go-moose/database/models"
	"go-moose/src/feed/inputs"
	"go-moose/src/feed/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// BookmarkImage creates a new user image relation
func BookmarkImage(ctx *gin.Context) {

	isBookmarked, _ := ctx.Get("is_bookmarked")
	assertedIsBookmarked := isBookmarked.(bool)

	if assertedIsBookmarked {
		ctx.AbortWithStatusJSON(http.StatusNoContent, nil)
		return
	}

	image, _ := ctx.Get("image")
	assertedImage := image.(models.Image)

	userToken, _ := ctx.Get("user_token")
	assertedUserToken := userToken.(models.UserToken)

	services.BookmarkImage(&assertedUserToken.User, &assertedImage)

	ctx.JSON(http.StatusNoContent, nil)
}

// RemoveImageBookmark removes bookmark relationship
// between an image and a user
func RemoveImageBookmark(ctx *gin.Context) {

	isBookmarked, _ := ctx.Get("is_bookmarked")
	assertedIsBookmarked := isBookmarked.(bool)

	if !assertedIsBookmarked {
		ctx.AbortWithStatusJSON(http.StatusNoContent, nil)
		return
	}

	image, _ := ctx.Get("image")
	assertedImage := image.(models.Image)

	userToken, _ := ctx.Get("user_token")
	assertedUserToken := userToken.(models.UserToken)

	services.RemoveBookmarkImage(&assertedUserToken.User, &assertedImage)

	ctx.JSON(http.StatusNoContent, nil)
}

// ModifyImageTags updates the tags associated
// with an image
func ModifyImageTags(ctx *gin.Context) {

	image, _ := ctx.Get("image")
	assertedImage := image.(models.Image)

	imageTags, _ := ctx.Get("image_tags")
	assertedImageTags := imageTags.(inputs.ImageTags)

	services.UpdateImageTags(&assertedImage, &assertedImageTags)

	ctx.JSON(http.StatusNoContent, nil)
}

// RemoveUploadedImage deletes an image
// that was uploaded by the current user
func RemoveUploadedImage(ctx *gin.Context) {

	imageID, _ := strconv.Atoi(ctx.Param("id"))

	database.DB.Delete(&models.Image{Model: gorm.Model{ID: uint(imageID)}})

	ctx.JSON(http.StatusNoContent, nil)
}
