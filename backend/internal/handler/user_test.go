package handler

import (
	"bytes"
	"encoding/json"
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

func (m *mockUserUseCase) ListAllUsers() ([]*entity.UserEntity, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockUserUseCase) DeleteUserById(id *int64) error {
	//TODO implement me
	panic("implement me")
}

func (m *mockUserUseCase) FindByUserName(username string) (*entity.UserEntity, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockUserUseCase) UpdateUser(id *int64, user *entity.UserUpdateDetails) error {
	//TODO implement me
	panic("implement me")
}

func (m *mockUserUseCase) UpdateUserPassword(id *int64, password string) error {
	//TODO implement me
	panic("implement me")
}

func (m *mockUserUseCase) FindPasswordById(id *int64) (*string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockUserUseCase) CreateUser(user *entity.UserEntity) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestCreateUser(t *testing.T) {
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
