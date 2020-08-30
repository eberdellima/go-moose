package authorization

import (
	"go-moose/src/authorization/controllers"
	"go-moose/src/authorization/middlewares"
	"go-moose/src/authorization/validators"

	"github.com/gin-gonic/gin"
)

// ConfigureRoutes configure routes for authorization module
func ConfigureRoutes(router *gin.Engine) {

	router.POST("/login",
		validators.ValidateLoginInput(),
		middlewares.CheckUserCredentials(),
		controllers.Login,
	)

	router.POST("/logout",
		middlewares.CheckJWT(),
		controllers.Logout,
	)

	router.POST("/register",
		validators.ValidateRegistrationInput(),
		middlewares.IsUserRegistered(),
		controllers.Register,
	)

	router.PATCH("/refresh-token",
		validators.ValidateRefreshTokenInput(),
		middlewares.CheckTokenExists(),
		controllers.RefreshToken,
	)
}
