package utils

import (
	"time"
)

// ProfileInfo structure representing neccessary
// profile data related to user
type ProfileInfo struct {
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

// PaginationResult structure representing data
// thet will be returned after pagination.
// Should always be extended and another field
// named `Results` should be added to the new structure
type PaginationResult struct {
	TotalResults uint `json:"total_results"`
	From         uint `json:"from"`
	Size         uint `json:"size"`
}

// PaginatedImageResults stucture representing
// data for paginated list of images
type PaginatedImageResults struct {
	PaginationResult
	Results []ImageInfo `json:"results"`
}

// ImageInfo structure representing client related
// data for an image
type ImageInfo struct {
	ID         uint     `json:"id"`
	URL        string   `json:"url"`
	Name       string   `json:"name"`
	UploadedBy string   `json:"uploaded_by"`
	UploadedAt string   `json:"uploaded_at"`
	Tags       []string `json:"tags"`
}
