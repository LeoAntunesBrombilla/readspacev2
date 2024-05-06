package interfaces

import "github.com/LeoAntunesBrombilla/readspacev2/internal/entity"

type BookListRepository interface {
	Create(bookList *entity.BookList) error
	UpdateBookList(id *int64, bookList *entity.BookListDetails) error
	DeleteBookListById(id *int64) error
	ListAllBookLists() ([]*entity.BookList, error)
}
