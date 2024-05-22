package contract

import (
	"github.com/gin-gonic/gin"
	"github.com/zakirkun/test-elabram/app/domain/model"
	"github.com/zakirkun/test-elabram/app/domain/types"
)

type EmployeRepository interface {
	CreateEmploye(data model.Employe) error
	GetAllEmploye() (*[]model.Employe, error)
	GetEmploye(id int) (*model.Employe, error)
	UpdateEmploye(id int, data model.Employe) error
	DeleteEmploye(id int) error
	FindByUsername(username string) (*model.Employe, error)
}

type EmployeUseCase interface {
	CreateEmploye(request types.RequestCreatEmploye) (*types.ResponseCreateEmploye, error)
	GetAllEmploye() (*[]types.ResponseGetEmploye, error)
	GetEmploye(request types.RequestGetEmploye) (*types.ResponseGetEmploye, error)
	UpdateEmploye(request types.RequestUpdateEmploye) (*types.ResponseUpdateEmploye, error)
	DeleteEmploye(request types.RequestGetEmploye) (*types.ResponseDeleteEmploye, error)
	LoginEmploye(request types.RequestLoginEmploye) (*types.ResponseLoginEmploye, error)
	RegisterEmploye(request types.RequestCreatEmploye) (*types.ResponseCreateEmploye, error)
}

type EmployeController interface {
	CreateEmploye(ctx *gin.Context)
	GetAllEmploye(ctx *gin.Context)
	GetEmploye(ctx *gin.Context)
	UpdateEmploye(ctx *gin.Context)
	DeleteEmploye(ctx *gin.Context)
	LoginEmploy(ctx *gin.Context)
	RegisterEmploye(ctx *gin.Context)
}
