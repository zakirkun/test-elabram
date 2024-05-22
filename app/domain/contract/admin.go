package contract

import (
	"github.com/gin-gonic/gin"
	"github.com/zakirkun/test-elabram/app/domain/model"
	"github.com/zakirkun/test-elabram/app/domain/types"
)

type AdminRepository interface {
	CreateAdmin(data model.Admin) error
	GetAllAdmin() (*[]model.Admin, error)
	GetAdmin(id int) (*model.Admin, error)
	UpdateAdmin(id int, data model.Admin) error
	DeleteAdmin(id int) error
	FindByUsername(username string) (*model.Admin, error)
}

type AdminUseCase interface {
	CreateAdmin(request types.RequestCreatAdmin) (*types.ResponseCreateAdmin, error)
	GetAllAdmin() (*[]types.ResponseGetAdmin, error)
	GetAdmin(request types.RequestGetAdmin) (*types.ResponseGetAdmin, error)
	UpdateAdmin(request types.RequestUpdateAdmin) (*types.ResponseUpdateAdmin, error)
	DeleteAdmin(request types.RequestGetAdmin) (*types.ResponseDeleteAdmin, error)
	LoginAdmin(request types.RequestLoginAdmin) (*types.ResponseLoginAdmin, error)
	RegisterAdmin(request types.RequestCreatAdmin) (*types.ResponseCreateAdmin, error)
}

type AdminController interface {
	CreateAdmin(ctx *gin.Context)
	GetAllAdmin(ctx *gin.Context)
	GetAdmin(ctx *gin.Context)
	UpdateAdmin(ctx *gin.Context)
	DeleteAdmin(ctx *gin.Context)
	LoginAdmin(ctx *gin.Context)
	RegisterAdmin(ctx *gin.Context)
}
