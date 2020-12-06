package services

import (
	"encoding/json"
	"fmt"
	"go-moose/database"
	"go-moose/database/models"
	"os"
	"path/filepath"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
)

// UploadImage save image data in the db.
// Returns the ID of the newly saved image and
// the path where it was uploaded to
func UploadImage(filename string, user *models.User) (uint, string) {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	uploadPath := generateNewFilename(os.Getenv("UPLOAD_DIR"), filepath.Ext(filename))

	var tags []string
	marshalledTags, _ := json.Marshal(tags)

	image := models.Image{
		OriginalName: filename,
		Name:         filepath.Base(uploadPath),
		User:         *user,
		Tags:         marshalledTags,
		UploadPath:   uploadPath,
	}

	database.DB.Create(&image)

	return image.ID, uploadPath
}

// RemoveUploadedImageOnScanningError should be called only when
// image was stored successfully on db but an error didn't allow it to be stored
// in the filesystem as well
func RemoveUploadedImageOnScanningError(imageID uint) {

	database.DB.Delete(&models.Image{Model: gorm.Model{ID: imageID}})
}

// generateNewFilename creates the new name for the image
// being uploaded
func generateNewFilename(uploadDir string, mimetype string) string {

	newFilename := uuid.NewV4()

	return fmt.Sprintf("%s/%s%s", uploadDir, newFilename.String(), mimetype)
}
