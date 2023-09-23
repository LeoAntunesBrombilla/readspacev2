package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"readspacev2/internal/entity"
	"readspacev2/internal/usecase"
	"strconv"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user entity.UserEntity

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

func (h *UserHandler) DeleteUserById(c *gin.Context) {
	var id int64

	idStr := c.DefaultQuery("id", "0")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID query parameter"})
		return
	}

	err = h.userUseCase.DeleteUserById(&id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user by Id"})
	}

	c.JSON(http.StatusOK, gin.H{"success": "UserEntity deleted with success"})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var id int64

	idStr := c.DefaultQuery("id", "0")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID query parameter"})
		return
	}

	var userUpdateDetails entity.UserUpdateDetails

	if err := c.ShouldBindJSON(&userUpdateDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err = h.userUseCase.UpdateUser(&id, &userUpdateDetails)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "UserEntity updated with success"})
}
