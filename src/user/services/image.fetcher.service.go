package services

import (
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

	database.DB.Table("images").Where(models.Image{User: *user}).Count(&totalResults).QueryExpr()

	var images []models.Image

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

	var totalResults uint

	database.DB.Model(&models.Image{}).Related(user, "BookmarkedBy").Where(models.Image{User: *user}).Count(&totalResults).QueryExpr()

	var bookmarkedImages []models.Image

	database.DB.Model(&models.Image{}).Related(user, "BookmarkedBy").Where(models.Image{User: *user}).Offset(paginator.From - 1).Limit(paginator.Size).Find(&bookmarkedImages)

	return &utils.PaginatedImageResults{
		PaginationResult: utils.PaginationResult{
			TotalResults: totalResults,
			From:         paginator.From,
			Size:         paginator.Size,
		},
		Results: mapImagesToUrls(bookmarkedImages, user),
	}
}

// mapImagesToUrls returns list of url for each of the images provided
func mapImagesToUrls(images []models.Image, user *models.User) []utils.ImageInfo {

	imageInfos := []utils.ImageInfo{}

	for _, image := range images {

		imageInfos = append(imageInfos, utils.ImageInfo{
			ID:         image.ID,
			Name:       image.Name,
			UploadedBy: user.Username,
			UploadedAt: image.CreatedAt.Format(time.RFC3339),
			URL:        GenerateImageURL(&image),
		})
	}

	return imageInfos
}
