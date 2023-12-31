package postgres

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
	"readspacev2/internal/entity"
	"readspacev2/internal/repository"
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
		return err
	}
	defer tx.Rollback(ctx)

	var bookId int64
	err = tx.QueryRow(ctx, "SELECT id FROM books WHERE googleBookId = $1", book.GoogleBookId).Scan(&bookId)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err == sql.ErrNoRows {
		bookInsertQuery := `INSERT INTO books (googleBookId, title, subtitle, authors, publisher, description, page_count, categories, language, small_thumbnail, thumbnail) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`
		err = tx.QueryRow(ctx, bookInsertQuery, book.GoogleBookId, book.Title, book.Subtitle, pq.Array(book.Authors), book.Publisher, book.Description, book.PageCount, pq.Array(book.Categories), book.Language, book.ImageLinks.SmallThumbnail, book.ImageLinks.Thumbnail).Scan(&bookId)
		if err != nil {
			return err
		}
	}

	linkInsertQuery := `INSERT INTO book_list_books (list_id, book_id) VALUES ($1, $2)`
	_, err = tx.Exec(ctx, linkInsertQuery, book.BookListID, bookId)
	if err != nil {
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
