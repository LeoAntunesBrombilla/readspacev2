package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"readspacev2/internal/auth"
	"readspacev2/internal/entity"
	"readspacev2/internal/usecase"
	"strconv"
	"strings"
)

type BookListHandler struct {
	bookListUseCase usecase.BookListUseCaseInterface
}

func NewBookListHandler(bookListUseCase usecase.BookListUseCaseInterface) *BookListHandler {
	return &BookListHandler{bookListUseCase: bookListUseCase}
}

// Create godoc
// @Summary Create a new bookList
// @Description Create a new bookList with the input payload
// @Tags bookList
// @Accept  json
// @Produce  json
// @Param user body entity.BookListInput true "BookList input for creation"
// @Success 201 {string} string "Created"
// @Failure 400 {object} entity.ErrorEntity
// @Failure 500 {object} entity.ErrorEntity
// @Router /bookList [post]
func (h *BookListHandler) Create(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	tokenClaims, err := auth.ParseToken(tokenString)

	if err != nil {
		c.JSON(http.StatusUnauthorized, entity.ErrorEntity{Code: 401, Message: "Unauthorized"})
		return
	}

	var bookListInput entity.BookList
	if err = c.ShouldBindJSON(&bookListInput); err != nil {
		c.JSON(http.StatusBadRequest, entity.ErrorEntity{Code: 400, Message: "Bad Request"})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, entity.ErrorEntity{Code: 400, Message: "Bad Request"})
		return
	}

	bookList := entity.BookList{
		UserID: tokenClaims.UserID,
		Name:   bookListInput.Name,
	}

	err = h.bookListUseCase.Create(&bookList)

	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.ErrorEntity{Code: 500, Message: "Error creating book list"})
		return
	}

	c.Status(http.StatusCreated)
	return
}

// ListAllBookLists godoc
// @Summary List all bookList
// @Description Retrieve a list of all bookLists in the database
// @Tags bookLists
// @Produce  json
// @Success 200 {array} entity.UserEntity
// @Failure 500 {object} entity.ErrorEntity
// @Router /bookList [get]
func (h *BookListHandler) ListAllBookLists(c *gin.Context) {
	bookLists, err := h.bookListUseCase.ListAllBookLists()

	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.ErrorEntity{Code: 500, Message: "Error returning book list"})
		return
	}

	c.JSON(http.StatusOK, bookLists)
	return
}

// DeleteBookListById godoc
// @Summary Delete a book list by ID
// @Description Delete the book list identified by the given ID
// @Tags bookList
// @Produce  json
// @Param id query int64 true "Book List ID to delete"
// @Success 200 {object} string "Book list deleted with success"
// @Failure 400 {object} entity.ErrorEntity
// @Failure 500 {object} entity.ErrorEntity
// @Router /bookList [delete]
func (h *BookListHandler) DeleteBookListById(c *gin.Context) {
	var id int64

	idStr := c.DefaultQuery("id", "0")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, entity.ErrorEntity{Code: 400, Message: "Invalid ID query parameter"})
		return
	}

	err = h.bookListUseCase.DeleteBookListById(&id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.ErrorEntity{Code: 500, Message: "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Book List deleted with success"})
	return
}

func (h *BookListHandler) UpdateBookList(c *gin.Context) {
	var id int64

	idStr := c.DefaultQuery("id", "0")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, entity.ErrorEntity{Code: 400, Message: "Invalid ID query parameter"})
		return
	}

	var bookListUpdateDetails entity.BookListDetails

	if err = c.ShouldBindJSON(&bookListUpdateDetails); err != nil {
		c.JSON(http.StatusBadRequest, entity.ErrorEntity{Code: 400, Message: "Invalid Payload"})
		return
	}

	err = h.bookListUseCase.UpdateBookList(&id, &bookListUpdateDetails)

	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.ErrorEntity{Code: 500, Message: "Error updating bookList"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "BookList updated with success"})
}
