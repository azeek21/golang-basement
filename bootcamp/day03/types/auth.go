package types

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

const AUTH_TOKEN_KEY = "Bearer"
