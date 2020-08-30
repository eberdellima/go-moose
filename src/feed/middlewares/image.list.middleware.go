package middlewares

import (
	"errors"
	"go-moose/src/feed/inputs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckImageListQueryingParams check the received values for
// querying the image list and checks whether they are valid or not
func CheckImageListQueryingParams() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var imageListQueryParams inputs.ImageListQueryingParams

		if err := ctx.ShouldBindQuery(&imageListQueryParams); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := validateOrderDir(imageListQueryParams.OrderDir); err != nil {
			imageListQueryParams.OrderDir = ""
		}

		if err := validateOrderImageBy(imageListQueryParams.OrderBy); err != nil {
			imageListQueryParams.OrderBy = ""
			imageListQueryParams.OrderDir = ""
		}

		if imageListQueryParams.OrderBy != "" && imageListQueryParams.OrderDir == "" {
			imageListQueryParams.OrderDir = "ASC"
		}

		imageListQueryParams.OrderBy = mapOrderImageBy(imageListQueryParams.OrderBy)

		ctx.Set("image_list_query_params", imageListQueryParams)
		ctx.Next()
	}
}

func validateOrderImageBy(orderBy string) error {

	switch orderBy {
	case "ID", "Date", "Name":
		return nil
	}

	return errors.New("Invalid `order_by` type")
}

func mapOrderImageBy(orderBy string) string {

	switch orderBy {
	case "ID":
		return "id"
	case "Name":
		return "original_name"
	case "Date":
		return "created_at"
	}

	return ""
}

func validateOrderDir(orderDir string) error {

	switch orderDir {
	case "ASC", "DESC":
		return nil
	}

	return errors.New("Invalid `order_dir` type")
}
