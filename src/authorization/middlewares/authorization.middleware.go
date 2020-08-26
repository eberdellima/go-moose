package middlewares

import (
	"go-moose/database"
	"go-moose/database/models"
	"go-moose/src/authorization/inputs"
	"go-moose/src/authorization/services"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CheckUserCredentials check user's credentials for login
func CheckUserCredentials() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		payload, _ := ctx.Get("payload")
		assertedPayload, _ := payload.(inputs.LoginInput)

		var user models.User

		database.DB.Where(&models.User{Email: assertedPayload.Email}).First(&user)

		if user.ID == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Email or password incorrect"})
			return
		}

		isMatch, err := services.ComparePasswords(assertedPayload.Password, user.Password)
		if !isMatch || err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Email or password incorrect"})
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}

// IsUserRegistered check if user already exists
func IsUserRegistered() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		payload, _ := ctx.Get("payload")
		assertedPayload, _ := payload.(inputs.RegisterInput)

		user := models.User{}

		database.DB.Where(models.User{Email: assertedPayload.Email}).Or(models.User{Username: assertedPayload.Username}).First(&user)

		if user.ID != 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "User already exists"})
			return
		}

		ctx.Next()
	}
}

// CheckJWT check if jwt is valid
func CheckJWT() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		token := ctx.GetHeader("Authorization")
		token = string(token[7:])

		// Parse token correctly
		parsedToken, err := services.ValidateAccessToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
			return
		}

		// Assert token claims
		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok || !parsedToken.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
			return
		}

		// Check for expiration
		exp := string(claims["exp"].(string))
		expTime, err := time.Parse(time.RFC3339, exp)

		if expTime.Before(time.Now()) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
			return
		}

		// Check if token exists in db
		userID := uint(claims["user_id"].(float64))

		var userToken models.UserToken

		database.DB.Where(models.UserToken{UserID: userID, AccessToken: token}).Preload("User").First(&userToken)

		if userToken.ID == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
			return
		}

		ctx.Set("user_token", userToken)
		ctx.Next()
	}
}

// CheckTokenExists checks if token exists
// before refreshing them
func CheckTokenExists() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		payload, _ := ctx.Get("payload")
		assertedPayload, _ := payload.(inputs.RefreshTokenInput)

		var userToken models.UserToken

		database.DB.Where(models.UserToken{RefreshToken: assertedPayload.RefreshToken, UserID: assertedPayload.UserID}).First(&userToken)

		if userToken.ID == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authenticated"})
			return
		}

		ctx.Set("user_token", userToken)
		ctx.Next()
	}
}
