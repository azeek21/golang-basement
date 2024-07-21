package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type PasswordServ struct {
}

func NewPasswordSerice() PasswordService {
	return PasswordServ{}
}

var ERR_FAILED_HASH_PASSWD = errors.New("There's something wrong with your password. We failed to encrypt it. Plase try changing it.")

func (s PasswordServ) CreateHash(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	if err != nil {
		return "", ERR_FAILED_HASH_PASSWD
	}
	return string(hash), nil
}

func (s PasswordServ) VerifyPasswordAgainstHash(pwd string, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	if err != nil {
		return false, err
	}
	return true, err
}
