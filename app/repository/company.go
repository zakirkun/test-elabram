package repository

import (
	"github.com/zakirkun/test-elabram/app/domain/contract"
	"github.com/zakirkun/test-elabram/app/domain/model"
	"github.com/zakirkun/test-elabram/internal/globals"
)

type companyRepositoryCtx struct{}

func NewCompanyRepository() contract.CompanyRepository {
	return companyRepositoryCtx{}
}

func (c companyRepositoryCtx) CreateCompany(data model.Company) error {

	db, err := globals.DB.OpenDB()
	if err != nil {
		return *err
	}

	if err := db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (c companyRepositoryCtx) GetAllCompany() (*[]model.Company, error) {

	db, err := globals.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data []model.Company
	if err := db.Model(&model.Company{}).Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
