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
	ID        string    `json:"id"`
	CreatorID string    `json:"creator_id"`
	Content   string    `json:"content"`
	RoomID    string    `json:"room_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UseCase interface {
	Create(name string) (string, error)
	AddUser(roomID, userID string) error // This could be a RoomMember
	SendMessage(roomID, userID, message string) error
	DeleteMessage(messageID, userID string) error
	GetMessages(roomID, userID, cursorID string, cursorTime time.Time) ([]Message, string, error)
}

type RoomRepository interface {
	Create(room Room) error
	AddUser(roomID, userID string) error
	SendMessage(message Message) error
	// Ideally we should have a struct with the different pagination options/values
	GetMessages(roomID, cursorID string, cursorTime time.Time) ([]Message, string, error)
	DeleteMessage(messageID, userID string) error
	Get(roomID string) (*Room, error)
	UserExists(roomID, userID string) (bool, error)
	GetMessageForUser(messageID, userID string) (*Message, error)
}
