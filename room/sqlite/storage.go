package storage_room

import (
	"chat_room/internal/http"
	"chat_room/room"
	"database/sql"
	"errors"
	"fmt"
	"time"

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

func (s *Storage) DeleteMessage(id, userID string) error {
	q := "DELETE FROM messages where id = ? and creator_id = ?;"
	// TODO: Logs
	_, err := s.DB.Exec(q, id, userID)
	if err != nil {
		return fmt.Errorf("Something went wrong deleting the mesage")
	}

	return nil
}

func (s *Storage) GetMessageForUser(id, userID string) (*room.Message, error) {
	q := "SELECT id, content, creator_id, room_id, created_at FROM messages WHERE id = ? AND creator_id = ? LIMIT 1;"

	// TODO: Logs
	row := s.DB.QueryRow(q, id, userID)
	var m room.Message
	var err error
	if err = row.Scan(&m.ID, &m.Content, &m.CreatorID, &m.RoomID, &m.CreatedAt); err == sql.ErrNoRows {
		return nil, fmt.Errorf("Message not found")
	}

	return &m, nil
}

func (s *Storage) GetMessages(id, cursorID string, cursorTime time.Time) ([]room.Message, string, error) {
	const limit = 10

	var q string
	var rows *sql.Rows
	var err error

	if cursorID != "" && !cursorTime.IsZero() {
		q = "SELECT id, content, creator_id, room_id, created_at FROM messages WHERE room_id = ? LIMIT ?;"
		rows, err = s.DB.Query(q, id, limit+1)
	} else {
		q = `
        SELECT id, content, creator_id, room_id, created_at FROM messages 
        WHERE room_id = ? AND (messages.created_at, messages.id) > (?, ?) LIMIT ?;`
		rows, err = s.DB.Query(q, id, cursorTime, cursorID, limit+1)
	}

	// TODO: Logs
	if err != nil {
		fmt.Println(err)
		return nil, "", fmt.Errorf("Something went wrong getting the messages")
	}
	defer rows.Close()
	var mm []room.Message
	for rows.Next() {
		var m room.Message
		err = rows.Scan(&m.ID, &m.Content, &m.CreatorID, &m.RoomID, &m.CreatedAt)
		if err != nil {
			return nil, "", err
		}
		mm = append(mm, m)
	}

	err = rows.Err()
	if err != nil {
		return nil, "", err
	}

	existsMore := len(mm) > limit
	cursor := ""
	if existsMore {
		mm = mm[:limit]
		lastTs := mm[limit-1]
		cursor = http.EncodeCursor(lastTs.CreatedAt.UTC(), lastTs.ID)
	}

	return mm, cursor, nil
}
