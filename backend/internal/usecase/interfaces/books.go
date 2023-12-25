package interfaces

import (
	"context"
	"readspacev2/internal/entity"
)

type BooksUseCaseInterface interface {
	Create(ctx context.Context, book *entity.Book) error
	Delete(id *int64) error
}
