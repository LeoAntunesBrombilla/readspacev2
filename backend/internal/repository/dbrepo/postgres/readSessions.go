package postgres

import (
	"context"
	"fmt"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/entity"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/repository"
	"github.com/jackc/pgx/v4/pgxpool"
)

type readSessionsRepository struct {
	db *pgxpool.Pool
}

func NewReadSessionsRepository(db *pgxpool.Pool) repository.ReadingSessionRepository {
	return &readSessionsRepository{db: db}
}

func (r readSessionsRepository) CreatReadingSession(ctx context.Context, readingSession entity.ReadingSession) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		fmt.Println("Error starting transaction:", err)
		return err
	}

	defer tx.Rollback(ctx)

	var sessionId int64
	sessionInsertQuery := `INSERT INTO read_sessions (user_id, book_id, created_at) VALUES ($1, $2, $3) RETURNING id`
	err = tx.QueryRow(ctx, sessionInsertQuery, readingSession.UserID, readingSession.BookID, readingSession.CreatedAt).Scan(&sessionId)
	if err != nil {
		fmt.Println("Error inserting new reading session:", err)
		return err
	}

	for _, duration := range readingSession.Durations {
		durationInsertQuery := `INSERT INTO session_durations (session_id, date, duration) VALUES ($1, $2, $3)`
		_, err = tx.Exec(ctx, durationInsertQuery, sessionId, duration.Date, duration.Time)
		if err != nil {
			fmt.Println("Error inserting session duration:", err)
			return err
		}
	}

	return tx.Commit(ctx)
}

func (r readSessionsRepository) UpdateReadingSession(ctx context.Context, sessionId string, session entity.ReadingSession) error {
	//TODO implement me
	panic("implement me")
}

func (r readSessionsRepository) GetReadingSessions(ctx context.Context, userId string, bookId string) ([]entity.ReadingSession, error) {
	//TODO implement me
	panic("implement me")
}

func (r readSessionsRepository) GetReadingSessionsById(ctx context.Context, readingSession entity.ReadingSession) ([]entity.ReadingSession, error) {
	//TODO implement me
	panic("implement me")
}

func (r readSessionsRepository) DeleteReadingSession(ctx context.Context, sessionId string) error {
	//TODO implement me
	panic("implement me")
}
