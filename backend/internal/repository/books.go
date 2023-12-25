package repository

import (
	"context"
	"readspacev2/internal/entity"
)

type BooksRepository interface {
	Create(c context.Context, book *entity.Book) error
	Delete(id *int64) error
}
