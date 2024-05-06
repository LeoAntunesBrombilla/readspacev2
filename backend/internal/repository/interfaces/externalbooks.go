package interfaces

import (
	"context"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/entity"
)

type ExternalBookRepository interface {
	SearchBooks(ctx context.Context, queryParam string, pagination int) ([]entity.ExternalBookResponse, error)
}
