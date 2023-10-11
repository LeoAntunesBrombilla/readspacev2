package entity

import "golang.org/x/crypto/bcrypt"

type RealBcryptWrapper struct{}

func (rbw RealBcryptWrapper) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, cost)
}

func (rbw RealBcryptWrapper) CompareHashAndPassword(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}
