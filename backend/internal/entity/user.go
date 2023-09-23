package entity

import "time"

type UserEntity struct {
	ID        int64
	Email     string
	Username  string
	Password  string
	CreatedAt time.Time
}

type UserUpdateDetails struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
