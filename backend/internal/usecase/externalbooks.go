package usecase

import (
	"context"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/entity"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/repository"
)

type ExternalBookServiceUseCaseInterface interface {
	SearchBooks(ctx context.Context, queryParam string, pagination int) ([]entity.ExternalBookResponse, error)
}

type ExternalBookServiceUseCase struct {
	repo repository.ExternalBookRepository
}

func NewExternalBookServiceUseCase(repo repository.ExternalBookRepository) *ExternalBookServiceUseCase {
	return &ExternalBookServiceUseCase{repo: repo}
}

func (useCase *ExternalBookServiceUseCase) SearchBooks(ctx context.Context, queryParam string, pagination int) ([]entity.ExternalBookResponse, error) {
	return useCase.repo.SearchBooks(ctx, queryParam, pagination)
}
