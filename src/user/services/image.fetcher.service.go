package services

import (
	"encoding/json"
	"go-moose/database"
	"go-moose/database/models"
	"go-moose/src/user/inputs"
	"go-moose/src/user/utils"
	"time"
)

// GetUploadedImages returns list of uploaded images
// by the current user
func GetUploadedImages(user *models.User, paginator *inputs.Paginator) *utils.PaginatedImageResults {

	var totalResults uint

	database.DB.Table("images").Where(models.Image{User: *user}).Count(&totalResults)

	var images []*models.Image

	database.DB.Where(models.Image{User: *user}).Offset(paginator.From - 1).Limit(paginator.Size).Find(&images)

	return &utils.PaginatedImageResults{
		PaginationResult: utils.PaginationResult{
			TotalResults: totalResults,
			From:         paginator.From,
			Size:         paginator.Size,
		},
		Results: mapImagesToUrls(images, user),
	}
}

// GetBookmarkedImages returns list of bookmarked images
// by the current user
func GetBookmarkedImages(user *models.User, paginator *inputs.Paginator) *utils.PaginatedImageResults {

	totalResults := database.DB.Model(user).Association("BookmarkedImages").Count()

	database.DB.Preload("BookmarkedImages").Offset(paginator.From - 1).Limit(paginator.Size).First(user)
	return &utils.PaginatedImageResults{
		PaginationResult: utils.PaginationResult{
			TotalResults: uint(totalResults),
			From:         paginator.From,
			Size:         paginator.Size,
		},
		Results: mapImagesToUrls(user.BookmarkedImages, user),
	}
}

// mapImagesToUrls returns list of url for each of the images provided
func mapImagesToUrls(images []*models.Image, user *models.User) []utils.ImageInfo {

	imageInfos := []utils.ImageInfo{}

	for _, image := range images {

		var tags []string

		json.Unmarshal(image.Tags, &tags)

		imageInfos = append(imageInfos, utils.ImageInfo{
			ID:         image.ID,
			Name:       image.Name,
			UploadedBy: user.Username,
			UploadedAt: image.CreatedAt.Format(time.RFC3339),
			URL:        GenerateImageURL(image),
			Tags:       tags,
		})
	}

	return imageInfos
}
