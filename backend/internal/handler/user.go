package handler

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"readspacev2/internal/entity"
	"readspacev2/internal/usecase"
	"strconv"
)

type BcryptWrapper interface {
	GenerateFromPassword(password []byte, cost int) ([]byte, error)
	CompareHashAndPassword(hashedPassword, password []byte) error
}

type UserHandler struct {
	userUseCase   usecase.UserUseCaseInterface
	bcryptWrapper BcryptWrapper
}

func NewUserHandler(userUseCase usecase.UserUseCaseInterface, bcrypt BcryptWrapper) *UserHandler {
	return &UserHandler{
		userUseCase:   userUseCase,
		bcryptWrapper: bcrypt,
	}
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

func (h *UserHandler) UpdateUserPassword(c *gin.Context) {

	idStr := c.DefaultQuery("id", "0")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID query parameter"})
		return
	}

	var userUpdatePassword entity.UserUpdatePassword

	if err := c.ShouldBindJSON(&userUpdatePassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	oldHashedPassword, err := h.userUseCase.FindPasswordById(&id)

	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding user"})
			return
		}
	}

	if oldHashedPassword == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Old password not found"})
		return
	}

	if err := h.bcryptWrapper.CompareHashAndPassword([]byte(*oldHashedPassword), []byte(userUpdatePassword.OldPassword)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid old password"})
		return
	}

	newHashedPassword, hashErr := h.bcryptWrapper.GenerateFromPassword([]byte(userUpdatePassword.NewPassword), bcrypt.DefaultCost)

	if hashErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing new password"})
		return
	}

	if updateErr := h.userUseCase.UpdateUserPassword(&id, string(newHashedPassword)); updateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "User password updated with success"})
}
