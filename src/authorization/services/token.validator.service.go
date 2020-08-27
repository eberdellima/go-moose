package services

import (
	"go-moose/src/authorization/utils"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

// ValidateAccessToken validates if access token is correct
func ValidateAccessToken(token string) (jwt.MapClaims, error) {

	// Parse token correctly
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		err := godotenv.Load()
		if err != nil {
			panic(err)
		}

		key := []byte(os.Getenv("APP_SECRET"))
		return key, nil
	})

	if err != nil {
		return nil, &utils.TokenValidationError{}
	}

	// Assert token claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, &utils.TokenValidationError{}
	}

	// Check for expiration
	exp := string(claims["exp"].(string))
	expTime, err := time.Parse(time.RFC3339, exp)

	if expTime.Before(time.Now()) {
		return nil, &utils.TokenValidationError{}
	}

	return claims, err
}
