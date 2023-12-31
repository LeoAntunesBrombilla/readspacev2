package usecase

import (
	"context"
	"readspacev2/internal/entity"
	"readspacev2/internal/repository"
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
