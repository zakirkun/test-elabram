package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) string {

	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed hash password: %v", err)
		return ""
	}

	return string(pass)
}
