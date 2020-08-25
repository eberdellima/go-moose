package validators

import (
	"go-moose/src/authorization/inputs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ValidateLoginInput validates payload for login
// before proceeding
func ValidateLoginInput() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var payload inputs.LoginInput

		if err := ctx.ShouldBindJSON(&payload); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.Set("payload", payload)
		ctx.Next()
	}
}

// ValidateRegistrationInput validates payload for registration
// before proceeding
func ValidateRegistrationInput() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var payload inputs.RegisterInput

		if err := ctx.ShouldBindJSON(&payload); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.Set("payload", payload)
		ctx.Next()
	}
}

// ValidateRefreshTokenInput validates payload for refreshing token
// before proceeding
func ValidateRefreshTokenInput() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var payload inputs.RefreshTokenInput

		if err := ctx.ShouldBindJSON(&payload); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.Set("payload", payload)
		ctx.Next()
	}
}
