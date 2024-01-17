package handler

import (
	"github.com/LeoAntunesBrombilla/readspacev2/internal/entity"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/usecase"
	"net/http"

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
	var loginDetails entity.Login

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

// Logout godoc
// @Summary Log out a user and invalidate the session
// @Description Invalidate the user's current session, effectively logging them out
// @Tags authentication
// @Produce  json
// @Success 200 {object} map[string]string "Successfully logged out"
// @Failure 500 {object} entity.ErrorEntity
// @Router /logout [post]
func (h *AuthenticationHandler) Logout(c *gin.Context) {
	session, _ := h.store.Get(c.Request, "user-session")
	session.Options.MaxAge = -1
	err := session.Save(c.Request, c.Writer)

	//TODO fix this
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.ErrorEntity{Code: 500, Message: "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}
