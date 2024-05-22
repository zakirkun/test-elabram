package model

import (
	"time"

	"gorm.io/gorm"
)

type RequestClaim struct {
	*gorm.Model

	IDEmploye   int
	IDCompany   int
	Category    string
	Date        time.Time
	Currency    string
	Amount      int64
	Description string
	Document    string
	Status      string
}
