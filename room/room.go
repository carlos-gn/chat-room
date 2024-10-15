package room

import (
	"time"
)

type Room struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type RoomMember struct {
	ID     int
	RoomID string
	UserID string
}

type Message struct {
	ID        string
	CreatorID string
	Content   string
	RoomID    string
	CreatedAt time.Time
}

type UseCase interface {
	Create(name string) (string, error)
	AddUser(roomID, userID string) error // This could be a RoomMember
	SendMessage(roomID, userID, message string) error
	DeleteMessage(messageID, userID string) error
}

type RoomRepository interface {
	Create(room Room) error
	AddUser(roomID, userID string) error
	SendMessage(message Message) error
	DeleteMessage(messageID, userID string) error
	Get(roomID string) (*Room, error)
	UserExists(roomID, userID string) (bool, error)
	GetMessageForUser(messageID, userID string) (*Message, error)
}
