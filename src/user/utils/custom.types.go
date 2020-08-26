package utils

import (
	"time"
)

// ProfileInfo structure representing neccessary
// profile data related to user
type ProfileInfo struct {
	Email     string
	Username  string
	ID        uint
	CreatedAt time.Time
}
