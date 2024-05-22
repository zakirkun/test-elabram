package controller

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zakirkun/test-elabram/app/domain/contract"
	"github.com/zakirkun/test-elabram/app/domain/types"
	"github.com/zakirkun/test-elabram/app/usecase"
	"github.com/zakirkun/test-elabram/internal/utils"
)

type claimControllerCtx struct{}

func NewClaimController() contract.ClaimController {
	return claimControllerCtx{}
}

func (c claimControllerCtx) CreateClaim(ctx *gin.Context) {

	var request types.RequestCreateClaim
	id := uuid.NewString()

	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	_employeID, _ := strconv.Atoi(ctx.GetString("_id"))
	request.IDEmploye = _employeID

	file, _ := ctx.FormFile("document")
	if file != nil {
		uploadPath := filepath.Join("public", file.Filename)
		if err := ctx.SaveUploadedFile(file, uploadPath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		request.URLDocument = uploadPath
	}

	uc := usecase.NewClaimUseCase()
	response, err := uc.CreateClaim(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	ctx.JSON(http.StatusCreated, utils.SetGeneralResponse("SUCCESS", response.Message, id, nil))
}

func (c claimControllerCtx) ApproveClaim(ctx *gin.Context) {

	var request types.RequestApproveClaim
	id := uuid.NewString()

	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	uc := usecase.NewClaimUseCase()
	response, err := uc.ApproveClaim(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("SUCCESS", response.Message, id, nil))
}

func (c claimControllerCtx) RejectClaim(ctx *gin.Context) {

	var request types.RequestRejectClaim
	id := uuid.NewString()

	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	uc := usecase.NewClaimUseCase()
	response, err := uc.RejectClaim(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("SUCCESS", response.Message, id, nil))
}

func (c claimControllerCtx) GetByEmploye(ctx *gin.Context) {

	var request types.RequestGetByEmploye
	id := uuid.NewString()

	_employeID, _ := strconv.Atoi(ctx.GetString("_id"))
	request.ID = _employeID

	uc := usecase.NewClaimUseCase()
	response, err := uc.GetByEmploye(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("SUCCESS", "Detail Claim", id, response))
}

func (c claimControllerCtx) GetByCompany(ctx *gin.Context) {

	var request types.RequestGetByCompany
	id := uuid.NewString()

	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	uc := usecase.NewClaimUseCase()
	response, err := uc.GetByCompany(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("SUCCESS", "List Claim", id, response))
}

func (c claimControllerCtx) GetAllClaim(ctx *gin.Context) {

	id := uuid.NewString()
	uc := usecase.NewClaimUseCase()
	response, err := uc.GetAllClaim()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("SUCCESS", "List Claim", id, response))
}

func (c claimControllerCtx) UpdateClaim(ctx *gin.Context) {

	var request types.RequestUpdateClaim
	id := uuid.NewString()

	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	_employeID, _ := strconv.Atoi(ctx.GetString("_id"))
	request.IDEmploye = _employeID

	file, _ := ctx.FormFile("document")
	if file != nil {
		filename := filepath.Base(file.Filename)
		filepath := filepath.Join("./public", filename)
		if err := ctx.SaveUploadedFile(file, filepath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		request.URLDocument = filepath
	}

	uc := usecase.NewClaimUseCase()
	response, err := uc.UpdateClaim(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("SUCCESS", response.Message, id, nil))
}

func (c claimControllerCtx) DeleteClaim(ctx *gin.Context) {

	var request types.RequestDeleteClaim
	id := uuid.NewString()

	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	_employeID, _ := strconv.Atoi(ctx.GetString("_id"))
	request.IDEmploye = _employeID

	uc := usecase.NewClaimUseCase()
	response, err := uc.DeleteClaim(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("SUCCESS", response.Message, id, nil))
}

func (c claimControllerCtx) DetailClaim(ctx *gin.Context) {

	var request types.RequestDetailClaim
	id := uuid.NewString()

	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	uc := usecase.NewClaimUseCase()
	response, err := uc.DetailClaim(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("ERROR", err.Error(), id))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("SUCCESS", "Detail Claim", id, response))
}
