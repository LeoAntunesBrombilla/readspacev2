package repository

import "readspacev2/internal/entity"

type UserRepository interface {
	Create(user *entity.UserEntity) error
	GetByID(id int64) (*entity.UserEntity, error)
	UpdateUser(id *int64, user *entity.UserUpdateDetails) error
	DeleteUserById(id *int64) error
	ListAllUsers() ([]*entity.UserEntity, error)
	FindByUserName(string) (*entity.UserEntity, error)
	UpdatePassword(id *int64, password string) error
}
