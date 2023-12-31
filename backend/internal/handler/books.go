package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"readspacev2/internal/entity"
	"readspacev2/internal/usecase/interfaces"
)

type BooksHandler struct {
	booksUseCase interfaces.BooksUseCaseInterface
}

func NewBooksHandler(booksUseCase interfaces.BooksUseCaseInterface) *BooksHandler {
	return &BooksHandler{booksUseCase: booksUseCase}
}

func (h *BooksHandler) Create(c *gin.Context) {
	var bookInput entity.ExternalBook

	err := c.ShouldBindJSON(&bookInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, entity.ErrorEntity{Code: 400, Message: "Bad Request"})
		return
	}

	book := entity.Book{
		BookListID:   bookInput.BookListID,
		GoogleBookId: bookInput.GoogleBookId,
		Title:        bookInput.Title,
		Subtitle:     bookInput.Subtitle,
		Authors:      bookInput.Authors,
		Publisher:    bookInput.Publisher,
		Description:  bookInput.Description,
		PageCount:    bookInput.PageCount,
		Categories:   bookInput.Categories,
		Language:     bookInput.Language,
		ImageLinks: struct {
			SmallThumbnail string
			Thumbnail      string
		}(bookInput.ImageLinks),
	}

	err = h.booksUseCase.Create(c, &book)

	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusInternalServerError, entity.ErrorEntity{Code: 500, Message: "Error creating book"})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *BooksHandler) Delete(c *gin.Context) {
	var deleteBook entity.DeleteBookInput

	err := c.ShouldBindJSON(&deleteBook)

	if err != nil {
		c.JSON(http.StatusBadRequest, entity.ErrorEntity{Code: 400, Message: "Bad Request"})
		return
	}

	err = h.booksUseCase.Delete(c, &deleteBook.BookListID, &deleteBook.BookID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.ErrorEntity{Code: 500, Message: "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Book deleted with success"})
	return
}
