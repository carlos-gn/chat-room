package main

import (
	"chat_room/infra"
	"chat_room/internal/http"
	"chat_room/room"
	storage_room "chat_room/room/sqlite"
	storage_user "chat_room/user/sqlite"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := infra.NewDB()
	if err != nil {
		panic(err)
	}

	mr := gin.Default()

	// Room
	rr := storage_room.New(db)
	ur := storage_user.New(db)
	rs := room.NewService(rr, ur)
	http.Handlers(rs, mr)

	mr.Run(":3000")
}
