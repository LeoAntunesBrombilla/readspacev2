package interfaces

type AuthenticationUseCase interface {
	Login(username, password string) (token string, error error)
	Logout(token string) error
}
