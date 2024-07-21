package models

const AUTH_MODEL_NAME = "authed"
const AUTH_COOKIE_NAME = "Authentication"
const PASSWORD_REGEX = "^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{8,}$"

type SignInDTO struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type SignUpDTO struct {
	Name           string `form:"name"`
	Username       string `form:"username" binding:"required"`
	Password       string `form:"password" binding:"required"`
	Email          string `form:"email" binding:"required"`
	PasswordRepeat string `form:"repeat-password" binding:"required"`
}
