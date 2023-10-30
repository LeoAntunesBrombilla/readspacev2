package usecase

import (
	"context"
	"readspacev2/internal/entity"
	"readspacev2/internal/repository"
)

type ExternalBookServiceUseCaseInterface interface {
	SearchBooks(ctx context.Context, queryParam string, pagination int) ([]entity.ExternalBook, error)
}

type ExternalBookServiceUseCase struct {
	repo repository.ExternalBookRepository
}

func NewExternalBookServiceUseCase(repo repository.ExternalBookRepository) *ExternalBookServiceUseCase {
	return &ExternalBookServiceUseCase{repo: repo}
}

func (useCase *ExternalBookServiceUseCase) SearchBooks(ctx context.Context, queryParam string, pagination int) ([]entity.ExternalBook, error) {
	return useCase.repo.SearchBooks(ctx, queryParam, pagination)
}
