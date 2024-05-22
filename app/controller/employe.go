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

type employeControllerCtx struct{}

func NewEmployeController() contract.EmployeController {
	return employeControllerCtx{}
}

func (e employeControllerCtx) CreateEmploye(ctx *gin.Context) {
	var request types.RequestCreatEmploye
	id := uuid.NewString()
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	uc := usecase.NewEmployeUseCase()
	response, err := uc.CreateEmploye(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	ctx.JSON(http.StatusCreated, utils.SetGeneralResponse("SUCCESS", response.Message, id, nil))
}

func (e employeControllerCtx) GetAllEmploye(ctx *gin.Context) {

	id := uuid.NewString()
	uc := usecase.NewEmployeUseCase()
	response, err := uc.GetAllEmploye()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("SUCCESS", "List employe", id, response))
}

func (e employeControllerCtx) GetEmploye(ctx *gin.Context) {

	var request types.RequestGetEmploye
	id := uuid.NewString()
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	uc := usecase.NewEmployeUseCase()
	response, err := uc.GetEmploye(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("SUCCESS", "Detail employe", id, response))
}

func (e employeControllerCtx) UpdateEmploye(ctx *gin.Context) {

	var request types.RequestUpdateEmploye
	id := uuid.NewString()
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	uc := usecase.NewEmployeUseCase()
	response, err := uc.UpdateEmploye(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("SUCCESS", response.Message, id, nil))
}

func (e employeControllerCtx) DeleteEmploye(ctx *gin.Context) {

	var request types.RequestGetEmploye
	id := uuid.NewString()
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	uc := usecase.NewEmployeUseCase()
	response, err := uc.DeleteEmploye(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("SUCCESS", response.Message, id, nil))
}

func (e employeControllerCtx) LoginEmploy(ctx *gin.Context) {

	var request types.RequestLoginEmploye
	id := uuid.NewString()
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	uc := usecase.NewEmployeUseCase()
	response, err := uc.LoginEmploye(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("SUCCESS", "Login success", id, response))
}

func (e employeControllerCtx) RegisterEmploye(ctx *gin.Context) {

	var request types.RequestCreatEmploye
	id := uuid.NewString()
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	uc := usecase.NewEmployeUseCase()
	response, err := uc.CreateEmploye(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	ctx.JSON(http.StatusCreated, utils.SetGeneralResponse("SUCCESS", response.Message, id, nil))
}
