package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(Password string) (string, error) {

	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(Password), 14)

	if err != nil {
		return "", err
	}

	return string(hashedpassword), err
}

func CheckPasswordHash(password, hashedpassword string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
	return err == nil
}
