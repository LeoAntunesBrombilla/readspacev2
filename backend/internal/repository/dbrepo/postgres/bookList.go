package postgres

import (
	"context"
	"readspacev2/internal/entity"
	"readspacev2/internal/repository"

	"github.com/jackc/pgx/v4/pgxpool"
)

type bookListRepository struct {
	db *pgxpool.Pool
}

func (b *bookListRepository) UpdateBookList(id *int64, bookList *entity.BookListDetails) error {
	//TODO implement me
	panic("implement me")
}

func (b *bookListRepository) DeleteBookListById(id *int64) error {
	//TODO implement me
	panic("implement me")
}

func (b *bookListRepository) ListAllBookLists() ([]*entity.BookList, error) {
	//TODO implement me
	panic("implement me")
}

func (b *bookListRepository) FindBookListByName(name string) (*entity.BookList, error) {
	//TODO implement me
	panic("implement me")
}

func NewBookListRepository(db *pgxpool.Pool) repository.BookListRepository {
	return &bookListRepository{db: db}
}

func (b *bookListRepository) Create(bookList *entity.BookList) error {

	query := `INSERT INTO book_lists (user_id, name, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`

	conn, err := b.db.Acquire(context.Background())

	if err != nil {
		return err
	}

	defer conn.Release()

	row := conn.QueryRow(context.Background(), query, bookList.UserID, bookList.Name, bookList.CreatedAt, bookList.UpdatedAt)

	err = row.Scan(&bookList.ID)

	if err != nil {
		return err
	}

	return nil
}
