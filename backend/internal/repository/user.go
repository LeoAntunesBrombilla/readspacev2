package repository

import "readspacev2/internal/entity"

type UserRepository interface {
	Create(user *entity.User) error
	GetByID(id int64) (*entity.User, error)
	Update(user *entity.User) error
	DeleteUserById(id *int64) error
	ListAllUsers() ([]*entity.User, error)
	FindByUserName(string) (*entity.User, error)
}
