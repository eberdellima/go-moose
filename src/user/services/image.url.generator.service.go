package services

import (
	"fmt"
	"go-moose/database/models"
	"os"

	"github.com/joho/godotenv"
)

// GenerateImageURL returns url for fetching the image provided
func GenerateImageURL(image *models.Image) string {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading env file content")
	}

	return fmt.Sprintf("%s%s/%s", os.Getenv("BACKEND_URL"), os.Getenv("IMAGE_SERVE_ROUTE"), image.Name)
}
