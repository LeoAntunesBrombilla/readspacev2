package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
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

func convertTextArrayToStringSlice(textArray pgtype.TextArray) []string {
	if textArray.Status != pgtype.Present {
		return nil
	}

	result := make([]string, len(textArray.Elements))
	for i, elem := range textArray.Elements {
		if elem.Status == pgtype.Present {
			result[i] = elem.String
		}
	}
	return result
}

func (b *bookListRepository) ListAllBookLists() ([]*entity.BookList, error) {

	query := `
    SELECT bl.id, bl.user_id, bl.name, bl.created_at, bl.updated_at,
           b.id, b.title, b.subtitle, b.authors, b.publisher, b.description, 
           b.page_count, b.categories, b.language, b.small_thumbnail, b.thumbnail
    FROM book_lists bl
    LEFT JOIN book_list_books blb ON bl.id = blb.list_id
    LEFT JOIN books b ON blb.book_id = b.id
    `

	rows, err := b.db.Query(context.Background(), query)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var bookLists []*entity.BookList
	bookListMap := make(map[int64]*entity.BookList)

	for rows.Next() {
		var bookList entity.BookList
		var book entity.Book
		var bookID sql.NullInt64

		var title, subtitle, publisher, description, language, smallThumbnail, thumbnail sql.NullString
		var authorsArray, categoriesArray pgtype.TextArray
		var pageCount sql.NullInt64

		err = rows.Scan(
			&bookList.ID, &bookList.UserID, &bookList.Name,
			&bookList.CreatedAt, &bookList.UpdatedAt,
			&bookID, &title, &subtitle, &authorsArray,
			&publisher, &description, &pageCount,
			&categoriesArray, &language,
			&smallThumbnail, &thumbnail,
		)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		existingBookList, exists := bookListMap[bookList.ID]
		if !exists {
			bookList.Books = []*entity.Book{}
			existingBookList = &bookList
			bookListMap[bookList.ID] = existingBookList
			bookLists = append(bookLists, existingBookList)
		}

		if bookID.Valid {
			book.ID = bookID.Int64
			book.Title = title.String
			existingBookList.Books = append(existingBookList.Books, &book)
			if subtitle.Valid {
				book.Subtitle = subtitle.String
			}
			book.Authors = convertTextArrayToStringSlice(authorsArray)
			book.Categories = convertTextArrayToStringSlice(categoriesArray)
			if publisher.Valid {
				book.Publisher = publisher.String
			}
			if description.Valid {
				book.Description = description.String
			}
			if pageCount.Valid {
				book.PageCount = int(pageCount.Int64)
			}
			if language.Valid {
				book.Language = language.String
			}
			if smallThumbnail.Valid {
				book.ImageLinks.SmallThumbnail = smallThumbnail.String
			}
			if thumbnail.Valid {
				book.ImageLinks.Thumbnail = thumbnail.String
			}
		}
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
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

//TODO create GET bookList
