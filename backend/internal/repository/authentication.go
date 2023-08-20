package repository

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
)

type AuthenticationRepository interface {
	StoreToken(ctx context.Context, userID string, token string) error
	RetrieveToken(ctx context.Context, userID string) (string, error)
	DeleteToken(ctx context.Context, userID string) error
	VerifyToken(tokenString string) (*jwt.Token, error)
}
