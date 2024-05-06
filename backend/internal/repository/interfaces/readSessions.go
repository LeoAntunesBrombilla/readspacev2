package interfaces

import (
	"context"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/entity"
)

type ReadingSessionRepository interface {
	CreatReadingSession(ctx context.Context, readingSession entity.ReadingSession) error
	GetReadingSessionsByBook(ctx context.Context, userId int, bookId string) ([]entity.ReadingSession, error)
	GetReadingSessionsById(ctx context.Context, readingSession entity.ReadingSession) ([]entity.ReadingSession, error)
	GetAllReadingSessions(ctx context.Context) ([]entity.ReadingSession, error)
	DeleteReadingSession(ctx context.Context, sessionId string) error
}
