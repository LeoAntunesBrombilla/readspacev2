package interfaces

import (
	"context"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/entity"
)

type ReadSessionUseCaseInterface interface {
	CreatReadingSession(ctx context.Context, readingSession entity.ReadingSession) error
}
