package feed

import (
	AuthorizationMiddlewares "go-moose/src/authorization/middlewares"
	"go-moose/src/feed/controllers"
	"go-moose/src/feed/middlewares"
	"go-moose/src/feed/validators"
	UserValidators "go-moose/src/user/validators"

	"github.com/gin-gonic/gin"
)

// ConfigureRoutes configure routes for feed module
func ConfigureRoutes(router *gin.Engine) {

	router.POST("/bookmarks/:id",
		AuthorizationMiddlewares.CheckJWT(),
		middlewares.CheckImageExists(),
		middlewares.IsAlreadyBookmarkedByUser(),
		controllers.BookmarkImage,
	)

	router.DELETE("/bookmarks/:id",
		AuthorizationMiddlewares.CheckJWT(),
		middlewares.CheckImageExists(),
		middlewares.IsAlreadyBookmarkedByUser(),
		controllers.RemoveImageBookmark,
	)

	router.GET("/download/:id",
		AuthorizationMiddlewares.CheckJWT(),
		middlewares.CheckImageExists(),
		middlewares.CheckImageFileExists(),
		controllers.DownloadImage,
	)

	router.GET("/images",
		AuthorizationMiddlewares.CheckJWT(),
		UserValidators.ValidateRequestPagination(),
		middlewares.CheckImageListQueryingParams(),
		controllers.ListImages,
	)

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
