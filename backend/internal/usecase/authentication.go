package usecase

import (
	"errors"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/auth"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/repository/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationUseCase interface {
	Login(username, password string) (token string, error error)
	Logout(token string) error
}

type AuthenticationUseCaseImpl struct {
	authRepo interfaces.AuthenticationRepository
	userRepo interfaces.UserRepository
}

func NewAuthenticationUseCase(authRepo interfaces.AuthenticationRepository, userRepo interfaces.UserRepository) AuthenticationUseCase {
	return &AuthenticationUseCaseImpl{
		authRepo: authRepo,
		userRepo: userRepo,
	}
}

func (a *AuthenticationUseCaseImpl) Login(username, password string) (string, error) {

	user, err := a.userRepo.FindByUserName(username)

	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("incorrect password")
	}

	token, err := auth.CreateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *AuthenticationUseCaseImpl) Logout(token string) error {
	//TODO implement me
	panic("implement me")
}
