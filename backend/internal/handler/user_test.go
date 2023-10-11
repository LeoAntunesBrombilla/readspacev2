package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"readspacev2/internal/entity"
	"testing"
)

type mockUserUseCase struct {
	mock.Mock
}

type MockBcryptWrapper struct {
	CompareHashAndPasswordFunc func(hashedPassword, password []byte) error
	GenerateFromPasswordFunc   func(password []byte, cost int) ([]byte, error)
}

func (mbw MockBcryptWrapper) CompareHashAndPassword(hashedPassword, password []byte) error {
	return mbw.CompareHashAndPasswordFunc(hashedPassword, password)
}

func (mbw MockBcryptWrapper) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	return mbw.GenerateFromPasswordFunc(password, cost)
}

func (m *mockUserUseCase) CreateUser(user *entity.UserEntity) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestUserHandler_CreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := new(mockUserUseCase)
	mockUseCase.On("CreateUser", mock.AnythingOfType("*entity.UserEntity")).Return(nil)

	handler := NewUserHandler(mockUseCase, MockBcryptWrapper{})

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

	handler := NewUserHandler(mockUseCase, MockBcryptWrapper{})

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

	handler := NewUserHandler(mockUseCase, MockBcryptWrapper{})
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

	handler := NewUserHandler(mockUseCase, MockBcryptWrapper{})

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

	handler := NewUserHandler(mockUseCase, MockBcryptWrapper{})

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

	handler := NewUserHandler(mockUseCase, MockBcryptWrapper{})

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

	handler := NewUserHandler(mockUseCase, MockBcryptWrapper{})

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

	handler := NewUserHandler(mockUseCase, MockBcryptWrapper{})

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

	handler := NewUserHandler(mockUseCase, MockBcryptWrapper{})

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

