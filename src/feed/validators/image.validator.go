package validators

import (
	"go-moose/src/feed/inputs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ValidateImageTags validates whether the tags
// to be applied to an image are valid
func ValidateImageTags() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var imageTags inputs.ImageTags

		if err := ctx.ShouldBindJSON(&imageTags); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.Set("image_tags", imageTags)
		ctx.Next()
	}
}
