package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"readspacev2/internal/entity"
	"readspacev2/internal/repository"
)

type bookListRepository struct {
	db *pgxpool.Pool
}

func NewBookListRepository(db *pgxpool.Pool) repository.BookListRepository {
	return &bookListRepository{db: db}
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
	query := `SELECT id, user_id, name, created_at, updated_at FROM book_lists`

	rows, err := b.db.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var bookLists []*entity.BookList

	for rows.Next() {
		var bookList entity.BookList

		if err := rows.Scan(&bookList.ID, bookList.UserID, bookList.Name, bookList.CreatedAt, bookList.UpdatedAt); err != nil {
			return nil, err
		}

		bookLists = append(bookLists, &bookList)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return bookLists, nil
}

func (b *bookListRepository) FindBookListByName(name string) (*entity.BookList, error) {
	//TODO implement me
	panic("implement me")
}

func (b *bookListRepository) Create(bookList *entity.BookList) error {

	query := `INSERT INTO book_lists (user_id, name) VALUES ($1, $2) RETURNING id`

	conn, err := b.db.Acquire(context.Background())

	if err != nil {
		return err
	}

	defer conn.Release()

	row := conn.QueryRow(context.Background(), query, bookList.UserID, bookList.Name)

	err = row.Scan(&bookList.ID)

	if err != nil {
		return err
	}

	return nil
}
