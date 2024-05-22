package usecase

import (
	"github.com/zakirkun/test-elabram/app/domain/contract"
	"github.com/zakirkun/test-elabram/app/domain/model"
	"github.com/zakirkun/test-elabram/app/domain/types"
	"github.com/zakirkun/test-elabram/app/repository"
)

type companyUseCaseCtx struct{}

func NewCompanyUseCase() contract.CompanyUseCase {
	return companyUseCaseCtx{}
}

func (c companyUseCaseCtx) CreateCompany(request types.RequestCreateCompany) (*types.ResponseCreateCompany, error) {

	repo := repository.NewCompanyRepository()
	if err := repo.CreateCompany(model.Company{
		Name:        request.Name,
		Description: request.Description,
		Logo:        request.Logo,
	}); err != nil {
		return nil, err
	}

	return &types.ResponseCreateCompany{
		Message: "Company Created",
	}, nil
}

func (c companyUseCaseCtx) GetAllCompany() (*[]types.ResponseGetCompany, error) {

	repo := repository.NewCompanyRepository()
	getCompany, err := repo.GetAllCompany()
	if err != nil {
		return nil, err
	}

	result := make([]types.ResponseGetCompany, 0)

	if len(*getCompany) > 0 {
		for _, company := range *getCompany {
			result = append(result, types.ResponseGetCompany{
				Name:        company.Name,
				Description: company.Description,
				Logo:        company.Logo,
			})
		}
	}

	return &result, nil
}
