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

type adminUseCaseCtx struct{}

func NewAdminUseCase() contract.AdminUseCase {
	return adminUseCaseCtx{}
}

func (a adminUseCaseCtx) CreateAdmin(request types.RequestCreatAdmin) (*types.ResponseCreateAdmin, error) {

	repo := repository.NewAdminRepository()
	if err := repo.CreateAdmin(model.Admin{
		Email:    request.Email,
		Username: request.Username,
		FullName: request.FullName,
		Password: utils.EncryptPassword(request.Password),
	}); err != nil {
		return nil, err
	}

	return &types.ResponseCreateAdmin{
		Message: "Admin was created",
	}, nil
}

func (a adminUseCaseCtx) GetAllAdmin() (*[]types.ResponseGetAdmin, error) {

	repo := repository.NewAdminRepository()
	getAll, err := repo.GetAllAdmin()
	if err != nil {
		return nil, err
	}

	result := make([]types.ResponseGetAdmin, 0)
	if len(*getAll) > 0 {
		for _, admin := range *getAll {
			result = append(result, types.ResponseGetAdmin{
				Email:    admin.Email,
				FullName: admin.FullName,
				Username: admin.Username,
			})
		}
	}

	return &result, nil
}

func (a adminUseCaseCtx) GetAdmin(request types.RequestGetAdmin) (*types.ResponseGetAdmin, error) {

	repo := repository.NewAdminRepository()
	getAdmin, err := repo.GetAdmin(request.ID)
	if err != nil {
		return nil, err
	}

	if getAdmin.ID == 0 {
		return nil, errors.New("admin not found")
	}

	return &types.ResponseGetAdmin{
		Email:    getAdmin.Email,
		Username: getAdmin.Username,
		FullName: getAdmin.FullName,
	}, nil
}

func (a adminUseCaseCtx) UpdateAdmin(request types.RequestUpdateAdmin) (*types.ResponseUpdateAdmin, error) {

	repo := repository.NewAdminRepository()
	getAdmin, err := repo.GetAdmin(request.ID)
	if err != nil {
		return nil, err
	}

	if getAdmin.ID == 0 {
		return nil, errors.New("admin not found")
	}

	if err := repo.UpdateAdmin(request.ID, model.Admin{
		Email:    request.Email,
		FullName: request.FullName,
	}); err != nil {
		return nil, err
	}

	return &types.ResponseUpdateAdmin{
		Message: "Admin updated.",
	}, nil
}

func (a adminUseCaseCtx) DeleteAdmin(request types.RequestGetAdmin) (*types.ResponseDeleteAdmin, error) {

	repo := repository.NewAdminRepository()
	getAdmin, err := repo.GetAdmin(request.ID)
	if err != nil {
		return nil, err
	}

	if getAdmin.ID == 0 {
		return nil, errors.New("admin not found")
	}

	if err := repo.DeleteAdmin(request.ID); err != nil {
		return nil, err
	}

	return &types.ResponseDeleteAdmin{
		Message: "Admin Deleted",
	}, nil
}

func (a adminUseCaseCtx) LoginAdmin(request types.RequestLoginAdmin) (*types.ResponseLoginAdmin, error) {

	repo := repository.NewAdminRepository()
	getAdmin, err := repo.FindByUsername(request.Username)
	if err != nil {
		return nil, err
	}

	if getAdmin.ID == 0 {
		return nil, errors.New("admin not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(getAdmin.Password), []byte(request.Password)); err != nil {
		return nil, errors.New("invalid password")
	}

	_toJWT := map[string]interface{}{
		"id":       strconv.Itoa(int(getAdmin.ID)),
		"acc_type": "admin",
		"email":    getAdmin.Email,
	}

	_jwt := jwt.NewJWTImpl(config.GetString("jwt.signature_key"), config.GetInt("jwt.day_expired"))
	token, err := _jwt.GenerateToken(_toJWT)
	if err != nil {
		log.Printf("Error generate token: %v", err.Error())
	}

	return &types.ResponseLoginAdmin{
		Token: token,
	}, nil
}

func (a adminUseCaseCtx) RegisterAdmin(request types.RequestCreatAdmin) (*types.ResponseCreateAdmin, error) {

	repo := repository.NewAdminRepository()
	if err := repo.CreateAdmin(model.Admin{
		Email:    request.Email,
		Username: request.Username,
		FullName: request.FullName,
		Password: utils.EncryptPassword(request.Password),
	}); err != nil {
		return nil, err
	}

	return &types.ResponseCreateAdmin{
		Message: "Register success",
	}, nil
}
