package storage_room

import (
	"chat_room/room"
	"database/sql"
	"errors"
	"fmt"

	"github.com/mattn/go-sqlite3"
)

type Storage struct {
	DB *sql.DB
}

func New(db *sql.DB) *Storage {
	return &Storage{
		DB: db,
	}
}

// TODO: Return human error from repo or from service?

func (s *Storage) Create(r room.Room) error {
	q := "INSERT INTO rooms(id, name, created_at) VALUES (?,?,?);"

	// TODO: Logs
	_, err := s.DB.Exec(q, r.ID, r.Name, r.CreatedAt)
	if err != nil {
		if errors.Is(err.(sqlite3.Error).Code, sqlite3.ErrConstraint) {
			return fmt.Errorf("Already exists a room with that name")
		}
		return fmt.Errorf("Something went wrong inserting the room")
	}

	return nil
}

func (s *Storage) Get(id string) (*room.Room, error) {
	q := "SELECT id, name, created_at FROM rooms WHERE id = ? LIMIT 1;"

	// TODO: Logs
	row := s.DB.QueryRow(q, id)
	var r room.Room
	var err error
	if err = row.Scan(&r.ID, &r.Name, &r.CreatedAt); err == sql.ErrNoRows {
		return nil, fmt.Errorf("Room not found")
	}

	return &r, nil
}

func (s *Storage) AddUser(id string, userID string) error {
	q := "INSERT INTO room_members(room_id, user_id) VALUES (?, ?);"

	// TODO: Logs
	_, err := s.DB.Exec(q, id, userID)
	if err != nil {
		if errors.Is(err.(sqlite3.Error).Code, sqlite3.ErrConstraint) {
			return fmt.Errorf("User is already member of the room")
		}
		return fmt.Errorf("Something went wrong adding the user")
	}

	return nil
}

func (s *Storage) UserExists(id, userID string) (bool, error) {
	q := "SELECT id, room_id, user_id FROM room_members WHERE room_id = ? and user_id = ? LIMIT 1;"

	// TODO: Logs
	row := s.DB.QueryRow(q, id, userID)
	var r room.RoomMember
	err := row.Scan(&r.ID, &r.RoomID, &r.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	if r.ID == 0 {
		return false, nil
	}

	return true, nil
}

func (s *Storage) SendMessage(m room.Message) error {
	q := "INSERT INTO messages(id, room_id, creator_id, content, created_at) VALUES (?,?,?,?,?);"

	// TODO: Logs
	_, err := s.DB.Exec(q, m.ID, m.RoomID, m.CreatorID, m.Content, m.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Something went wrong sending the mesage")
	}

	return nil
}

func (s *Storage) DeleteMessage(m string) error {
	return nil
}
