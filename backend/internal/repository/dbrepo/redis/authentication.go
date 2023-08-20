package redis

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"golang.org/x/xerrors"
	"os"
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

func (r *AuthRepository) VerifyToken(tokenString string) (*jwt.Token, error) {
	secretKey := os.Getenv("SECRET_KEY")
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the alg
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
}

func NewRedisAuthRepository(ctx context.Context) (*AuthRepository, error) {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, xerrors.Errorf("failed to connect to redis: %v", err)
	}

	return &AuthRepository{client: client}, nil
}

func (r *AuthRepository) StoreToken(ctx context.Context, userId string, token string) error {
	return r.client.Set(ctx, userId, token, 0).Err()
}
