package room

import (
	"chat_room/user"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type RoomService struct {
	rr RoomRepository
	ur user.UserRepository
}

func NewService(rr RoomRepository, ur user.UserRepository) *RoomService {
	return &RoomService{
		rr: rr,
		ur: ur,
	}
}

func (s *RoomService) Create(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("Room name cannot be empty")
	}

	r := Room{
		ID:        uuid.NewString(),
		Name:      name,
		CreatedAt: time.Now().UTC(),
	}

	err := s.rr.Create(r)
	if err != nil {
		return "", err
	}

	return r.ID, nil
}

func (s *RoomService) AddUser(id, userID string) error {
	if id == "" || userID == "" {
		return fmt.Errorf("Missing id or userID")
	}

	// Check if room exists
	r, err := s.rr.Get(id)
	if err != nil {
		return err
	}

	// Check if user exists
	u, err := s.ur.Get(userID)
	if err != nil {
		return err
	}

	// Check if user is already part of the room
	m, err := s.rr.UserExists(id, userID)
	if err != nil {
		return err
	}

	if m {
		return fmt.Errorf("User is already member of the room")
	}

	err = s.rr.AddUser(r.ID, u.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *RoomService) SendMessage(roomID, userID, message string) error {
	// I can really skip room check and user check and see if the user is part of the group.
	// Check if room exists
	r, err := s.rr.Get(roomID)
	if err != nil {
		return err
	}

	// Don't need to check if user exists, we have the middleware for that
	// Check if user is already part of the room
	m, err := s.rr.UserExists(r.ID, userID)
	if err != nil {
		return err
	}

	if !m {
		return fmt.Errorf("Cannot send messages to this room")
	}

	msg := Message{
		ID:        uuid.NewString(),
		RoomID:    r.ID,
		CreatorID: userID,
		Content:   message,
		CreatedAt: time.Now().UTC(),
	}

	err = s.rr.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}

func (s *RoomService) DeleteMessage(messageID, userID string) error {
	m, err := s.rr.GetMessageForUser(messageID, userID)
	if err != nil {
		return err
	}

	err = s.rr.DeleteMessage(m.ID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (s *RoomService) GetMessages(roomID, userID, cursorID string, cursorTime time.Time) ([]Message, string, error) {
	exists, err := s.rr.UserExists(roomID, userID)
	if err != nil {
		return nil, "", err
	}

	if !exists {
		return nil, "", fmt.Errorf("No authorized to check the messages")
	}

	m, cursor, err := s.rr.GetMessages(roomID, cursorID, cursorTime)
	if err != nil {
		return nil, "", err
	}
	return m, cursor, nil
}
