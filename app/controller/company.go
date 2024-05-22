package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zakirkun/test-elabram/app/domain/contract"
	"github.com/zakirkun/test-elabram/app/domain/types"
	"github.com/zakirkun/test-elabram/app/usecase"
	"github.com/zakirkun/test-elabram/internal/utils"
)

type companyControllerCtx struct{}

func NewCompanyController() contract.CompanyController {
	return companyControllerCtx{}
}

func (c companyControllerCtx) CreateCompany(ctx *gin.Context) {

	var request types.RequestCreateCompany
	id := uuid.NewString()
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	uc := usecase.NewCompanyUseCase()
	response, err := uc.CreateCompany(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	ctx.JSON(http.StatusCreated, utils.SetGeneralResponse("SUCCESS", response.Message, id, nil))
}

func (c companyControllerCtx) GetAllCompany(ctx *gin.Context) {

	id := uuid.NewString()
	uc := usecase.NewCompanyUseCase()
	response, err := uc.GetAllCompany()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("SUCCESS", "List Company", id, response))
}
