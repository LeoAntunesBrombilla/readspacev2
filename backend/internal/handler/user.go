package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"readspacev2/internal/entity"
	"readspacev2/internal/usecase"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user entity.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err = h.userUseCase.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *UserHandler) ListAllUsers(c *gin.Context) {

	users, err := h.userUseCase.ListAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error returning list of users"})
		return
	}

	c.JSON(http.StatusOK, users)
}
