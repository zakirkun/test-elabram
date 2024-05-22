package model

import "gorm.io/gorm"

type Admin struct {
	*gorm.Model

	Email    string
	Username string
	Password string
	FullName string
}
