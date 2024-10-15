package room

import (
	"time"
)

type Room struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type UseCase interface {
	Create(name string) (string, error)
	AddUser(id, userId string) error
}

type RoomRepository interface {
	Create(r *Room) error
	Get(id string) (*Room, error)
	AddUser(id, userID string) error
  UserExists(id, userID string) (bool, error)
}

type RoomMember struct {
  ID int
  RoomID string
  UserID string
}


type Message struct {
	ID       string
	SenderID string
	Content  string
	RoomID   string
	SentAt   time.Time
	EditedAt time.Time
}
