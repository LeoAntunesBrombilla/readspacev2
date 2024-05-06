package interfaces

import (
	"context"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/entity"
)

type ReadSessionUseCaseInterface interface {
	CreatReadingSession(ctx context.Context, readingSession entity.ReadingSessionInput, userId int) error
	GetReadingSessionsByBook(ctx context.Context, userId int, bookId string) (*entity.ReadingSessionModel, error)
}
