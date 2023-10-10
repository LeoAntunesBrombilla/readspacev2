package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"readspacev2/internal/entity"
	"testing"
)

type mockUserUseCase struct {
	mock.Mock
}

func (m *mockUserUseCase) CreateUser(user *entity.UserEntity) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestUserHandler_CreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := new(mockUserUseCase)
	mockUseCase.On("CreateUser", mock.AnythingOfType("*entity.UserEntity")).Return(nil)

	handler := NewUserHandler(mockUseCase)

	r := gin.Default()
	r.POST("/user", handler.CreateUser)

	user := entity.UserEntity{
		Username: "test",
		Password: "test",
		Email:    "test@gmail",
	}

	payload, err := json.Marshal(user)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestUserHandler_CreateUser_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := new(mockUserUseCase)
	mockUseCase.On("CreateUser", mock.AnythingOfType("*entity.UserEntity")).Return(nil)

	handler := NewUserHandler(mockUseCase)

	r := gin.Default()
	r.POST("/user", handler.CreateUser)

	payload := []byte(`{bad json}`)

	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUserHandler_CreateUser_UseCaseError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := new(mockUserUseCase)
	mockUseCase.On("CreateUser", mock.AnythingOfType("*entity.UserEntity")).Return(errors.New("some error"))

	handler := NewUserHandler(mockUseCase)
	r := gin.Default()
	r.POST("/user", handler.CreateUser)

	user := entity.UserEntity{
		Username: "test",
		Password: "test",
		Email:    "test@gmail",
	}

	payload, err := json.Marshal(user)

	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(payload))
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func (m *mockUserUseCase) ListAllUsers() ([]*entity.UserEntity, error) {
	args := m.Called()

	var list []*entity.UserEntity
	if args.Get(0) != nil {
		list = args.Get(0).([]*entity.UserEntity)
	}

	err := args.Error(1)

	return list, err
}

func TestUserHandler_ListAllUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := new(mockUserUseCase)

	users := []*entity.UserEntity{
		{Username: "user1", Password: "pass1", Email: "email1"},
		{Username: "user2", Password: "pass2", Email: "email2"},
	}

	mockUseCase.On("ListAllUsers").Return(users, nil)

	handler := NewUserHandler(mockUseCase)

	r := gin.Default()
	r.GET("/user", handler.ListAllUsers)

	req, err := http.NewRequest("GET", "/user", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserHandler_ListAllUsers_InternalServerError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := new(mockUserUseCase)
	mockUseCase.On("ListAllUsers").Return(nil, errors.New("some error"))

	handler := NewUserHandler(mockUseCase)

	r := gin.Default()
	r.GET("/user", handler.ListAllUsers)

	req, err := http.NewRequest("GET", "/user", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func (m *mockUserUseCase) DeleteUserById(id *int64) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestUserHandler_DeleteUserById(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := new(mockUserUseCase)
	mockUseCase.On("DeleteUserById", mock.AnythingOfType("*int64")).Return(nil)

	handler := NewUserHandler(mockUseCase)

	r := gin.Default()

	r.DELETE("/user", handler.DeleteUserById)

	req, err := http.NewRequest("DELETE", "/user?id=1", nil)

	assert.NoError(t, err)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserHandler_DeleteUserById_Invalid_Query_Param(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := new(mockUserUseCase)
	mockUseCase.On("DeleteUserById", mock.AnythingOfType("*int64")).Return(nil)

	handler := NewUserHandler(mockUseCase)

	r := gin.Default()
	r.DELETE("/user", handler.DeleteUserById)

	req, err := http.NewRequest("DELETE", "/user", nil)

	assert.NoError(t, err)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUserHandler_DeleteUserById_Internal_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := new(mockUserUseCase)
	mockUseCase.On("DeleteUserById", mock.AnythingOfType("*int64")).Return(errors.New("Some error"))

	handler := NewUserHandler(mockUseCase)

	r := gin.Default()
	r.DELETE("/user", handler.DeleteUserById)

	req, err := http.NewRequest("DELETE", "/user?id=1", nil)

	assert.NoError(t, err)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func (m *mockUserUseCase) UpdateUser(id *int64, user *entity.UserUpdateDetails) error {
	args := m.Called(id, user)
	return args.Error(0)
}

func TestUserHandler_UpdateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := new(mockUserUseCase)
	mockUseCase.On("UpdateUser", mock.AnythingOfType("*int64"), mock.AnythingOfType("*entity.UserUpdateDetails")).Return(nil)

	handler := NewUserHandler(mockUseCase)

	r := gin.Default()

	r.PATCH("/user", handler.UpdateUser)

	user := entity.UserEntity{
		Username: "testUpdated",
		Password: "test",
		Email:    "testeupdated@gmail",
	}

	payload, err := json.Marshal(user)

	assert.NoError(t, err)

	req, err := http.NewRequest("PATCH", "/user?id=1", bytes.NewBuffer(payload))

	assert.NoError(t, err)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserHandler_UpdateUser_Bad_Request(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := new(mockUserUseCase)
	mockUseCase.On("UpdateUser", mock.AnythingOfType("*int64"), mock.AnythingOfType("*entity.UserUpdateDetails")).Return(nil)

	handler := NewUserHandler(mockUseCase)

	r := gin.Default()

	r.PATCH("/user", handler.UpdateUser)

	user := []byte(`{bad json}`)

	payload, err := json.Marshal(user)

	assert.NoError(t, err)

	req, err := http.NewRequest("PATCH", "/user?id=1", bytes.NewBuffer(payload))

	assert.NoError(t, err)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUserHandler_UpdateUser_Internal_Server_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := new(mockUserUseCase)
	mockUseCase.On("UpdateUser", mock.AnythingOfType("*int64"), mock.AnythingOfType("*entity.UserUpdateDetails")).Return(errors.New("Some error"))

	handler := NewUserHandler(mockUseCase)

	r := gin.Default()

	r.PATCH("/user", handler.UpdateUser)

	user := entity.UserEntity{
		Username: "testUpdated",
		Password: "test",
		Email:    "testeupdated@gmail",
	}

	payload, err := json.Marshal(user)

	assert.NoError(t, err)

	req, err := http.NewRequest("PATCH", "/user?id=1", bytes.NewBuffer(payload))

	assert.NoError(t, err)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func (m *mockUserUseCase) UpdateUserPassword(id *int64, password string) error {
	//TODO implement me
	panic("implement me")
}

func (m *mockUserUseCase) FindPasswordById(id *int64) (*string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockUserUseCase) FindByUserName(username string) (*entity.UserEntity, error) {
	//TODO implement me
	panic("implement me")
}
