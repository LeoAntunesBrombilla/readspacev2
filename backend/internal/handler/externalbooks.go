package handler

import (
	"github.com/LeoAntunesBrombilla/readspacev2/internal/usecase"
	"net/http"

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
	query := c.DefaultQuery("q", "")

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
