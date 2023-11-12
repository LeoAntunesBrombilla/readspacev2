package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"readspacev2/internal/auth"
	"readspacev2/internal/entity"
	"readspacev2/internal/usecase"
	"strings"
)

type BookListHandler struct {
	bookListUseCase usecase.BookListUseCaseInterface
}

func NewBookListHandler(bookListUseCase usecase.BookListUseCaseInterface) *BookListHandler {
	return &BookListHandler{bookListUseCase: bookListUseCase}
}

func (h *BookListHandler) Create(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	tokenClaims, err := auth.ParseToken(tokenString)

	if err != nil {
		c.JSON(http.StatusUnauthorized, entity.ErrorEntity{Code: 401, Message: "Unauthorized"})
		return
	}

	var bookListInput entity.BookList
	if err := c.ShouldBindJSON(&bookListInput); err != nil {
		c.JSON(http.StatusBadRequest, entity.ErrorEntity{Code: 400, Message: "Bad Request"})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, entity.ErrorEntity{Code: 400, Message: "Bad Request"})
	}

	bookList := entity.BookList{
		UserID: tokenClaims.UserID,
		Name:   bookListInput.Name,
	}

	err = h.bookListUseCase.Create(&bookList)

	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.ErrorEntity{Code: 500, Message: "Error creating book list"})
	}

	c.Status(http.StatusCreated)
}

func (h *BookListHandler) ListAllBookLists(c *gin.Context) {
	bookLists, err := h.bookListUseCase.ListAllBookLists()

	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.ErrorEntity{Code: 500, Message: "Error returning book list"})
	}

	c.JSON(http.StatusOK, bookLists)
}
