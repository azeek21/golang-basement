package service

import (
	"errors"
	"fmt"
	"server/pkg/repository"
	"server/types"
	"server/types/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	repo repository.User
}

func NewAuthService(repo repository.User) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (authservice *AuthService) LogIn(credentials types.LoginDTO) (string, error) {
	user, err := authservice.repo.GetByEmail(credentials.Email)

	if err != nil {
		return "", errors.New(fmt.Sprintf("auth: bad credentials: %s", err.Error()))
	}

	var secret []byte
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(25 * time.Hour).Unix(),
	})
	secret = []byte(types.GLOABAL_CONFIG.Secret)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (authservice *AuthService) LogOut(user models.User) (bool, error) {
	return true, nil
}

func (authservice *AuthService) Register(user models.User) (string, error) {

	res, err := authservice.repo.Create(user)
	if err != nil {
		return "", err
	}
	return res, nil
}
