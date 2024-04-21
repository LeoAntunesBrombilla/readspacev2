package handler

import (
	"fmt"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/auth"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/entity"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/usecase/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
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

	fmt.Println(tokenClaims)

	if err != nil {
		c.JSON(http.StatusBadRequest, entity.ErrorEntity{Code: 400, Message: "Bad Request"})
		return
	}

	readSession := entity.ReadingSession{
		UserID:    tokenClaims.UserID,
		BookID:    readSessionInput.BookID,
		CreatedAt: time.Now(),
		Durations: []entity.ReadingDuration{
			{
				Date: time.Now(),
				Time: readSessionInput.Durations.Time,
			},
		},
	}

	err = h.readSessionUseCase.CreatReadingSession(c, readSession)

	c.Status(http.StatusCreated)
}
