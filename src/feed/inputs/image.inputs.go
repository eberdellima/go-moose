package inputs

// ImageTags structure representing tags
// to be added to an image
type ImageTags struct {
	Tags []string `json:"tags" binding:"required"`
}
