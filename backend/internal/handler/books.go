package handler

import (
	"fmt"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/entity"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/usecase/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type BooksHandler struct {
	booksUseCase interfaces.BooksUseCaseInterface
}

func NewBooksHandler(booksUseCase interfaces.BooksUseCaseInterface) *BooksHandler {
	return &BooksHandler{booksUseCase: booksUseCase}
}

// Create godoc
// @Summary Create a new book
// @Description Create a new book with the input payload
// @Tags books
// @Accept  json
// @Produce  json
// @Param user body entity.ExternalBook true "Book input for creation"
// @Success 201 {string} string "Created"
// @Failure 400 {object} entity.ErrorEntity
// @Failure 500 {object} entity.ErrorEntity
// @Router /books [post]
func (h *BooksHandler) Create(c *gin.Context) {
	var bookInput entity.Book

	err := c.ShouldBindJSON(&bookInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, entity.ErrorEntity{Code: 400, Message: "Bad Request"})
		return
	}

	book := entity.Book{
		BookListID:   bookInput.BookListID,
		GoogleBookId: bookInput.GoogleBookId,
		Title:        bookInput.Title,
		CreatedAt:    time.Now(),
	}

	err = h.booksUseCase.Create(c, &book)

	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusInternalServerError, entity.ErrorEntity{Code: 500, Message: "Error creating book"})
		return
	}

	c.Status(http.StatusCreated)
}

// Delete godoc
// @Summary Delete a book
// @Description Delete a book with the input payload
// @Tags books
// @Accept  json
// @Produce  json
// @Param user body entity.DeleteBookInput true "Book input for deletion"
// @Success 200 {string} string "Deleted"
// @Failure 400 {object} entity.ErrorEntity
// @Failure 500 {object} entity.ErrorEntity
// @Router /books [delete]
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
