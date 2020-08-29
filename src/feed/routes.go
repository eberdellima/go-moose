package feed

import (
	AuthorizationMiddlewares "go-moose/src/authorization/middlewares"
	"go-moose/src/feed/controllers"
	"go-moose/src/feed/middlewares"
	"go-moose/src/feed/validators"

	"github.com/gin-gonic/gin"
)

// ConfigureRoutes configure routes for feed module
func ConfigureRoutes(router *gin.Engine) {

	router.GET("/download/:id",
		AuthorizationMiddlewares.CheckJWT(),
		middlewares.CheckImageExists(),
		middlewares.CheckImageFileExists(),
		controllers.DownloadImage,
	)

	// TODO: Add route to bookmark an image
	// router.POST("/images/:id/bookmark")

	// TODO: Add route to get list of images
	// router.GET("/images")

	router.POST("/images",
		AuthorizationMiddlewares.CheckJWT(),
		controllers.UploadImage,
	)

	router.PATCH("/images/:id",
		AuthorizationMiddlewares.CheckJWT(),
		middlewares.CheckImageExists(),
		middlewares.IsUploadedByCurrentUser(),
		validators.ValidateImageTags(),
		controllers.ModifyImageTags,
	)

	router.DELETE("/images/:id",
		AuthorizationMiddlewares.CheckJWT(),
		middlewares.CheckImageExists(),
		middlewares.IsUploadedByCurrentUser(),
		controllers.RemoveUploadedImage,
	)
}
