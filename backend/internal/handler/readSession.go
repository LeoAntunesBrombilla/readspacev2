package handler

import (
	"github.com/LeoAntunesBrombilla/readspacev2/internal/auth"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/entity"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/usecase/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type ReadSessionHandler struct {
	readSessionUseCase interfaces.ReadSessionUseCaseInterface
}

func NewReadSessionHandler(readSessionUseCase interfaces.ReadSessionUseCaseInterface) *ReadSessionHandler {
	return &ReadSessionHandler{
		readSessionUseCase: readSessionUseCase,
	}
}

// CreateReadSession Create godoc
// @Summary Create a reading session
// @Description Create a new reading session with the input payload
// @Tags readSession
// @Accept  json
// @Produce  json
// @Param user body entity.ReadingSessionInput true "ReadingSession input for creation"
// @Success 201 {string} string "Created"
// @Failure 400 {object} entity.ErrorEntity
// @Failure 500 {object} entity.ErrorEntity
// @Router /readSession [post]
func (h *ReadSessionHandler) CreateReadSession(c *gin.Context) {
	var readSessionInput entity.ReadingSessionInput
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	tokenClaims, err := auth.ParseToken(tokenString)

	err = c.BindJSON(&readSessionInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, entity.ErrorEntity{Code: 400, Message: "Bad Request"})
		return
	}

	err = h.readSessionUseCase.CreatReadingSession(c, readSessionInput, tokenClaims.UserID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.ErrorEntity{Code: 500, Message: "Error creating reading session"})
		return
	}

	c.Status(http.StatusCreated)
}

// GetReadSessionByBook godoc
// @Summary List all reading sessions
// @Description Retrieve a list of all reading sessions in the database
// @Tags readSession
// @Produce  json
// @Success 200 {array} entity.ReadingSession
// @Failure 500 {object} entity.ErrorEntity
// @Router /readSession [get]
func (h *ReadSessionHandler) GetReadSessionByBook(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	tokenClaims, err := auth.ParseToken(tokenString)

	query := c.DefaultQuery("bookId", "")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter is missing"})
		return
	}

	if err != nil {
		c.JSON(http.StatusUnauthorized, entity.ErrorEntity{Code: 401, Message: "Unauthorized"})
		return
	}

	readSessions, err := h.readSessionUseCase.GetReadingSessionsByBook(c, tokenClaims.UserID, query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.ErrorEntity{Code: 500, Message: "Error returning list of reading sessions"})
		return
	}

	c.JSON(http.StatusOK, readSessions)
}
