package user

import (
	AuthorizationMiddlewares "go-moose/src/authorization/middlewares"
	"go-moose/src/user/controllers"
	"go-moose/src/user/validators"

	"github.com/gin-gonic/gin"
)

// ConfigureRoutes configure routes for user module
func ConfigureRoutes(router *gin.Engine) {

	router.GET("/me",
		AuthorizationMiddlewares.CheckJWT(),
		controllers.GetProfile,
	)

	router.PATCH("/me",
		AuthorizationMiddlewares.CheckJWT(),
		validators.ValidateProfileInput(),
		controllers.UpdateProfile,
	)

	router.DELETE("/me",
		AuthorizationMiddlewares.CheckJWT(),
		controllers.DeleteUser,
	)

	router.GET("/me/uploads",
		AuthorizationMiddlewares.CheckJWT(),
		validators.ValidateRequestPagination(),
		controllers.GetUploadedImages,
	)

	router.GET("/me/bookmarks",
		AuthorizationMiddlewares.CheckJWT(),
		validators.ValidateRequestPagination(),
		controllers.GetBookmarkedImages,
	)
}
