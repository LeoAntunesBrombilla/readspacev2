package repository

import (
	"context"
	"readspacev2/internal/entity"
)

type ExternalBookRepository interface {
	SearchBooks(ctx context.Context, queryParam string, pagination int) ([]entity.ExternalBook, error)
}
