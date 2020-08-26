package inputs

// ProfileInput structure representing data
// that can be updated in user's profile
type ProfileInput struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
