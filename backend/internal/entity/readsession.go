package entity

import "time"

type ReadingSession struct {
	ID        int64
	UserID    int
	BookID    int64
	CreatedAt time.Time
	Durations []ReadingDuration // Assuming you have multiple durations per session
}

type ReadingSessionInput struct {
	BookID    int64                `json:"book_id"`
	Durations ReadingDurationInput `json:"durations"` // Assuming you have multiple durations per session
}

type ReadingDuration struct {
	Date time.Time `json:"date"`
	Time int       `json:"time"`
}

type ReadingDurationInput struct {
	Time int `json:"time"`
}
