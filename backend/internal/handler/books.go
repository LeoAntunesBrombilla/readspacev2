package handler

import (
	"net/http"
	"readspacev2/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ExternalBookServiceHandler struct {
	externalBookServiceUseCase usecase.ExternalBookServiceUseCaseInterface
}

func NewExternalBookServiceHandler(externalBookServiceUseCase usecase.ExternalBookServiceUseCaseInterface) *ExternalBookServiceHandler {
	return &ExternalBookServiceHandler{
		externalBookServiceUseCase: externalBookServiceUseCase,
	}
}

func (h *ExternalBookServiceHandler) SearchBooks(c *gin.Context) {
	query := c.DefaultQuery("query", "")

  if query == "" {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter is missing"})
    return
  }	

  pagination := 10

	books, err := h.externalBookServiceUseCase.SearchBooks(c, query, pagination)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}
