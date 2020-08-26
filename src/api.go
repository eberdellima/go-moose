package src

import (
	AuthorizationModule "go-moose/src/authorization"
	UserModule "go-moose/src/user"

	"github.com/gin-gonic/gin"
)

// ConfigureAPI configure all available routes of the API
func ConfigureAPI() *gin.Engine {

	router := gin.Default()

	AuthorizationModule.ConfigureRoutes(router)

	UserModule.ConfigureRoutes(router)

	return router
}
