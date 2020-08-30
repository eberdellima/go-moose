package utils

import (
	"go-moose/database/models"
)

// GetImageIndex checks a slice of images
// for the index of the provided image. Returns
// -1 if the image is not in the slice
func GetImageIndex(images []*models.Image, image *models.Image) int {

	index := -1

	for i, sliceImage := range images {

		if sliceImage.ID == image.ID {
			index = i
			break
		}
	}

	return index
}
