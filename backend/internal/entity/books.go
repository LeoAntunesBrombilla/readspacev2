package entity

import "time"

type ImageLinks struct {
	SmallThumbnail string `json:"smallThumbnail"`
	Thumbnail      string `json:"thumbnail"`
}

type ExternalBook struct {
	BookListID  int64    `json:"bookListId"`
	Title       string   `json:"title"`
	Subtitle    string   `json:"subtitle"`
	Authors     []string `json:"authors"`
	Publisher   string   `json:"publisher"`
	Description string   `json:"description"`
	PageCount   int      `json:"pageCount"`
	Categories  []string `json:"categories"`
	Language    string   `json:"language"`
	ImageLinks  struct {
		SmallThumbnail string `json:"smallThumbnail"`
		Thumbnail      string `json:"thumbnail"`
	} `json:"imageLinks"`
}

type ExternalBookResponse struct {
	Title       string   `json:"title"`
	Subtitle    string   `json:"subtitle"`
	Authors     []string `json:"authors"`
	Publisher   string   `json:"publisher"`
	Description string   `json:"description"`
	PageCount   int      `json:"pageCount"`
	Categories  []string `json:"categories"`
	Language    string   `json:"language"`
	ImageLinks  struct {
		SmallThumbnail string `json:"smallThumbnail"`
		Thumbnail      string `json:"thumbnail"`
	} `json:"imageLinks"`
}

type Book struct {
	ID          int64
	BookListID  int64
	Title       string
	Subtitle    string
	Authors     []string
	Publisher   string
	Description string
	PageCount   int
	Categories  []string
	Language    string
	ImageLinks  struct {
		SmallThumbnail string
		Thumbnail      string
	}
}

type BookList struct {
	ID        int64     `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type BookListInput struct {
	Name string `json:"name" db:"name"`
}

type BookListDetails struct {
	Name string `json:"name" db:"name"`
}
