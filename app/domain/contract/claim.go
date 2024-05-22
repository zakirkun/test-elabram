package contract

import (
	"github.com/gin-gonic/gin"
	"github.com/zakirkun/test-elabram/app/domain/model"
	"github.com/zakirkun/test-elabram/app/domain/types"
)

type ClaimRepository interface {
	CreateClaim(data model.RequestClaim) error
	GetAllClaim() (*[]model.RequestClaim, error)
	GetClaimWithWhere(cond map[string]interface{}) (*[]model.RequestClaim, error)
	UpdateClaim(id int, data model.RequestClaim) error
	DeleteClaim(id int) error
	GetClaimOne(cond map[string]interface{}) (*model.RequestClaim, error)
}

type ClaimUseCase interface {
	CreateClaim(request types.RequestCreateClaim) (*types.ResponseCreateClaim, error)
	GetByEmploye(request types.RequestGetByEmploye) (*types.ResponseGetByEmploye, error)
	ApproveClaim(request types.RequestApproveClaim) (*types.ResponseApproveClaim, error)
	RejectClaim(request types.RequestRejectClaim) (*types.ResponseRejectClaim, error)
	GetByCompany(request types.RequestGetByCompany) (*types.ResponseGetByCompany, error)
	GetAllClaim() (*types.ResponseGetAllClaim, error)
	UpdateClaim(request types.RequestUpdateClaim) (*types.ResponseUpdateClaim, error)
	DeleteClaim(request types.RequestDeleteClaim) (*types.ResponseDeleteClaim, error)
	DetailClaim(request types.RequestDetailClaim) (*types.ResponseDetailClaim, error)
}

type ClaimController interface {
	CreateClaim(ctx *gin.Context)
	ApproveClaim(ctx *gin.Context)
	RejectClaim(ctx *gin.Context)
	GetByCompany(ctx *gin.Context)
	GetByEmploye(ctx *gin.Context)
	GetAllClaim(ctx *gin.Context)
	UpdateClaim(ctx *gin.Context)
	DeleteClaim(ctx *gin.Context)
	DetailClaim(ctx *gin.Context)
}
