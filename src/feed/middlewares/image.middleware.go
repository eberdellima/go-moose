package middlewares

import (
	"fmt"
	"go-moose/database"
	"go-moose/database/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// CheckImageExists checks if the image with the provided ID
// actually exists and if it was uploaded by the current user.
// Otherwise. it returns `Not found` error code
func CheckImageExists() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		imageID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("DELETE %s not found", ctx.FullPath())})
			return
		}

		var image models.Image

		database.DB.Preload("User").Where(models.Image{Model: gorm.Model{ID: uint(imageID)}}).First(&image)

		if image.ID == 0 {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.Set("image", image)
		ctx.Next()
	}
}

// IsUploadedByCurrentUser checks whether the respective
// image has been uploaded by the current user
func IsUploadedByCurrentUser() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		user, _ := ctx.Get("user_token")
		assertedUserToken := user.(models.UserToken)

		image, _ := ctx.Get("image")
		assertedImage := image.(models.Image)

		if assertedUserToken.User.ID != assertedImage.User.ID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Permission denied"})
			return
		}

		ctx.Next()
	}
}
