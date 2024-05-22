package usecase

import (
	"errors"

	"github.com/zakirkun/test-elabram/app/domain/contract"
	"github.com/zakirkun/test-elabram/app/domain/model"
	"github.com/zakirkun/test-elabram/app/domain/types"
	"github.com/zakirkun/test-elabram/app/repository"
)

type claimUseCaseCtx struct{}

func NewClaimUseCase() contract.ClaimUseCase {
	return claimUseCaseCtx{}
}

func (c claimUseCaseCtx) DetailClaim(request types.RequestDetailClaim) (*types.ResponseDetailClaim, error) {

	repo := repository.NewClaimRepository()

	getClaim, err := repo.GetClaimOne(map[string]interface{}{
		"id": request.ID,
	})

	if err != nil {
		return nil, err
	}

	if getClaim.ID == 0 {
		return nil, errors.New("claim not found")
	}

	return &types.ResponseDetailClaim{
		Data: getClaim,
	}, nil
}

func (c claimUseCaseCtx) DeleteClaim(request types.RequestDeleteClaim) (*types.ResponseDeleteClaim, error) {

	repo := repository.NewClaimRepository()

	getClaim, err := repo.GetClaimOne(map[string]interface{}{
		"id":         request.ID,
		"id_employe": request.IDEmploye,
	})

	if err != nil {
		return nil, err
	}

	if getClaim.ID == 0 {
		return nil, errors.New("claim not found")
	}

	if err := repo.DeleteClaim(request.ID); err != nil {
		return nil, err
	}

	return &types.ResponseDeleteClaim{
		Message: "Claim deleted",
	}, nil
}

func (c claimUseCaseCtx) UpdateClaim(request types.RequestUpdateClaim) (*types.ResponseUpdateClaim, error) {
	repo := repository.NewClaimRepository()

	getClaim, err := repo.GetClaimOne(map[string]interface{}{
		"id":         request.IDClaim,
		"id_employe": request.IDEmploye,
	})

	if err != nil {
		return nil, err
	}

	if getClaim.ID == 0 {
		return nil, errors.New("claim not found")
	}

	if err := repo.UpdateClaim(int(getClaim.ID), model.RequestClaim{
		Category:    request.Category,
		Currency:    request.Currency,
		Date:        request.Date,
		Amount:      request.Amount,
		Description: request.Description,
		Document:    request.URLDocument,
	}); err != nil {
		return nil, err
	}

	return &types.ResponseUpdateClaim{
		Message: "Claim updated",
	}, nil
}

func (c claimUseCaseCtx) CreateClaim(request types.RequestCreateClaim) (*types.ResponseCreateClaim, error) {

	repoEmploye := repository.NewEmployeRepository()
	repoClaim := repository.NewClaimRepository()

	getEmplploye, err := repoEmploye.GetEmploye(request.IDEmploye)
	if err != nil {
		return nil, err
	}

	if err := repoClaim.CreateClaim(model.RequestClaim{
		IDEmploye:   int(getEmplploye.ID),
		IDCompany:   getEmplploye.IDCOmpany,
		Category:    request.Category,
		Currency:    request.Currency,
		Date:        request.Date,
		Amount:      request.Amount,
		Description: request.Description,
		Document:    request.URLDocument,
		Status:      "PENDING",
	}); err != nil {
		return nil, err
	}

	return &types.ResponseCreateClaim{
		Message: "Claim was sent",
	}, nil
}

func (c claimUseCaseCtx) GetByEmploye(request types.RequestGetByEmploye) (*types.ResponseGetByEmploye, error) {

	repo := repository.NewClaimRepository()
	getAllClaim, err := repo.GetClaimWithWhere(map[string]interface{}{
		"id_employe": request.ID,
	})

	if err != nil {
		return nil, err
	}

	return &types.ResponseGetByEmploye{
		Data: getAllClaim,
	}, nil
}

func (c claimUseCaseCtx) ApproveClaim(request types.RequestApproveClaim) (*types.ResponseApproveClaim, error) {

	repo := repository.NewClaimRepository()

	if err := repo.UpdateClaim(request.ID, model.RequestClaim{
		Status: "APPROVE",
	}); err != nil {
		return nil, err
	}

	return &types.ResponseApproveClaim{
		Message: "Claim Approved",
	}, nil
}

func (c claimUseCaseCtx) RejectClaim(request types.RequestRejectClaim) (*types.ResponseRejectClaim, error) {

	repo := repository.NewClaimRepository()

	if err := repo.UpdateClaim(request.ID, model.RequestClaim{
		Status: "REJECT",
	}); err != nil {
		return nil, err
	}

	return &types.ResponseRejectClaim{
		Message: "Claim Rejected",
	}, nil
}

func (c claimUseCaseCtx) GetByCompany(request types.RequestGetByCompany) (*types.ResponseGetByCompany, error) {

	repo := repository.NewClaimRepository()
	getAllClaim, err := repo.GetClaimWithWhere(map[string]interface{}{
		"id_company": request.ID,
	})

	if err != nil {
		return nil, err
	}

	return &types.ResponseGetByCompany{
		Data: getAllClaim,
	}, nil
}

func (c claimUseCaseCtx) GetAllClaim() (*types.ResponseGetAllClaim, error) {

	repo := repository.NewClaimRepository()
	getAllClaim, err := repo.GetAllClaim()

	if err != nil {
		return nil, err
	}

	return &types.ResponseGetAllClaim{
		Data: getAllClaim,
	}, nil
}
