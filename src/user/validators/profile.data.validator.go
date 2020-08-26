package validators

import (
	"go-moose/src/user/inputs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ValidateProfileInput validates that all required data
// for updating a user's profile are provided correctly
func ValidateProfileInput() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var input inputs.ProfileInput

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.Set("payload", input)
		ctx.Next()
	}
}
