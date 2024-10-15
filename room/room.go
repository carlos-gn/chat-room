package room

import (
	"context"
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
	Create(ctx context.Context, name string) (string, error)
	AddUser(ctx context.Context, roomID, userID string) error // This could be a RoomMember
	SendMessage(ctx context.Context, roomID, userID, message string) error
	DeleteMessage(ctx context.Context, messageID, userID string) error
	GetMessages(ctx context.Context, roomID, userID, cursorID string, cursorTime time.Time) ([]Message, string, error)
}

type RoomRepository interface {
	Create(ctx context.Context, room Room) error
	AddUser(ctx context.Context, roomID, userID string) error
	SendMessage(ctx context.Context, message Message) error
	// Ideally we should have a struct with the different pagination options/values
	GetMessages(ctx context.Context, roomID, cursorID string, cursorTime time.Time) ([]Message, string, error)
	DeleteMessage(ctx context.Context, messageID, userID string) error
	Get(ctx context.Context, roomID string) (*Room, error)
	UserExists(ctx context.Context, roomID, userID string) (bool, error)
	GetMessageForUser(ctx context.Context, messageID, userID string) (*Message, error)
}
