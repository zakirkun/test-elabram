package repository

import (
	"github.com/zakirkun/test-elabram/app/domain/contract"
	"github.com/zakirkun/test-elabram/app/domain/model"
	"github.com/zakirkun/test-elabram/internal/globals"
)

type employeRepositoryCtx struct{}

func NewEmployeRepository() contract.EmployeRepository {
	return employeRepositoryCtx{}
}

func (e employeRepositoryCtx) CreateEmploye(data model.Employe) error {

	db, err := globals.DB.OpenDB()
	if err != nil {
		return *err
	}

	if err := db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (e employeRepositoryCtx) GetAllEmploye() (*[]model.Employe, error) {

	db, err := globals.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data []model.Employe
	if err := db.Model(&model.Employe{}).Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (e employeRepositoryCtx) GetEmploye(id int) (*model.Employe, error) {

	db, err := globals.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data model.Employe
	if err := db.Model(&model.Employe{}).Where("id = ?", id).Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (e employeRepositoryCtx) UpdateEmploye(id int, data model.Employe) error {

	db, err := globals.DB.OpenDB()
	if err != nil {
		return *err
	}

	if err := db.Model(&model.Employe{}).Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func (e employeRepositoryCtx) DeleteEmploye(id int) error {
	db, err := globals.DB.OpenDB()
	if err != nil {
		return *err
	}

	if err := db.Delete(&model.Employe{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (e employeRepositoryCtx) FindByUsername(username string) (*model.Employe, error) {

	db, err := globals.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data model.Employe
	if err := db.Model(&model.Employe{}).Where("username = ?", username).Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil

}