func TestUserHandler_UpdateUser_Bad_Id(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := new(mockUserUseCase)
	mockUseCase.On("UpdateUser", mock.AnythingOfType("*int64"), mock.AnythingOfType("*entity.UserUpdateDetails")).Return(nil)

	handler := NewUserHandler(mockUseCase, MockBcryptWrapper{})

	r := gin.Default()

	r.PATCH("/user", handler.UpdateUser)

	user := []byte(`{bad json}`)

	payload, err := json.Marshal(user)

	assert.NoError(t, err)

	req, err := http.NewRequest("PATCH", "/user?id=", bytes.NewBuffer(payload))

	assert.NoError(t, err)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUserHandler_UpdateUser_Bad_Request(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := new(mockUserUseCase)
	mockUseCase.On("UpdateUser", mock.AnythingOfType("*int64"), mock.AnythingOfType("*entity.UserUpdateDetails")).Return(nil)

	handler := NewUserHandler(mockUseCase, MockBcryptWrapper{})

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

	handler := NewUserHandler(mockUseCase, MockBcryptWrapper{})

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
	args := m.Called(id, password)
	return args.Error(0)
}

func (m *mockUserUseCase) FindPasswordById(id *int64) (*string, error) {
	args := m.Called(id)

	return args.Get(0).(*string), args.Error(1)
}

func (m *mockUserUseCase) FindByUserName(username string) (*entity.UserEntity, error) {
	args := m.Called(username)

	return args.Get(0).(*entity.UserEntity), args.Error(1)
}

func TestUserHandler_UpdateUserPassword(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockBcrypt := MockBcryptWrapper{
		CompareHashAndPasswordFunc: func(hashedPassword, password []byte) error {
			return nil // mimic success
		},
		GenerateFromPasswordFunc: func(password []byte, cost int) ([]byte, error) {
			return []byte("hashed_password"), nil // mimic success
		},
	}

	mockUseCase := new(mockUserUseCase)
	mockUseCase.On("UpdateUserPassword", mock.AnythingOfType("*int64"), mock.AnythingOfType("string")).Return(nil)
	oldHashedPassword, _ := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
	oldHashedPasswordStr := string(oldHashedPassword)
	mockUseCase.On("FindPasswordById", mock.AnythingOfType("*int64")).Return(&oldHashedPasswordStr, nil)

	handler := NewUserHandler(mockUseCase, mockBcrypt)

	r := gin.Default()
	r.PATCH("/user/password", handler.UpdateUserPassword)

	user := entity.UserUpdatePassword{
		OldPassword: "test",
		NewPassword: "testWithBiggerPass2",
	}
	payload, err := json.Marshal(user)
	assert.NoError(t, err)

	req, err := http.NewRequest("PATCH", "/user/password?id=1", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserHandler_UpdateUserPassword_User_Not_Found(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := new(mockUserUseCase)
	mockUseCase.On("UpdateUserPassword", mock.AnythingOfType("*int64"), mock.AnythingOfType("string")).Return(nil)
	oldPassMock := ""
	mockUseCase.On("FindPasswordById", mock.AnythingOfType("*int64")).Return(&oldPassMock, errors.New("user not found"))

	handler := NewUserHandler(mockUseCase, MockBcryptWrapper{})

	r := gin.Default()

	r.PATCH("/user/password", handler.UpdateUserPassword)

	user := entity.UserUpdatePassword{
		OldPassword: "test",
		NewPassword: "testWithBiggerPass2",
	}

	payload, err := json.Marshal(user)

	assert.NoError(t, err)

	req, err := http.NewRequest("PATCH", "/user/password?id=100", bytes.NewBuffer(payload))

	assert.NoError(t, err)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestUserHandler_UpdateUserPassword_Internal_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := new(mockUserUseCase)
	mockUseCase.On("UpdateUserPassword", mock.AnythingOfType("*int64"), mock.AnythingOfType("string")).Return(nil)
	oldPassMock := ""
	mockUseCase.On("FindPasswordById", mock.AnythingOfType("*int64")).Return(&oldPassMock, errors.New("some error"))

	handler := NewUserHandler(mockUseCase, MockBcryptWrapper{})

	r := gin.Default()

	r.PATCH("/user/password", handler.UpdateUserPassword)

	user := entity.UserUpdatePassword{
		OldPassword: "test",
		NewPassword: "testWithBiggerPass2",
	}

	payload, err := json.Marshal(user)

	assert.NoError(t, err)

	req, err := http.NewRequest("PATCH", "/user/password?id=1", bytes.NewBuffer(payload))

	assert.NoError(t, err)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestUserHandler_UpdateUser_Password_Pass_Not_Found(t *testing.T) {
	gin.SetMode(gin.TestMode)

	var nilString *string = nil
	mockUseCase := new(mockUserUseCase)
	mockUseCase.On("UpdateUserPassword", mock.AnythingOfType("*int64"), mock.AnythingOfType("string")).Return(nil)
	mockUseCase.On("FindPasswordById", mock.AnythingOfType("*int64")).Return(nilString, nil)

	handler := NewUserHandler(mockUseCase, MockBcryptWrapper{})

	r := gin.Default()

	r.PATCH("/user/password", handler.UpdateUserPassword)

	user := entity.UserUpdatePassword{
		OldPassword: "test",
		NewPassword: "testWithBiggerPass2",
	}

	payload, err := json.Marshal(user)

	assert.NoError(t, err)

	req, err := http.NewRequest("PATCH", "/user/password?id=1", bytes.NewBuffer(payload))

	assert.NoError(t, err)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestUserHandler_UpdateUserPassword_Invalid_ID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := new(mockUserUseCase)
	mockUseCase.On("UpdateUserPassword", mock.AnythingOfType("*int64"), mock.AnythingOfType("string")).Return(nil)

	handler := NewUserHandler(mockUseCase, MockBcryptWrapper{})

	r := gin.Default()

	r.PATCH("/user/password", handler.UpdateUserPassword)

	user := entity.UserUpdatePassword{
		OldPassword: "test",
		NewPassword: "testWithBiggerPass2",
	}

	payload, err := json.Marshal(user)

	assert.NoError(t, err)

	req, err := http.NewRequest("PATCH", "/user/password?id=", bytes.NewBuffer(payload))

	assert.NoError(t, err)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUserHandler_UpdateUserPassword_Invalid_Payload(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := new(mockUserUseCase)
	mockUseCase.On("UpdateUserPassword", mock.AnythingOfType("*int64"), mock.AnythingOfType("string")).Return(nil)

	handler := NewUserHandler(mockUseCase, MockBcryptWrapper{})

	r := gin.Default()

	r.PATCH("/user/password", handler.UpdateUserPassword)

	payload, err := json.Marshal([]byte(`{bad json}`))

	assert.NoError(t, err)

	req, err := http.NewRequest("PATCH", "/user/password?id=1", bytes.NewBuffer(payload))

	assert.NoError(t, err)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUserHandler_UpdateUserPassword_Diff_Pass(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockBcrypt := MockBcryptWrapper{
		CompareHashAndPasswordFunc: func(hashedPassword, password []byte) error {
			return errors.New("invalid password")
		},
		GenerateFromPasswordFunc: func(password []byte, cost int) ([]byte, error) {
			return []byte("somehash"), nil
		},
	}

	mockUseCase := new(mockUserUseCase)
	mockUseCase.On("UpdateUserPassword", mock.AnythingOfType("*int64"), mock.AnythingOfType("string")).Return(nil)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
	hashedPasswordStr := string(hashedPassword)
	mockUseCase.On("FindPasswordById", mock.AnythingOfType("*int64")).Return(&hashedPasswordStr, nil)

	handler := NewUserHandler(mockUseCase, mockBcrypt)

	r := gin.Default()

	r.PATCH("/user/password", handler.UpdateUserPassword)

	user := entity.UserUpdatePassword{
		OldPassword: "fakeDiffPass",
		NewPassword: "testWithBiggerPass2",
	}

	payload, err := json.Marshal(user)

	assert.NoError(t, err)

	req, err := http.NewRequest("PATCH", "/user/password?id=1", bytes.NewBuffer(payload))

	assert.NoError(t, err)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUserHandler_UpdateUserPassword_Generate_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockBcrypt := MockBcryptWrapper{
		CompareHashAndPasswordFunc: func(hashedPassword, password []byte) error {
			return nil
		},
		GenerateFromPasswordFunc: func(password []byte, cost int) ([]byte, error) {
			return nil, errors.New("some error")
		},
	}

	mockUseCase := new(mockUserUseCase)
	mockUseCase.On("UpdateUserPassword", mock.AnythingOfType("*int64"), mock.AnythingOfType("string")).Return(nil)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
	hashedPasswordStr := string(hashedPassword)
	mockUseCase.On("FindPasswordById", mock.AnythingOfType("*int64")).Return(&hashedPasswordStr, nil)

	handler := NewUserHandler(mockUseCase, mockBcrypt)

	r := gin.Default()

	r.PATCH("/user/password", handler.UpdateUserPassword)

	user := entity.UserUpdatePassword{
		OldPassword: "fakeDiffPass",
		NewPassword: "testWithBiggerPass2",
	}

	payload, err := json.Marshal(user)

	assert.NoError(t, err)

	req, err := http.NewRequest("PATCH", "/user/password?id=1", bytes.NewBuffer(payload))

	assert.NoError(t, err)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestUserHandler_UpdateUserPassword_Error_Updating_Password(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockBcrypt := MockBcryptWrapper{
		CompareHashAndPasswordFunc: func(hashedPassword, password []byte) error {
			return nil
		},
		GenerateFromPasswordFunc: func(password []byte, cost int) ([]byte, error) {
			return []byte("somehash"), nil
		},
	}

	mockUseCase := new(mockUserUseCase)
	mockUseCase.On("UpdateUserPassword", mock.AnythingOfType("*int64"), mock.AnythingOfType("string")).Return(errors.New("some error"))
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
	hashedPasswordStr := string(hashedPassword)
	mockUseCase.On("FindPasswordById", mock.AnythingOfType("*int64")).Return(&hashedPasswordStr, nil)

	handler := NewUserHandler(mockUseCase, mockBcrypt)

	r := gin.Default()

	r.PATCH("/user/password", handler.UpdateUserPassword)

	user := entity.UserUpdatePassword{
		OldPassword: "fakeDiffPass",
		NewPassword: "testWithBiggerPass2",
	}

	payload, err := json.Marshal(user)

	assert.NoError(t, err)

	req, err := http.NewRequest("PATCH", "/user/password?id=1", bytes.NewBuffer(payload))

	assert.NoError(t, err)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
