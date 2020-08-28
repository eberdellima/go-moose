package src

import (
	AuthorizationModule "go-moose/src/authorization"
	FeedModule "go-moose/src/feed"
	UserModule "go-moose/src/user"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// ConfigureAPI configure all available routes of the API
func ConfigureAPI() *gin.Engine {

	router := gin.Default()

	err := godotenv.Load()
	if err != nil {
		panic("Error loading env file content")
	}

	maxMultipartMemoryMB, err := strconv.Atoi(os.Getenv("MAX_MULTIPART_MEMORY_MB"))
	if err != nil {
		panic(err)
	}

	router.MaxMultipartMemory = int64(maxMultipartMemoryMB)
	router.Static(os.Getenv("IMAGE_SERVE_ROUTE"), os.Getenv("UPLOAD_DIR"))

	configureRoutes(router)

	return router
}

// configureRoutes server for registering all available routes
// into the router engine
func configureRoutes(router *gin.Engine) {

	AuthorizationModule.ConfigureRoutes(router)

	UserModule.ConfigureRoutes(router)

	FeedModule.ConfigureRoutes(router)
}
