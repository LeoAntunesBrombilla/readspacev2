package handler

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"readspacev2/internal/entity"
	"readspacev2/internal/usecase"
	"strconv"
	"time"
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

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body entity.UserInput true "User details for creation"
// @Success 201 {string} string "Created"
// @Failure 400 {object} entity.ErrorEntity
// @Failure 500 {object} entity.ErrorEntity
// @Router /user [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var userInput entity.UserInput

	err := c.ShouldBindJSON(&userInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, entity.ErrorEntity{Code: 400, Message: "Bad Request"})
		return
	}

	user := entity.UserEntity{
		Email:     userInput.Email,
		Username:  userInput.Username,
		Password:  userInput.Password,
		CreatedAt: time.Now().UTC(),
	}

	err = h.userUseCase.CreateUser(&user)
	//TODO melhorar tratamento de erro, caso o cliente envie um email ou username ja existente
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.ErrorEntity{Code: 500, Message: "Error creating user"})
		return
	}

	c.Status(http.StatusCreated)
}

// ListAllUsers godoc
// @Summary List all users
// @Description Retrieve a list of all users in the database
// @Tags users
// @Produce  json
// @Success 200 {array} entity.UserEntity
// @Failure 500 {object} entity.ErrorEntity
// @Router /user [get]
func (h *UserHandler) ListAllUsers(c *gin.Context) {

	users, err := h.userUseCase.ListAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.ErrorEntity{Code: 500, Message: "Error returning list of users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// DeleteUserById godoc
// @Summary Delete a user by ID
// @Description Delete the user identified by the given ID
// @Tags users
// @Produce  json
// @Param id query int64 true "User ID to delete"
// @Success 200 {object} string "User deleted with success"
// @Failure 400 {object} entity.ErrorEntity
// @Failure 500 {object} entity.ErrorEntity
// @Router /user [delete]
func (h *UserHandler) DeleteUserById(c *gin.Context) {
	var id int64

	idStr := c.DefaultQuery("id", "0")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, entity.ErrorEntity{Code: 400, Message: "Invalid ID query parameter"})
		return
	}

	err = h.userUseCase.DeleteUserById(&id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.ErrorEntity{Code: 500, Message: "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "User deleted with success"})
	return
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update the user identified by the given ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id query int64 true "User ID to update"
// @Param details body entity.UserUpdateDetails true "Details to update"
// @Success 200 {object} string "User updated with success"
// @Failure 400 {object} entity.ErrorEntity
// @Failure 500 {object} entity.ErrorEntity
// @Router /user [patch]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var id int64

	idStr := c.DefaultQuery("id", "0")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, entity.ErrorEntity{Code: 400, Message: "Invalid ID query parameter"})
		return
	}

	var userUpdateDetails entity.UserUpdateDetails

	if err = c.ShouldBindJSON(&userUpdateDetails); err != nil {
		c.JSON(http.StatusBadRequest, entity.ErrorEntity{Code: 400, Message: "Invalid payload"})
		return
	}

	err = h.userUseCase.UpdateUser(&id, &userUpdateDetails)

	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.ErrorEntity{Code: 500, Message: "Error updating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "User updated with success"})
}

// UpdateUserPassword godoc
// @Summary Update user password
// @Description Update the password of the user identified by the given ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id query int64 true "User ID to update password"
// @Param details body entity.UserUpdatePassword true "New password details"
// @Success 200 {object} string "User password updated with success"
// @Failure 400 {object} entity.ErrorEntity
// @Failure 500 {object} entity.ErrorEntity
// @Router /user/password [patch]
func (h *UserHandler) UpdateUserPassword(c *gin.Context) {

	idStr := c.DefaultQuery("id", "0")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, entity.ErrorEntity{Code: 400, Message: "Invalid ID query parameter"})
		return
	}

	var userUpdatePassword entity.UserUpdatePassword

	if err := c.ShouldBindJSON(&userUpdatePassword); err != nil {
		c.JSON(http.StatusBadRequest, entity.ErrorEntity{Code: 400, Message: "Invalid request payload"})
		return
	}

	oldHashedPassword, err := h.userUseCase.FindPasswordById(&id)

	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, entity.ErrorEntity{Code: 404, Message: "User not found"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, entity.ErrorEntity{Code: 500, Message: "Error finding user"})
			return
		}
	}

	if oldHashedPassword == nil {
		c.JSON(http.StatusInternalServerError, entity.ErrorEntity{Code: 500, Message: "Old password not found"})
		return
	}

	if err := h.bcryptWrapper.CompareHashAndPassword([]byte(*oldHashedPassword), []byte(userUpdatePassword.OldPassword)); err != nil {
		c.JSON(http.StatusBadRequest, entity.ErrorEntity{Code: 400, Message: "Invalid old password"})
		return
	}

	newHashedPassword, hashErr := h.bcryptWrapper.GenerateFromPassword([]byte(userUpdatePassword.NewPassword), bcrypt.DefaultCost)

	if hashErr != nil {
		c.JSON(http.StatusInternalServerError, entity.ErrorEntity{Code: 500, Message: "Error hashing new password"})
		return
	}

	if updateErr := h.userUseCase.UpdateUserPassword(&id, string(newHashedPassword)); updateErr != nil {
		c.JSON(http.StatusInternalServerError, entity.ErrorEntity{Code: 500, Message: "Error updating user password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "User password updated with success"})
}
