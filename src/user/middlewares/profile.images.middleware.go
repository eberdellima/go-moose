package middlewares

import (
	"fmt"
	"go-moose/database"
	"go-moose/database/models"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

// CheckUploadedImageExists checks if the image with the provided ID
// actually exists and if it was uploaded by the current user.
// Otherwise. it returns `Not found` error code
func CheckUploadedImageExists() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		userToken, _ := ctx.Get("user_token")
		assertedUserToken := userToken.(models.UserToken)

		imageID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("DELETE %s not found", ctx.FullPath())})
			return
		}

		var image models.Image

		database.DB.Where(models.Image{Model: gorm.Model{ID: uint(imageID)}, User: assertedUserToken.User}).First(&image)

		if image.ID == 0 {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.Next()
	}
}
