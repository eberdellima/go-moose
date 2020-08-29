package middlewares

import (
	"go-moose/database/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// CheckImageFileExists check whether the file with
// the given path does actually exist in that path
func CheckImageFileExists() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		image, _ := ctx.Get("image")
		assertedImage := image.(models.Image)

		if _, err := os.Stat(assertedImage.UploadPath); os.IsNotExist(err) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "File not found"})
			return
		}

		ctx.Next()
	}
}
