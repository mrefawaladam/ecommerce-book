package service

import "golang.org/x/crypto/bcrypt"

func Encrypt(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
