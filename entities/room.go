package entities

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time

	Users    []*User
	Messages []Message
}

// Creates a new Room
func NewRoom(name string) *Room {
	return &Room{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Users:     []*User{},
	}
}

// Add the user to the Room if it's not there.
func (r *Room) AddUser(user *User) {
	found := false
	for i := range r.Users {
		if r.Users[i].ID == user.ID {
			found = true
		}
	}

	if found {
		return
	}
	r.Users = append(r.Users, user)
}

// Get latest <limit> messages in a room
func (r *Room) GetLatestMessages(limit int) []Message {
	var messages []Message
	for i, m := range r.Messages {
		if i == limit {
			break
		}
		messages = append(messages, m)
	}
	return messages
}

type RoomInfoResponse struct {
	UsersCount    int
	MessagesCount int
}

// Get basic info about the Room
func (r *Room) GetRoomInfo() RoomInfoResponse {
	return RoomInfoResponse{
		UsersCount:    len(r.Users),
		MessagesCount: len(r.Messages),
	}
}

// Add a message to the Room but only if the user is part of it
func (r *Room) AddMessage(msg string, sender User) error {
	// Validate user is part of the room
	found := false
	for i := range r.Users {
		if r.Users[i].ID == sender.ID {
			found = true
		}
	}

	if !found {
		return fmt.Errorf("Can't send message user not found in chat room")
	}

	m := Message{
		ID:       uuid.NewString(),
		Sender:   sender,
		Room:     r,
		Content:  msg,
		SentAt:   time.Now(),
		EditedAt: time.Now(),
	}
	r.Messages = append(r.Messages, m)
	return nil
}
