package postgres

import (
	"context"
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

	bookInsertQuery := `INSERT INTO books (title, subtitle, authors, publisher, description, page_count, categories, language, small_thumbnail, thumbnail) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`
	err = tx.QueryRow(ctx, bookInsertQuery, book.Title, book.Subtitle, book.Authors, book.Publisher, book.Description, book.PageCount, pq.Array(book.Categories), book.Language, book.ImageLinks.SmallThumbnail, book.ImageLinks.Thumbnail).Scan(&book.ID)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	linkInsertQuery := `INSERT INTO book_list_books (list_id, book_id) VALUES ($1, $2)`
	_, err = tx.Exec(ctx, linkInsertQuery, book.BookListID, book.ID)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}

func (b booksRepository) Delete(id *int64) error {
	//TODO implement me
	panic("implement me")
}
