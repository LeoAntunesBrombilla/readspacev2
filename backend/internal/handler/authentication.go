package handler

import (
	"github.com/boj/redistore"
	"github.com/gin-gonic/gin"
	"net/http"
	"readspacev2/internal/usecase"
)

type AuthenticationHandler struct {
	authUseCase usecase.AuthenticationUseCase
	store       *redistore.RediStore
}

func NewAuthenticationHandler(authUseCase usecase.AuthenticationUseCase, store *redistore.RediStore) *AuthenticationHandler {
	return &AuthenticationHandler{
		authUseCase: authUseCase,
		store:       store,
	}
}

func (h *AuthenticationHandler) Login(c *gin.Context) {
	var loginDetails struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	token, err := h.authUseCase.Login(loginDetails.Username, loginDetails.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session, _ := h.store.Get(c.Request, "user-session")

	session.Values["username"] = loginDetails.Username
	session.Values["token"] = token

	err = session.Save(c.Request, c.Writer)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *AuthenticationHandler) Logout(c *gin.Context) {
	session, _ := h.store.Get(c.Request, "user-session")
	session.Options.MaxAge = -1
	err := session.Save(c.Request, c.Writer)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}
