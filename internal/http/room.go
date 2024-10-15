package http

import (
	"chat_room/room"
	"chat_room/user"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Handlers(u room.UseCase, r *gin.Engine) {
	rr := r.Group("/rooms")
	{
		rr.POST("", createRoom(u))
		rr.POST("/:roomID/members", addUser(u))
		rr.POST("/:roomID/messages", sendMessage(u))
		rr.GET("/:roomID/messages", getMessages(u))
		rr.DELETE("/:roomID/messages/:messageID", deleteMessage(u))
	}
}

type createRoomInput struct {
	Name string `json:"name" binding:"required"`
}

func createRoom(u room.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		i := createRoomInput{}
		if err := c.ShouldBindJSON(&i); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"success": false,
				"message": "wrong payload",
			})
			return
		}
		id, err := u.Create(c.Request.Context(), i.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"success": true,
		})
	}
}

type addUserInput struct {
	UserID string `json:"user_id" binding:"required"`
}

func addUser(u room.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		roomID := c.Param("roomID")
		if roomID == "" {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"success": false,
				"message": "wrong payload",
			})
			return
		}
		i := addUserInput{}
		if err := c.ShouldBindJSON(&i); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"success": false,
				"message": "wrong payload",
			})
			return
		}
		err := u.AddUser(c.Request.Context(), roomID, i.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	}
}

type sendMessageInput struct {
	Message string `json:"message" binding:"required"`
}

func sendMessage(u room.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		roomID := c.Param("roomID")
		if roomID == "" {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"success": false,
				"message": "wrong payload",
			})
			return
		}
		i := sendMessageInput{}
		if err := c.ShouldBindJSON(&i); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"success": false,
				"message": "wrong payload",
			})
			return
		}

		// This can be in a "shared" method
		userContext, _ := c.Get("user")
		user := userContext.(user.User)
		err := u.SendMessage(c.Request.Context(), roomID, user.ID, i.Message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	}
}

func getMessages(u room.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		roomID := c.Param("roomID")

		if roomID == "" {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"success": false,
				"message": "wrong payload",
			})
			return
		}

		cursor := c.Query("cursor")
		userContext, _ := c.Get("user")
		user := userContext.(user.User)
		var cursorTime time.Time
		var cursorID string
		var err error
		if cursor != "" {
			cursorTime, cursorID, err = DecodeCursor(cursor)
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"success": false,
					"message": "wrong cursor",
				})
				return
			}
		}
		messages, cursor, err := u.GetMessages(c.Request.Context(), roomID, user.ID, cursorID, cursorTime)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    messages,
			"cursor":  cursor,
		})
	}
}

func deleteMessage(u room.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		messageID := c.Param("messageID")

		if messageID == "" {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"success": false,
				"message": "wrong payload",
			})
			return
		}

		// This can be in a "shared" method
		userContext, _ := c.Get("user")
		user := userContext.(user.User)
		err := u.DeleteMessage(c.Request.Context(), messageID, user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	}
}
