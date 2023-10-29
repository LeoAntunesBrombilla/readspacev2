package repository

import (
	"context"
	"time"
)

type ExternalBook struct {
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

type BookList struct {
	ID          int           `json:"id" db:"id"`
	Name        string        `json:"name" db:"name"`
	CreatedAt   time.Time     `json:"created_at" db:"created_at"`
	Books       []ExternalBook `json:"books"`
}

type ExternalBookService interface {
    SearchBooks(ctx context.Context, queryParam string, pagination int) ([]ExternalBook, error)
    GetBookByID(ctx context.Context, bookId string) (ExternalBook, error)
}

type BooksRepository interface {
    AddBookToList(ctx context.Context, externalBook ExternalBook, bookListId string) error
    CreateBookList(ctx context.Context, listName string) (BookList, error)
    GetBookListByID(ctx context.Context, bookListId string) (BookList, error)
    // UpdateBookList(ctx context.Context, bookListId string, newInfo BookListUpdateInfo) error
    DeleteBookList(ctx context.Context, bookListId string) error
    DeleteBookFromList(ctx context.Context, bookListId string, bookId string) error
}

