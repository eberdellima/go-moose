package services

import (
	"go-moose/database"
	"go-moose/database/models"
	"go-moose/src/authorization/utils"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// CreateAccessToken create new access token for user
func CreateAccessToken(user models.User) string {

	tokenExpTime, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXP_IN_MIN"))
	if err != nil {
		panic(err)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"email":    user.Email,
		"iat":      time.Now(),
		"exp":      time.Now().Add(time.Second * 60 * time.Duration(tokenExpTime)),
	}

	key := []byte(os.Getenv("APP_SECRET"))
	signedToken, err := token.SignedString(key)
	if err != nil {
		panic(err)
	}

	return signedToken
}

// CreateRefreshToken create new refresh token for user
func CreateRefreshToken() string {

	newUUID := uuid.NewV4()
	return newUUID.String()
}

// addTokenPair add new token pair for a user in db
func addTokenPair(user models.User, tokenPair utils.TokenPair) {

	refreshTokenExpTime, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXP_IN_MIN"))
	if err != nil {
		panic(err)
	}

	database.DB.Create(&models.UserToken{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
		UserID:       user.ID,
		ExpiresAt:    time.Now().Add(time.Second * 60 * time.Duration(refreshTokenExpTime)),
	})
}

// CreateTokenPair create new token pair for user
func CreateTokenPair(user models.User) utils.TokenPair {

	tokenPair := utils.TokenPair{
		AccessToken:  CreateAccessToken(user),
		RefreshToken: CreateRefreshToken(),
	}

	addTokenPair(user, tokenPair)

	return tokenPair
}
