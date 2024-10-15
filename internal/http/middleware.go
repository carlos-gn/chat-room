package http

import (
	"chat_room/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserMiddleware(u user.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		xUserID := c.Request.Header.Get("x-user-id")
		if xUserID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user is not authorized"})
			return
		}
		user, err := u.Get(xUserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Something went wrong")
		}
		c.Set("user", *user)
		c.Next()
	}
}
