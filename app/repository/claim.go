package repository

import (
	"github.com/zakirkun/test-elabram/app/domain/contract"
	"github.com/zakirkun/test-elabram/app/domain/model"
	"github.com/zakirkun/test-elabram/internal/globals"
)

type claimRepositoryCtx struct{}

func NewClaimRepository() contract.ClaimRepository {
	return claimRepositoryCtx{}
}

func (c claimRepositoryCtx) CreateClaim(data model.RequestClaim) error {

	db, err := globals.DB.OpenDB()
	if err != nil {
		return *err
	}

	if err := db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (c claimRepositoryCtx) GetAllClaim() (*[]model.RequestClaim, error) {

	db, err := globals.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data []model.RequestClaim
	if err := db.Model(&model.RequestClaim{}).Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil

}

func (c claimRepositoryCtx) GetClaimWithWhere(cond map[string]interface{}) (*[]model.RequestClaim, error) {

	db, err := globals.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data []model.RequestClaim

	if err := db.Model(&model.RequestClaim{}).Where(cond).Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (c claimRepositoryCtx) UpdateClaim(id int, data model.RequestClaim) error {

	db, err := globals.DB.OpenDB()
	if err != nil {
		return *err
	}

	if err := db.Model(&model.RequestClaim{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func (c claimRepositoryCtx) DeleteClaim(id int) error {

	db, err := globals.DB.OpenDB()
	if err != nil {
		return *err
	}

	if err := db.Delete(&model.RequestClaim{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (c claimRepositoryCtx) GetClaimOne(cond map[string]interface{}) (*model.RequestClaim, error) {

	db, err := globals.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data model.RequestClaim

	if err := db.Model(&model.RequestClaim{}).Where(cond).Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
