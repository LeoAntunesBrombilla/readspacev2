package usecase

import (
	"context"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/entity"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/repository/interfaces"
	"time"
)

type ReadSessionUseCase struct {
	repo interfaces.ReadingSessionRepository
}

func NewReadingSessionUseCase(repo interfaces.ReadingSessionRepository) *ReadSessionUseCase {
	return &ReadSessionUseCase{
		repo: repo,
	}
}

func (useCase *ReadSessionUseCase) CreatReadingSession(ctx context.Context, input entity.ReadingSessionInput, userID int) error {

	readSession := entity.ReadingSession{
		UserID:      userID,
		BookID:      input.BookID,
		CreatedAt:   time.Now(),
		ReadingTime: entity.ReadingTime{Date: time.Now(), Time: input.ReadingTime.Time},
	}

	return useCase.repo.CreatReadingSession(ctx, readSession)
}

func (useCase *ReadSessionUseCase) GetReadingSessionsByBook(ctx context.Context, userId int, bookId string) ([]entity.ReadingSession, error) {
	return useCase.repo.GetReadingSessionsByBook(ctx, userId, bookId)
}
