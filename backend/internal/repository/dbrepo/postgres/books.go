package postgres

import (
	"context"
	"fmt"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/entity"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/repository"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type booksRepository struct {
	db *pgxpool.Pool
}

func NewBooksRepository(db *pgxpool.Pool) repository.BooksRepository {
	return &booksRepository{db: db}
}

func (b booksRepository) Create(ctx context.Context, book *entity.Book) error {
	tx, err := b.db.Begin(ctx)
	if err != nil {
		fmt.Println("Error starting transaction:", err)
		return err
	}
	defer tx.Rollback(ctx)

	var bookId int64
	err = tx.QueryRow(ctx, "SELECT id FROM books WHERE google_book_id = $1", book.GoogleBookId).Scan(&bookId)

	if err != nil {
		if err == pgx.ErrNoRows {
			bookInsertQuery := `INSERT INTO books (google_book_id, title, created_at) VALUES ($1, $2, $3) RETURNING id`
			err = tx.QueryRow(ctx, bookInsertQuery, book.GoogleBookId, book.Title, book.CreatedAt).Scan(&bookId)
			if err != nil {
				fmt.Println("Error inserting new book:", err)
				return err
			}
		} else {
			fmt.Println("Error querying book ID:", err)
			return err
		}
	}

	// Insert link between the book and the book list
	linkInsertQuery := `INSERT INTO book_list_books (list_id, book_id) VALUES ($1, $2)`
	_, err = tx.Exec(ctx, linkInsertQuery, book.BookListID, bookId)
	if err != nil {
		fmt.Println("Error linking book to book list:", err)
		return err
	}

	return tx.Commit(ctx)
}

func (b booksRepository) Delete(ctx context.Context, bookListId *int64, bookId *int64) error {
	tx, err := b.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	deleteQuery := `DELETE FROM book_list_books WHERE list_id = $1 AND book_id = $2`
	_, err = tx.Exec(ctx, deleteQuery, *bookListId, *bookId)
	if err != nil {
		return err
	}

	var count int
	checkBookIfNotInBookList := `SELECT COUNT(*) FROM book_list_books WHERE book_id = $1`
	err = tx.QueryRow(ctx, checkBookIfNotInBookList, *bookId).Scan(&count)
	if err != nil {
		return err
	}

	removeFromBooks := `DELETE FROM books WHERE id = $1`
	if count == 0 {
		_, err = tx.Exec(ctx, removeFromBooks, *bookId)
		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}
