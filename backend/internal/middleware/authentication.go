package middleware

import (
	"github.com/boj/redistore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthenticationMiddleware(store *redistore.RediStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := store.Get(c.Request, "user-session")
		if err != nil || session.Values["username"] == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.Set("username", session.Values["username"])
		c.Next()
	}
}
