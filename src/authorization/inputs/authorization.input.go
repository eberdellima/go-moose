package inputs

// LoginInput structure representing how
// payload for login should look like
type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterInput structure representing how
// payload for registration should look like
type RegisterInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}

// RefreshTokenInput structure representing how
// payload for refreshing tokens should look like
type RefreshTokenInput struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
	UserID       uint   `json:"user_id" binding:"required"`
}
