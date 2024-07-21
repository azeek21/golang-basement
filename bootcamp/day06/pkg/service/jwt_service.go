package service

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/azeek21/blog/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

var ERR_BAD_TOKEN = errors.New("Bad token")
var ERR_CANT_PARS_CLAIMS = errors.New("Failed to get claims from token")
var ERR_CANT_PARSE_USER_ID = errors.New("Failed to parse user id from token")
var ERR_CANT_CREATE_TOKEN = errors.New("Failed to create new authentication token")

type JwtServ struct{}

func NewJwtSerice() JwtService {
	return JwtServ{}
}

func (s JwtServ) CreateJwt(user *models.User) (string, error) {
	token := ""
	key := []byte(viper.GetString("PRIVATE_KEY"))

	generatedAt := time.Now().Unix()
	expiresAt := generatedAt + int64((time.Hour * 720).Seconds())
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   fmt.Sprint(user.ID),
		"iss":   "blog.askaraliev.uz",
		"iat":   generatedAt,
		"exp":   expiresAt,
		"email": user.Email,
	})
	token, err := t.SignedString(key)
	if err != nil {
		return token, ERR_CANT_CREATE_TOKEN

	}
	return token, nil

}

func (s JwtServ) VerifyJwt(token string) (uint, error) {
	user_id := uint(0)
	jwtParsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ERR_BAD_TOKEN
		}

		secret := []byte(viper.GetString("PRIVATE_KEY"))
		return secret, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		return user_id, err
	}

	claims, ok := jwtParsed.Claims.(jwt.MapClaims)
	if !ok {
		return user_id, ERR_CANT_PARS_CLAIMS
	}

	subject, err := claims.GetSubject()

	if err != nil {
		return user_id, err
	}

	user_id_int, err := strconv.Atoi(subject)
	if err != nil {
		return user_id, ERR_CANT_PARSE_USER_ID
	}

	user_id = uint(user_id_int)
	return user_id, nil
}
