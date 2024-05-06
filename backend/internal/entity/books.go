package entity

import "time"

type ImageLinks struct {
	SmallThumbnail string `json:"smallThumbnail"`
	Thumbnail      string `json:"thumbnail"`
}

type ExternalBook struct {
	BookListID   int64    `json:"bookListId"`
	GoogleBookId string   `json:"googleBookId"`
	Title        string   `json:"title"`
	Subtitle     string   `json:"subtitle"`
	Authors      []string `json:"authors"`
	Publisher    string   `json:"publisher"`
	Description  string   `json:"description"`
	PageCount    int      `json:"pageCount"`
	Categories   []string `json:"categories"`
	Language     string   `json:"language"`
	ImageLinks   struct {
		SmallThumbnail string `json:"smallThumbnail"`
		Thumbnail      string `json:"thumbnail"`
	} `json:"imageLinks"`
}

type ExternalBookResponse struct {
	Title        string   `json:"title"`
	GoogleBookId string   `json:"id"`
	Subtitle     string   `json:"subtitle"`
	Authors      []string `json:"authors"`
	Publisher    string   `json:"publisher"`
	Description  string   `json:"description"`
	PageCount    int      `json:"pageCount"`
	Categories   []string `json:"categories"`
	Language     string   `json:"language"`
	ImageLinks   struct {
		SmallThumbnail string `json:"smallThumbnail"`
		Thumbnail      string `json:"thumbnail"`
	} `json:"imageLinks"`
}

type Book struct {
	ID           int64  `json:"id"`
	BookListID   int64  `json:"bookListId"`
	GoogleBookId string `json:"googleBookId"`
	Title        string `json:"title"`
	CreatedAt    time.Time
}

type BookResponseModel struct {
	ID           int64
	BookListID   int64
	GoogleBookId string
	Title        string
	Subtitle     string
	Authors      []string
	Publisher    string
	Description  string
	PageCount    int
	Categories   []string
	Language     string
	ImageLinks   struct {
		SmallThumbnail string
		Thumbnail      string
	}
}

type DeleteBookInput struct {
	BookID     int64 `json:"book_id"`
	BookListID int64 `json:"book_list_id"`
}
