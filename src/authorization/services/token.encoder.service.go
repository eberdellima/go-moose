package services

import (
	"go-moose/database"
	"go-moose/database/models"
	"go-moose/src/authorization/utils"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
)

// CreateAccessToken create new access token for user
func CreateAccessToken(user models.User) string {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

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

// ValidateAccessToken validates if access token is correct
func ValidateAccessToken(token string) (*jwt.Token, error) {

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		err := godotenv.Load()
		if err != nil {
			panic(err)
		}

		key := []byte(os.Getenv("APP_SECRET"))
		return key, nil
	})

	return parsedToken, err
}

// CreateRefreshToken create new refresh token for user
func CreateRefreshToken() string {

	newUUID := uuid.NewV4()
	return newUUID.String()
}

// addTokenPair add new token pair for a user in db
func addTokenPair(user models.User, tokenPair utils.TokenPair) {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading env file content")
	}

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

// UpdateTokenPair update the token pairs for a user
func UpdateTokenPair(user models.User, tokenPair utils.TokenPair) {

	var userToken models.UserToken

	database.DB.Where(&models.UserToken{UserID: user.ID, RefreshToken: tokenPair.RefreshToken}).First(&userToken)

	if (&userToken) == nil {
		panic("Token pair not exists for provided user with id: " + string(user.ID))
	}

	userToken.AccessToken = CreateAccessToken(user)

	database.DB.Save(&userToken)
}
