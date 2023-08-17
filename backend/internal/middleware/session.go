package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"os"
)

var secret = os.Getenv("SECRET_KEY")

var (
	store = sessions.NewCookieStore([]byte(secret))
)

// TODO create session for admin
func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := store.Get(c.Request, "user-session")
		c.Set("session", session)
		c.Next()
	}
}
