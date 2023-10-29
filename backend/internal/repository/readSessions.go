package repository

import (
	"context"
	"time"
)

type ReadingSession struct {
    ID        int           `json:"id"`
    UserID    int           `json:"user_id"`
    BookID    int           `json:"book_id"`
    Duration  []ReadingDuration `json:"duration"`
    Date      string        `json:"date"`
    CreatedAt time.Time     `json:"created_at"`
}

type ReadingDuration struct {
    Date string `json:"date"`
    Time int    `json:"time"`
}

type ReadingSessionRepository interface {
    AddReadingSession(ctx context.Context, userId string, bookId string, session ReadingSession) error
    UpdateReadingSession(ctx context.Context, sessionId string, session ReadingSession) error
    GetReadingSessions(ctx context.Context, userId string, bookId string) ([]ReadingSession, error)
    GetReadingSessionsByDate(ctx context.Context, userId string, bookId string, startDate string, endDate string) ([]ReadingSession, error)
    DeleteReadingSession(ctx context.Context, sessionId string) error
}

