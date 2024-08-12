package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hash), err
}

func CompairePAssword(passsword, HashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(HashPassword), []byte(passsword))
	return err == nil
}
