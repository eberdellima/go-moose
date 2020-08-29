package services

import (
	"encoding/json"
	"go-moose/database"
	"go-moose/database/models"
	"go-moose/src/feed/inputs"
)

// UpdateImageTags modifies tags of an image
// based on the tags array provided
func UpdateImageTags(image *models.Image, imageTags *inputs.ImageTags) {

	marshaledImageTags, err := json.Marshal(imageTags.Tags)
	if err != nil {
		panic(err)
	}

	image.Tags = string(marshaledImageTags)

	database.DB.Save(image)
}
