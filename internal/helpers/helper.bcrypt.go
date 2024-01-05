package helpers

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(text string) (string, error) {
	if text == "" {
		return "", errors.New("text can't be empty")
	}

	result, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func ValidateHash(text, hashedText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedText), []byte(text))
	return err == nil
}
