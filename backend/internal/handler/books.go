package handler

import (
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
		BookListID:  bookInput.BookListID,
		Title:       bookInput.Title,
		Subtitle:    bookInput.Subtitle,
		Authors:     bookInput.Authors,
		Publisher:   bookInput.Publisher,
		Description: bookInput.Description,
		PageCount:   bookInput.PageCount,
		Categories:  bookInput.Categories,
		Language:    bookInput.Language,
		ImageLinks: struct {
			SmallThumbnail string
			Thumbnail      string
		}(bookInput.ImageLinks),
	}

	err = h.booksUseCase.Create(c, &book)

	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.ErrorEntity{Code: 500, Message: "Error creating user"})
		return
	}

	c.Status(http.StatusCreated)
}
