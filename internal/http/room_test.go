package http_test

import (
	"bytes"
	"chat_room/infra"
	"chat_room/room"
	storage_room "chat_room/room/sqlite"
	storage_user "chat_room/user/sqlite"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	. "chat_room/internal/http"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Room API", Ordered, func() {
	var (
		roomService *room.RoomService
		r           *gin.Engine
		db          *sql.DB
	)
	_ = BeforeAll(func() {
		r = gin.Default()
		db, _ = infra.NewTestDB()
		rr := storage_room.New(db)
		ur := storage_user.New(db)
		roomService = room.NewService(rr, ur)
		Handlers(roomService, r)
	})
	_ = AfterAll(func() {
		err := infra.CleanDB(db)
		if err != nil {
			log.Fatal(err)
		}
		db.Close()
	})

	Describe("POST Create room", func() {
		Context("there's not room with that name", func() {
			It("should create the room", func(ctx SpecContext) {
				payload, _ := json.Marshal(map[string]string{"name": "Room1"})
				req, _ := http.NewRequest("POST", "/rooms", bytes.NewBuffer(payload))
				req.Header.Add("Content-Type", "application/json")
				resp := httptest.NewRecorder()
				r.ServeHTTP(resp, req)
				fmt.Println(resp)
				Expect(resp.Code).To(Equal(http.StatusOK))
			})
		})
	})
})
