package entity

import "time"

type BookList struct {
	ID        int64     `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Books     []*Book   `json:"books"`
}

type BookListInput struct {
	Name string `json:"name" db:"name"`
}

type BookListDetails struct {
	Name string `json:"name" db:"name"`
}
