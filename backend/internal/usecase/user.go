package usecase

import (
	"readspacev2/internal/entity"
	"readspacev2/internal/repository"
)

type UserUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (useCase *UserUseCase) CreateUser(user *entity.User) error {
	return useCase.repo.Create(user)
}

func (useCase *UserUseCase) ListAllUsers() ([]*entity.User, error) {
	return useCase.repo.ListAllUsers()
}

func (useCase *UserUseCase) DeleteUserById(id *int64) error {
	return useCase.repo.DeleteUserById(id)
}
