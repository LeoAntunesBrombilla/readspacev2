package repository

import "context"

type AuthenticationRepository interface {
	StoreToken(ctx context.Context, userID string, token string) error
	RetrieveToken(ctx context.Context, userID string) (string, error)
	DeleteToken(ctx context.Context, userID string) error
	VerifyToken(ctx context.Context, token string) (userID string, err error)
}
