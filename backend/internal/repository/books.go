package repository

import (
	"context"
	"readspacev2/internal/entity"
)

type BooksRepository interface {
	Create(c context.Context, book *entity.Book) error
	Delete(c context.Context, bookListId *int64, bookId *int64) error
}
