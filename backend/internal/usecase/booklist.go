package usecase

import (
	"readspacev2/internal/entity"
	"readspacev2/internal/repository"
)

type BookListUseCaseInterface interface {
	Create(bookList *entity.BookList) error
	UpdateBookList(id *int64, bookList *entity.BookListDetails) error
	DeleteBookListById(id *int64) error
	ListAllBookLists() ([]*entity.BookList, error)
}

type BookListUseCase struct {
	repo repository.BookListRepository
}

func NewBookListUseCase(repo repository.BookListRepository) *BookListUseCase {
	return &BookListUseCase{repo: repo}
}

func (useCase *BookListUseCase) Create(bookList *entity.BookList) error {
	return useCase.repo.Create(bookList)
}

func (useCase *BookListUseCase) UpdateBookList(id *int64, bookList *entity.BookListDetails) error {
	return useCase.repo.UpdateBookList(id, bookList)
}

func (useCase *BookListUseCase) DeleteBookListById(id *int64) error {
	return useCase.repo.DeleteBookListById(id)
}

func (useCase *BookListUseCase) ListAllBookLists() ([]*entity.BookList, error) {
	return useCase.repo.ListAllBookLists()
}
