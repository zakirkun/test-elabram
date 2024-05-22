package types

import (
	"mime/multipart"
	"time"

	"github.com/zakirkun/test-elabram/app/domain/model"
)

type RequestDetailClaim struct {
	ID int `uri:"id" binding:"required"`
}

type ResponseDetailClaim struct {
	Data *model.RequestClaim `json:"data"`
}

type RequestDeleteClaim struct {
	ID        int `uri:"id" binding:"required"`
	IDEmploye int
}

type ResponseDeleteClaim struct {
	Message string `json:"message"`
}

type RequestUpdateClaim struct {
	Category    string                `form:"category"`
	Date        time.Time             `form:"date"`
	Currency    string                `form:"currency"`
	Amount      int64                 `form:"amount"`
	Description string                `form:"description"`
	Document    *multipart.FileHeader `form:"document"`
	IDClaim     int                   `form:"id_claim"`
	IDEmploye   int
	URLDocument string
}

type ResponseUpdateClaim struct {
	Message string `json:"message"`
}

type RequestCreateClaim struct {
	Category    string                `form:"category"`
	Date        time.Time             `form:"date"`
	Currency    string                `form:"currency"`
	Amount      int64                 `form:"amount"`
	Description string                `form:"description"`
	Document    *multipart.FileHeader `form:"document"`
	IDEmploye   int
	URLDocument string
}

type ResponseCreateClaim struct {
	Message string `json:"message"`
}

type RequestApproveClaim struct {
	ID int `json:"id"`
}

type ResponseApproveClaim struct {
	Message string `json:"message"`
}

type RequestRejectClaim struct {
	ID int `json:"id"`
}

type ResponseRejectClaim struct {
	Message string `json:"message"`
}

type RequestGetByCompany struct {
	ID int `uri:"id" binding:"required"`
}

type ResponseGetByCompany struct {
	Data *[]model.RequestClaim `json:"data"`
}

type RequestGetByEmploye struct {
	ID int
}

type ResponseGetByEmploye struct {
	Data *[]model.RequestClaim `json:"data"`
}

type ResponseGetAllClaim struct {
	Data *[]model.RequestClaim `json:"data"`
}
