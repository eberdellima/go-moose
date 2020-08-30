package controllers

import (
	"go-moose/src/feed/inputs"
	"go-moose/src/feed/services"
	UserInput "go-moose/src/user/inputs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListImages returns a paginated list of images
func ListImages(ctx *gin.Context) {

	paginator, _ := ctx.Get("paginator")
	assertedPaginator := paginator.(UserInput.Paginator)

	imageListQueryingParams, _ := ctx.Get("image_list_query_params")
	assertedQueryingParams := imageListQueryingParams.(inputs.ImageListQueryingParams)

	imageList := services.GetImageList(&assertedQueryingParams, &assertedPaginator)

	ctx.JSON(http.StatusOK, imageList)
}
