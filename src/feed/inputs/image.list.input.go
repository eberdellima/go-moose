package inputs

// ImageListQueryingParams structure representing
// optional parameters for querying results for image list
type ImageListQueryingParams struct {
	SearchKeyword string `form:"search-keyword"`
	OrderBy       string `form:"order-by"`
	OrderDir      string `form:"order-dir"`
}
