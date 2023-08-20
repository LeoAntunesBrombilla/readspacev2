package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"readspacev2/internal/entity"
	"time"
)

func CreateToken(user *entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"id":       user.ID,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}
