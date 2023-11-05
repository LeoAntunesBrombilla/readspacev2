package auth

import (
	"errors"
	"os"
	"readspacev2/internal/entity"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(user *entity.UserEntity) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"id":       user.ID,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

type TokenClaims struct {
	UserID int `json:"id"`
}

func ParseToken(tokenString string) (TokenClaims, error) {
	claims := TokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return claims, err
	}

	if claimsData, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		userID, ok := (*claimsData)["id"].(float64)
		if !ok {
			return claims, errors.New("invalid token claims")
		}
		claims.UserID = int(userID)
		return claims, nil
	}

	return claims, errors.New("invalid token")
}
