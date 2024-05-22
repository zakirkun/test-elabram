package model

import "gorm.io/gorm"

type Company struct {
	*gorm.Model
	Name        string
	Description string
	Logo        string
}
