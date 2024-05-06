package interfaces

import "github.com/LeoAntunesBrombilla/readspacev2/internal/entity"

type UserRepository interface {
	Create(user *entity.UserEntity) error
	UpdateUser(id *int64, user *entity.UserUpdateDetails) error
	DeleteUserById(id *int64) error
	ListAllUsers() ([]*entity.UserEntity, error)
	FindByUserName(string) (*entity.UserEntity, error)
	UpdateUserPassword(id *int64, password string) error
	FindPasswordById(id *int64) (*string, error)
}
