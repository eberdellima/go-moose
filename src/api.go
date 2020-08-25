package src

import (
	AuthorizationModule "go-moose/src/authorization"

	"github.com/gin-gonic/gin"
)

// ConfigureAPI configure all available routes of the API
func ConfigureAPI() *gin.Engine {

	router := gin.Default()

	AuthorizationModule.ConfigureRoutes(router)

	return router
}
