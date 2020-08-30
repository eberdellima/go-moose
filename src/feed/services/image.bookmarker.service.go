package services

import (
	"go-moose/database"
	"go-moose/database/models"
	"go-moose/src/feed/utils"
)

// BookmarkImage add a new bookmarked image
// to the list of bookmarked images of that user
func BookmarkImage(user *models.User, image *models.Image) {

	user.BookmarkedImages = []*models.Image{image}

	database.DB.Save(&user)
}

// RemoveBookmarkImage removes a bookmarked image
// from the list of bookmarked images for that user
func RemoveBookmarkImage(user *models.User, image *models.Image) {

	database.DB.Preload("BookmarkedImages").First(user)

	index := utils.GetImageIndex(user.BookmarkedImages, image)

	bookmarkedImage := user.BookmarkedImages[index]

	database.DB.Model(user).Association("BookmarkedImages").Delete(bookmarkedImage)
}
