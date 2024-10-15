package http

import (
	"chat_room/room"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handlers(u room.UseCase, r *gin.Engine) {
	rr := r.Group("/room")
	{
		rr.POST("/create", createRoom(u))
    rr.POST("/member", addUser(u))
	}
}

type createRoomInput struct {
	Name string `json:"name" binding:"required"`
}

func createRoom(u room.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		i := createRoomInput{}
		if err := c.ShouldBind(&i); err != nil {
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
      "id": id,
			"success": true,
		})
	}
}

type addUserInput struct {
	RoomID string `json:"room_id" binding:"required"`
	UserID string `json:"user_id" binding:"required"`
}

func addUser(u room.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		i := addUserInput{}
		if err := c.ShouldBind(&i); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"success": false,
				"message": "wrong payload",
			})
			return
		}
		err := u.AddUser(i.RoomID, i.UserID)
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
