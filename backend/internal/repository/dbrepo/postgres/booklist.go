package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/entity"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/repository/interfaces"
	"github.com/jackc/pgx/v4/pgxpool"
)

type bookListRepository struct {
	db *pgxpool.Pool
}

func NewBookListRepository(db *pgxpool.Pool) interfaces.BookListRepository {
	return &bookListRepository{db: db}
}

func (b *bookListRepository) UpdateBookList(id *int64, bookList *entity.BookListDetails) error {
	query := `UPDATE book_lists SET name = $1 WHERE id = $2`
	_, err := b.db.Exec(context.Background(), query, bookList.Name, id)

	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error updating user: %w", err)
	}

	return nil
}

func (b *bookListRepository) DeleteBookListById(id *int64) error {
	query := `DELETE FROM book_lists WHERE id = $1`

	_, err := b.db.Exec(context.Background(), query, id)

	if err != nil {
		return err
	}

	return nil
}

func (b *bookListRepository) ListAllBookLists() ([]*entity.BookList, error) {

	query := `
		SELECT bl.id, bl.user_id, bl.name, bl.created_at, bl.updated_at,
       		b.id AS book_id, b.google_book_id, b.title, blb.list_id
		FROM book_lists bl
		LEFT JOIN book_list_books blb ON bl.id = blb.list_id
		LEFT JOIN books b ON blb.book_id = b.id
	`

	rows, err := b.db.Query(context.Background(), query)
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}
	defer rows.Close()

	var bookLists []*entity.BookList
	bookListMap := make(map[int64]*entity.BookList)

	for rows.Next() {
		var bookList entity.BookList
		var book entity.Book
		var bookID sql.NullInt64
		var bookListId sql.NullInt64
		var googleBookId sql.NullString // Changed to handle NULL values
		var title sql.NullString        // Changed to handle NULL values

		err = rows.Scan(
			&bookList.ID, &bookList.UserID, &bookList.Name,
			&bookList.CreatedAt, &bookList.UpdatedAt,
			&bookID, &googleBookId, &title, &bookListId,
		)
		if err != nil {
			fmt.Println("Scan error:", err)
			return nil, err
		}

		if existingBookList, exists := bookListMap[bookList.ID]; exists {
			bookList = *existingBookList
		} else {
			bookList.Books = []*entity.Book{}
			bookListMap[bookList.ID] = &bookList
			bookLists = append(bookLists, &bookList)
		}

		if bookID.Valid {
			book.ID = bookID.Int64
			book.GoogleBookId = googleBookId.String
			book.Title = title.String
			if googleBookId.Valid {
				book.GoogleBookId = googleBookId.String
			}
			if title.Valid {
				book.Title = title.String
			}
			if bookListId.Valid {
				book.BookListID = bookListId.Int64
			}
			bookList.Books = append(bookList.Books, &book)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookLists, nil
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
