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

type adminControllerCtx struct{}

func NewAdminController() contract.AdminController {
	return adminControllerCtx{}
}

func (a adminControllerCtx) CreateAdmin(ctx *gin.Context) {

	var request types.RequestCreatAdmin
	id := uuid.NewString()
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	uc := usecase.NewAdminUseCase()
	response, err := uc.CreateAdmin(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	ctx.JSON(http.StatusCreated, utils.SetGeneralResponse("SUCCESS", response.Message, id, nil))
}

func (a adminControllerCtx) GetAllAdmin(ctx *gin.Context) {

	id := uuid.NewString()
	uc := usecase.NewAdminUseCase()
	response, err := uc.GetAllAdmin()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("SUCCESS", "List admin", id, response))
}

func (a adminControllerCtx) GetAdmin(ctx *gin.Context) {

	var request types.RequestGetAdmin
	id := uuid.NewString()
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	uc := usecase.NewAdminUseCase()
	response, err := uc.GetAdmin(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("SUCCESS", "Detail admin", id, response))
}

func (a adminControllerCtx) UpdateAdmin(ctx *gin.Context) {

	var request types.RequestUpdateAdmin
	id := uuid.NewString()
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	uc := usecase.NewAdminUseCase()
	response, err := uc.UpdateAdmin(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("SUCCESS", response.Message, id, nil))
}

func (a adminControllerCtx) DeleteAdmin(ctx *gin.Context) {

	var request types.RequestGetAdmin
	id := uuid.NewString()
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	uc := usecase.NewAdminUseCase()
	response, err := uc.DeleteAdmin(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("SUCCESS", response.Message, id, nil))
}

func (a adminControllerCtx) LoginAdmin(ctx *gin.Context) {

	var request types.RequestLoginAdmin
	id := uuid.NewString()
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	uc := usecase.NewAdminUseCase()
	response, err := uc.LoginAdmin(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("SUCCESS", "Login success", id, response))
}

func (a adminControllerCtx) RegisterAdmin(ctx *gin.Context) {

	var request types.RequestCreatAdmin
	id := uuid.NewString()
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	uc := usecase.NewAdminUseCase()
	response, err := uc.CreateAdmin(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	ctx.JSON(http.StatusCreated, utils.SetGeneralResponse("SUCCESS", response.Message, id, nil))
}
