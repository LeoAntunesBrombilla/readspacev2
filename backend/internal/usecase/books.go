package usecase

import (
	"context"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/entity"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/repository"
)

type BooksUseCase struct {
	repo repository.BooksRepository
}

func NewBooksUseCase(repo repository.BooksRepository) *BooksUseCase {
	return &BooksUseCase{
		repo: repo,
	}
}

func (useCase *BooksUseCase) Create(ctx context.Context, book *entity.Book) error {
	return useCase.repo.Create(ctx, book)
}

func (useCase *BooksUseCase) Delete(c context.Context, bookListId *int64, bookId *int64) error {
	return useCase.repo.Delete(c, bookListId, bookId)
}
