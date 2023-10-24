package handler

import (
	"net/http"
	"readspacev2/internal/entity"
	"readspacev2/internal/usecase"

	"github.com/boj/redistore"
	"github.com/gin-gonic/gin"
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

// Login godoc
// @Summary Authenticate a user and obtain a token
// @Description Authenticate using username and password to get a token
// @Tags authentication
// @Accept  json
// @Produce  json
// @Param login body entity.Login true "Credentials for login"
// @Success 200 {object} map[string]string "Successfully authenticated, token returned"
// @Router /login [post]
func (h *AuthenticationHandler) Login(c *gin.Context) {
  var loginDetails entity.Login;

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
