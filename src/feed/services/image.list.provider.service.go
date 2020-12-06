package services

import (
	"go-moose/database"
	"go-moose/database/models"
	"go-moose/src/feed/inputs"
	UserInputs "go-moose/src/user/inputs"
	UserServices "go-moose/src/user/services"
	UserUtils "go-moose/src/user/utils"
)

// GetImageList returns list of mapped images
// based on the query params provided
func GetImageList(imageListParams *inputs.ImageListQueryingParams, paginator *UserInputs.Paginator) *UserUtils.PaginatedImageResults {

	var images []*models.Image

	db := database.DB

	if imageListParams.SearchKeyword != "" {
		db = db.Where("original_name LIKE '%" + imageListParams.SearchKeyword + "%' OR JSON_SEARCH(tags, 'one', '" + imageListParams.SearchKeyword + "' ) is not null")
	}

	if imageListParams.OrderBy != "" {
		db = db.Order(imageListParams.OrderBy + " " + imageListParams.OrderDir)
	}

	var totalResults int
	db.Table("images").Count(&totalResults)

	db.Preload("User").Limit(paginator.Size).Offset(paginator.From - 1).Find(&images)

	return &UserUtils.PaginatedImageResults{
		PaginationResult: UserUtils.PaginationResult{
			TotalResults: uint(totalResults),
			From:         paginator.From,
			Size:         paginator.Size,
		},
		Results: UserServices.MapImagesToUrls(images),
	}
}
