package http

import (
	"chat_room/room"
	"chat_room/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handlers(u room.UseCase, r *gin.Engine) {
	rr := r.Group("/rooms")
	{
		rr.POST("", createRoom(u))
		rr.POST("/:roomID/members", addUser(u))
		rr.POST("/:roomID/messages", sendMessage(u))
		// We can pass the room id to improve the DB search here.
		rr.DELETE("/messages/:messageID", deleteMessage(u))
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
		id, err := u.Create(i.Name)
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
		err := u.AddUser(roomID, i.UserID)
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
		err := u.SendMessage(roomID, user.ID, i.Message)
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
		err := u.DeleteMessage(messageID, user.ID)
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
