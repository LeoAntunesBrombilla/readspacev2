package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"readspacev2/internal/entity"
	"readspacev2/internal/repository"
)

type booksRepository struct {
	db *pgxpool.Pool
}

func NewBooksRepository(db *pgxpool.Pool) repository.BooksRepository {
	return &booksRepository{db: db}
}

func (b booksRepository) AddBookToList(ctx context.Context, externalBook entity.ExternalBook, bookListId string) error {
	//TODO implement me
	panic("implement me")
}

func (b booksRepository) CreateBookList(ctx context.Context, listName string) (entity.BookList, error) {
	//TODO implement me
	panic("implement me")
}

func (b booksRepository) GetBookListByID(ctx context.Context, bookListId string) (entity.BookList, error) {
	//TODO implement me
	panic("implement me")
}

func (b booksRepository) DeleteBookList(ctx context.Context, bookListId string) error {
	//TODO implement me
	panic("implement me")
}

func (b booksRepository) DeleteBookFromList(ctx context.Context, bookListId string, bookId string) error {
	//TODO implement me
	panic("implement me")
}
