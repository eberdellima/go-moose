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
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": http.StatusText(http.StatusNotFound)})
			return
		}

		ctx.Set("image", image)
		ctx.Next()
	}
}

// IsAlreadyBookmarkedByUser checks whether the user
// has already bookmarked the image
func IsAlreadyBookmarkedByUser() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		image, _ := ctx.Get("image")
		assertedImage := image.(models.Image)

		userToken, _ := ctx.Get("user_token")
		assertedUserToken := userToken.(models.UserToken)

		user := assertedUserToken.User

		database.DB.Preload("BookmarkedImages").First(&user)

		isBookmarked := false
		for _, image := range user.BookmarkedImages {

			if image.ID == assertedImage.ID {
				isBookmarked = true
				break
			}
		}

		ctx.Set("is_bookmarked", isBookmarked)
		ctx.Next()
	}
}

// IsUploadedByCurrentUser checks whether the respective
// image has been uploaded by the current user
func IsUploadedByCurrentUser() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		userToken, _ := ctx.Get("user_token")
		assertedUserToken := userToken.(models.UserToken)

		image, _ := ctx.Get("image")
		assertedImage := image.(models.Image)

		if assertedUserToken.User.ID != assertedImage.User.ID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Permission denied"})
			return
		}

		ctx.Next()
	}
}
