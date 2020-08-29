package controllers

import (
	"go-moose/database/models"

	"github.com/gin-gonic/gin"
)

// DownloadImage servers the requested image
// file to the client
func DownloadImage(ctx *gin.Context) {

	image, _ := ctx.Get("image")
	assertedImage := image.(models.Image)

	ctx.FileAttachment(assertedImage.UploadPath, assertedImage.OriginalName)
	return
}
