package utils

// TokenValidationError represents error
// when token validation fails for some reason
type TokenValidationError struct{}

func (e *TokenValidationError) Error() string {
	return "Not authorized"
}
