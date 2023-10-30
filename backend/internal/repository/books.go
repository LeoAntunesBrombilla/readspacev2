package repository

import (
	"context"
	"readspacev2/internal/entity"
)

type BooksRepository interface {
    AddBookToList(ctx context.Context, externalBook entity.ExternalBook, bookListId string) error
    CreateBookList(ctx context.Context, listName string) (entity.BookList, error)
    GetBookListByID(ctx context.Context, bookListId string) (entity.BookList, error)
    // UpdateBookList(ctx context.Context, bookListId string, newInfo BookListUpdateInfo) error
    DeleteBookList(ctx context.Context, bookListId string) error
    DeleteBookFromList(ctx context.Context, bookListId string, bookId string) error
}

type ExternalBookRepository interface {
    SearchBooks(ctx context.Context, queryParam string, pagination int) ([]entity.ExternalBook, error)
}
