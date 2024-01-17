package interfaces

import (
	"context"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/entity"
)

type BooksUseCaseInterface interface {
	Create(ctx context.Context, book *entity.Book) error
	Delete(c context.Context, bookListId *int64, bookId *int64) error
}
