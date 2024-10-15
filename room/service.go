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

	r := &Room{
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
