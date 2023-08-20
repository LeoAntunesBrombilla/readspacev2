package repository

import "readspacev2/internal/entity"

type UserRepository interface {
	Create(user *entity.User) error
	GetByID(id int64) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id int64) error
	ListAll() ([]*entity.User, error)
	FindByUserName(string) (*entity.User, error)
}
