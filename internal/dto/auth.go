// Package dto defines the data transfer objects (DTOs) for authentication-related operations in the CartQL application, including user registration, login, token refresh, and profile updates.
package dto

// RegisterRequest represents the request payload for user registration, containing the email, password, first name, last name, and optional phone number of the user.
type RegisterRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Phone     string `json:"phone"`
}

// LoginRequest represents the request payload for user login, containing the email and password of the user attempting to authenticate.
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// RefreshTokenRequest represents the request payload for refreshing an access token, containing the refresh token that was issued during the initial authentication process.
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// AuthResponse represents the response structure for authentication-related API endpoints, including the authenticated user's details, access token, and refresh token for maintaining the session.
type AuthResponse struct {
	User         UserResponse `json:"user"`
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
}

// UserResponse represents the response structure for user-related API endpoints, including the user's ID, email, first name, last name, phone number, role, and active status.
type UserResponse struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	IsActive  bool   `json:"is_active"`
}

// UpdateProfileRequest represents the request payload for updating a user's profile information, allowing changes to the first name, last name, and phone number of the user.
type UpdateProfileRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Phone     string `json:"phone"`
}
