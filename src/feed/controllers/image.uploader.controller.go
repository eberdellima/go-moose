package controllers

import (
	"go-moose/database/models"
	"go-moose/src/feed/services"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// UploadImage saves an image in the filesystem
func UploadImage(ctx *gin.Context) {

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filename := filepath.Base(file.Filename)

	userToken, _ := ctx.Get("user_token")
	assertedUserToken := userToken.(models.UserToken)

	imageID, fullFilePath := services.UploadImage(filename, &assertedUserToken.User)

	if err := ctx.SaveUploadedFile(file, fullFilePath); err != nil {

		services.RemoveUploadedImageOnScanningError(imageID)

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"image_id": imageID})
}
