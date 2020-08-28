package feed

import (
	AuthorizationMiddlewares "go-moose/src/authorization/middlewares"
	"go-moose/src/user/controllers"
	"go-moose/src/user/middlewares"

	"github.com/gin-gonic/gin"
)

// ConfigureRoutes configure routes for feed module
func ConfigureRoutes(router *gin.Engine) {

	router.DELETE("/images/:id",
		AuthorizationMiddlewares.CheckJWT(),
		middlewares.CheckUploadedImageExists(),
		controllers.RemoveUploadedImage,
	)
}
