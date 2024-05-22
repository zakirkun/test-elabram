package repository

import (
	"github.com/zakirkun/test-elabram/app/domain/contract"
	"github.com/zakirkun/test-elabram/app/domain/model"
	"github.com/zakirkun/test-elabram/internal/globals"
)

type adminRepositoryCtx struct{}

func NewAdminRepository() contract.AdminRepository {
	return adminRepositoryCtx{}
}

func (a adminRepositoryCtx) CreateAdmin(data model.Admin) error {
	db, err := globals.DB.OpenDB()
	if err != nil {
		return *err
	}

	if err := db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (a adminRepositoryCtx) GetAllAdmin() (*[]model.Admin, error) {
	db, err := globals.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data []model.Admin
	if err := db.Model(&model.Admin{}).Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (a adminRepositoryCtx) GetAdmin(id int) (*model.Admin, error) {
	db, err := globals.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data model.Admin
	if err := db.Model(&model.Admin{}).Where("id = ?", id).Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (a adminRepositoryCtx) UpdateAdmin(id int, data model.Admin) error {

	db, err := globals.DB.OpenDB()
	if err != nil {
		return *err
	}

	if err := db.Model(&model.Admin{}).Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func (a adminRepositoryCtx) DeleteAdmin(id int) error {

	db, err := globals.DB.OpenDB()
	if err != nil {
		return *err
	}

	if err := db.Delete(&model.Admin{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (a adminRepositoryCtx) FindByUsername(username string) (*model.Admin, error) {

	db, err := globals.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data model.Admin
	if err := db.Model(&model.Admin{}).Where("username = ?", username).Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
