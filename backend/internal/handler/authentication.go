package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"readspacev2/internal/usecase"
)

type AuthenticationHandler struct {
	authUseCase usecase.AuthenticationUseCase
}

func NewAuthenticationHandler(authUseCase usecase.AuthenticationUseCase) *AuthenticationHandler {
	return &AuthenticationHandler{
		authUseCase: authUseCase,
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

	c.JSON(http.StatusOK, gin.H{"token": token})
}
