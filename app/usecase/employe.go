package usecase

import (
	"errors"
	"log"
	"strconv"

	"github.com/zakirkun/test-elabram/app/domain/contract"
	"github.com/zakirkun/test-elabram/app/domain/model"
	"github.com/zakirkun/test-elabram/app/domain/types"
	"github.com/zakirkun/test-elabram/app/repository"
	"github.com/zakirkun/test-elabram/internal/pkg/config"
	"github.com/zakirkun/test-elabram/internal/pkg/jwt"
	"github.com/zakirkun/test-elabram/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type employeUseCaseCtx struct{}

func NewEmployeUseCase() contract.EmployeUseCase {
	return employeUseCaseCtx{}
}

func (e employeUseCaseCtx) CreateEmploye(request types.RequestCreatEmploye) (*types.ResponseCreateEmploye, error) {

	repo := repository.NewEmployeRepository()
	if err := repo.CreateEmploye(model.Employe{
		Email:     request.Email,
		Msisdn:    request.Msisdn,
		IDCOmpany: request.IDCOmpany,
		Username:  request.Username,
		Password:  utils.EncryptPassword(request.Password),
		FullName:  request.FullName,
		Position:  request.Position,
		Foto:      request.Foto,
	}); err != nil {
		return nil, err
	}

	return &types.ResponseCreateEmploye{
		Message: "Employe created",
	}, nil
}

func (e employeUseCaseCtx) GetAllEmploye() (*[]types.ResponseGetEmploye, error) {

	repo := repository.NewEmployeRepository()
	getEmploye, err := repo.GetAllEmploye()
	if err != nil {
		return nil, err
	}

	result := make([]types.ResponseGetEmploye, 0)
	if len(*getEmploye) > 0 {
		for _, employe := range *getEmploye {
			result = append(result, types.ResponseGetEmploye{
				Email:     employe.Email,
				Username:  employe.Username,
				Msisdn:    employe.Msisdn,
				FullName:  employe.FullName,
				IDCOmpany: employe.IDCOmpany,
				Foto:      employe.Foto,
			})
		}
	}

	return &result, nil
}

func (e employeUseCaseCtx) GetEmploye(request types.RequestGetEmploye) (*types.ResponseGetEmploye, error) {

	repo := repository.NewEmployeRepository()
	getEmploye, err := repo.GetEmploye(request.ID)
	if err != nil {
		return nil, err
	}

	if getEmploye.ID == 0 {
		return nil, errors.New("admin not found")
	}

	return &types.ResponseGetEmploye{
		Email:     getEmploye.Email,
		Username:  getEmploye.Username,
		Msisdn:    getEmploye.Msisdn,
		FullName:  getEmploye.FullName,
		IDCOmpany: getEmploye.IDCOmpany,
		Position:  getEmploye.Position,
		Foto:      getEmploye.Foto,
	}, nil
}

func (e employeUseCaseCtx) UpdateEmploye(request types.RequestUpdateEmploye) (*types.ResponseUpdateEmploye, error) {

	repo := repository.NewEmployeRepository()
	getEmploye, err := repo.GetEmploye(request.ID)
	if err != nil {
		return nil, err
	}

	if getEmploye.ID == 0 {
		return nil, errors.New("admin not found")
	}

	if err := repo.UpdateEmploye(request.ID, model.Employe{
		Email:    request.Email,
		FullName: request.FullName,
		Position: request.Position,
		Foto:     request.Foto,
	}); err != nil {
		return nil, err
	}

	return &types.ResponseUpdateEmploye{
		Message: "Employe updated",
	}, nil
}

func (e employeUseCaseCtx) DeleteEmploye(request types.RequestGetEmploye) (*types.ResponseDeleteEmploye, error) {

	repo := repository.NewEmployeRepository()
	getEmploye, err := repo.GetEmploye(request.ID)
	if err != nil {
		return nil, err
	}

	if getEmploye.ID == 0 {
		return nil, errors.New("employe not found")
	}

	if err := repo.DeleteEmploye(request.ID); err != nil {
		return nil, err
	}

	return &types.ResponseDeleteEmploye{
		Message: "Employe deleted",
	}, nil
}

func (e employeUseCaseCtx) LoginEmploye(request types.RequestLoginEmploye) (*types.ResponseLoginEmploye, error) {

	repo := repository.NewEmployeRepository()
	getEmploye, err := repo.FindByUsername(request.Username)
	if err != nil {
		return nil, err
	}

	if getEmploye.ID == 0 {
		return nil, errors.New("employe not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(getEmploye.Password), []byte(request.Password)); err != nil {
		return nil, errors.New("invalid password")
	}

	_toJWT := map[string]interface{}{
		"id":       strconv.Itoa(int(getEmploye.ID)),
		"msisdn":   getEmploye.Msisdn,
		"acc_type": "employe",
		"email":    getEmploye.Email,
	}

	_jwt := jwt.NewJWTImpl(config.GetString("jwt.signature_key"), config.GetInt("jwt.day_expired"))
	token, err := _jwt.GenerateToken(_toJWT)
	if err != nil {
		log.Printf("Error generate token: %v", err.Error())
	}

	return &types.ResponseLoginEmploye{
		Token: token,
	}, nil
}

func (e employeUseCaseCtx) RegisterEmploye(request types.RequestCreatEmploye) (*types.ResponseCreateEmploye, error) {

	repo := repository.NewEmployeRepository()

	if err := repo.CreateEmploye(model.Employe{
		Email:     request.Email,
		Msisdn:    request.Msisdn,
		IDCOmpany: request.IDCOmpany,
		Username:  request.Username,
		Password:  utils.EncryptPassword(request.Password),
		FullName:  request.FullName,
		Position:  request.Position,
		Foto:      request.Foto,
	}); err != nil {
		return nil, err
	}

	return &types.ResponseCreateEmploye{
		Message: "Employe registered",
	}, nil
}
