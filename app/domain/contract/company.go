package contract

import (
	"github.com/gin-gonic/gin"
	"github.com/zakirkun/test-elabram/app/domain/model"
	"github.com/zakirkun/test-elabram/app/domain/types"
)

type CompanyRepository interface {
	CreateCompany(data model.Company) error
	GetAllCompany() (*[]model.Company, error)
}

type CompanyUseCase interface {
	CreateCompany(request types.RequestCreateCompany) (*types.ResponseCreateCompany, error)
	GetAllCompany() (*[]types.ResponseGetCompany, error)
}

type CompanyController interface {
	CreateCompany(ctx *gin.Context)
	GetAllCompany(ctx *gin.Context)
}
