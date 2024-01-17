package postgres

import (
	"context"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/repository"
	"github.com/jackc/pgx/v4/pgxpool"
)

type readSessionsRepository struct {
	db *pgxpool.Pool
}

func NewReadSessionsRepository(db *pgxpool.Pool) repository.ReadingSessionRepository {
	return &readSessionsRepository{db: db}
}

func (r readSessionsRepository) AddReadingSession(ctx context.Context, userId string, bookId string, session repository.ReadingSession) error {
	//TODO implement me
	panic("implement me")
}

func (r readSessionsRepository) UpdateReadingSession(ctx context.Context, sessionId string, session repository.ReadingSession) error {
	//TODO implement me
	panic("implement me")
}

func (r readSessionsRepository) GetReadingSessions(ctx context.Context, userId string, bookId string) ([]repository.ReadingSession, error) {
	//TODO implement me
	panic("implement me")
}

func (r readSessionsRepository) GetReadingSessionsByDate(ctx context.Context, userId string, bookId string, startDate string, endDate string) ([]repository.ReadingSession, error) {
	//TODO implement me
	panic("implement me")
}

func (r readSessionsRepository) DeleteReadingSession(ctx context.Context, sessionId string) error {
	//TODO implement me
	panic("implement me")
}
