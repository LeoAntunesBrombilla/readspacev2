package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"golang.org/x/xerrors"
)

type AuthRepository struct {
	client *redis.Client
}

func (r *AuthRepository) RetrieveToken(ctx context.Context, userID string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AuthRepository) DeleteToken(ctx context.Context, userID string) error {
	//TODO implement me
	panic("implement me")
}

func (r *AuthRepository) VerifyToken(ctx context.Context, token string) (userID string, err error) {
	//TODO implement me
	panic("implement me")
}

func NewRedisAuthRepository(ctx context.Context) (*AuthRepository, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, xerrors.Errorf("failed to connect to redis: %v", err)
	}

	return &AuthRepository{client: client}, nil
}

func (r *AuthRepository) StoreToken(ctx context.Context, userId string, token string) error {
	return r.client.Set(ctx, userId, token, 0).Err()
}
