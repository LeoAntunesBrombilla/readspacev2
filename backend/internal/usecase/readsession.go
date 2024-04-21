package usecase

import (
	"context"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/entity"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/repository"
)

type ReadSessionUseCase struct {
	repo repository.ReadingSessionRepository
}

func NewReadingSessionUseCase(repo repository.ReadingSessionRepository) *ReadSessionUseCase {
	return &ReadSessionUseCase{
		repo: repo,
	}
}

func (useCase *ReadSessionUseCase) CreatReadingSession(ctx context.Context, readingSession entity.ReadingSession) error {
	return useCase.repo.CreatReadingSession(ctx, readingSession)
}
