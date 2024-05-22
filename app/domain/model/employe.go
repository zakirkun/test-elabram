package model

import "gorm.io/gorm"

type Employe struct {
	*gorm.Model

	IDCOmpany int
	Email     string
	Msisdn    string
	Username  string
	Password  string
	FullName  string
	Position  string
	Foto      string
}
