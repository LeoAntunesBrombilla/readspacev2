package entity

import "time"

type ReadingSession struct {
	ID          int64
	UserID      int
	BookID      int64
	CreatedAt   time.Time
	ReadingTime ReadingTime
}

type ReadingSessionInput struct {
	BookID      int64                `json:"book_id"`
	ReadingTime ReadingDurationInput `json:"reading_time"`
}

type ReadingTime struct {
	Date time.Time `json:"date"`
	Time int       `json:"time"`
}

type ReadingDurationInput struct {
	Time int `json:"time"`
}
