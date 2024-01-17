package usecase

import (
	"github.com/LeoAntunesBrombilla/readspacev2/internal/entity"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/repository"
)

type UserUseCaseInterface interface {
	CreateUser(user *entity.UserEntity) error
	ListAllUsers() ([]*entity.UserEntity, error)
	DeleteUserById(id *int64) error
	FindByUserName(username string) (*entity.UserEntity, error)
	UpdateUser(id *int64, user *entity.UserUpdateDetails) error
	UpdateUserPassword(id *int64, password string) error
	FindPasswordById(id *int64) (*string, error)
}

type UserUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (useCase *UserUseCase) CreateUser(user *entity.UserEntity) error {
	return useCase.repo.Create(user)
}

func (useCase *UserUseCase) ListAllUsers() ([]*entity.UserEntity, error) {
	return useCase.repo.ListAllUsers()
}

func (useCase *UserUseCase) DeleteUserById(id *int64) error {
	return useCase.repo.DeleteUserById(id)
}

func (useCase *UserUseCase) FindByUserName(username string) (*entity.UserEntity, error) {
	return useCase.repo.FindByUserName(username)
}

func (useCase *UserUseCase) UpdateUser(id *int64, user *entity.UserUpdateDetails) error {
	return useCase.repo.UpdateUser(id, user)
}

func (useCase *UserUseCase) UpdateUserPassword(id *int64, password string) error {
	return useCase.repo.UpdateUserPassword(id, password)
}

func (useCase *UserUseCase) FindPasswordById(id *int64) (*string, error) {
	return useCase.repo.FindPasswordById(id)
}
