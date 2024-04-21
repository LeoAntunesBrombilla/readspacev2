package repository

import (
	"context"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/entity"
)

type ReadingSessionRepository interface {
	CreatReadingSession(ctx context.Context, readingSession entity.ReadingSession) error
	UpdateReadingSession(ctx context.Context, sessionId string, session entity.ReadingSession) error
	GetReadingSessions(ctx context.Context, userId string, bookId string) ([]entity.ReadingSession, error)
	GetReadingSessionsById(ctx context.Context, readingSession entity.ReadingSession) ([]entity.ReadingSession, error)
	DeleteReadingSession(ctx context.Context, sessionId string) error
}
