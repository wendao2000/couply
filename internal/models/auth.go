package model

type Auth struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	*Auth
	Name        string `json:"name"`
	DateOfBirth string `json:"dob"` // Format: "2006-01-02"
}

type UserToken struct {
	Token        string `json:"token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
