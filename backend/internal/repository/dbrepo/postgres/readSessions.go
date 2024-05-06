package postgres

import (
	"context"
	"fmt"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/entity"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/repository/interfaces"
	"github.com/jackc/pgx/v4/pgxpool"
)

type readSessionsRepository struct {
	db *pgxpool.Pool
}

func NewReadSessionsRepository(db *pgxpool.Pool) interfaces.ReadingSessionRepository {
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
	readSessionInsertQuery := `INSERT INTO read_sessions (user_id, book_id, created_at) VALUES ($1, $2, $3) RETURNING id`
	err = tx.QueryRow(ctx, readSessionInsertQuery, readingSession.UserID, readingSession.BookID, readingSession.CreatedAt).Scan(&sessionId)
	if err != nil {
		fmt.Println("Error inserting new reading session:", err)
		return err
	}

	readingTimeInsertQuery := `INSERT INTO reading_time (session_id, date, reading_time) VALUES ($1, $2, $3)`
	_, err = tx.Exec(ctx, readingTimeInsertQuery, sessionId, readingSession.ReadingTime.Date, readingSession.ReadingTime.Time)
	if err != nil {
		fmt.Println("Error inserting reading time:", err)
		return err
	}

	return tx.Commit(ctx)
}

func (r readSessionsRepository) GetReadingSessionsByBook(ctx context.Context, userId int, bookId string) ([]entity.ReadingSession, error) {
	query := `
		SELECT rs.id, rs.user_id, rs.book_id, rs.created_at, rt.date, rt.reading_time
		FROM read_sessions rs
		LEFT JOIN reading_time rt ON rs.id = rt.session_id
		WHERE ($1 = 0 OR rs.user_id = $1) AND ($2 = 0 OR rs.book_id = $2)
	`
	rows, err := r.db.Query(ctx, query, userId, bookId)
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	defer rows.Close()

	var readingSessions []entity.ReadingSession
	for rows.Next() {
		var readingSession entity.ReadingSession
		var readingTime entity.ReadingTime
		err = rows.Scan(&readingSession.ID, &readingSession.UserID, &readingSession.BookID, &readingSession.CreatedAt, &readingTime.Date, &readingTime.Time)
		if err != nil {
			fmt.Println("Error scanning reading session:", err)
			return nil, err
		}
		readingSession.ReadingTime = readingTime
		readingSessions = append(readingSessions, readingSession)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Error scanning reading session:", err)
		return nil, err
	}

	return readingSessions, nil
}

func (r readSessionsRepository) GetReadingSessionsById(ctx context.Context, readingSession entity.ReadingSession) ([]entity.ReadingSession, error) {
	//TODO implement me
	panic("implement me")
}

func (r readSessionsRepository) GetAllReadingSessions(ctx context.Context) ([]entity.ReadingSession, error) {
	//TODO implement me
	panic("implement me")
}

func (r readSessionsRepository) DeleteReadingSession(ctx context.Context, sessionId string) error {
	//TODO implement me
	panic("implement me")
}
