package entity

import "time"

type UserEntity struct {
	ID        int64
	Email     string
	Username  string
	Password  string
	CreatedAt time.Time
}

type UserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateDetails struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserUpdatePassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
