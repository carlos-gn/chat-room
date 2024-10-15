package storage_user

import (
	"chat_room/user"
	"context"
	"database/sql"
	"fmt"
)

type Storage struct {
	DB *sql.DB
}

func New(db *sql.DB) *Storage {
	return &Storage{
		DB: db,
	}
}

func (s *Storage) Get(ctx context.Context, id string) (*user.User, error) {
	q := "SELECT id, name, created_at FROM users WHERE id = ? LIMIT 1;"

	row := s.DB.QueryRowContext(ctx, q, id)
	var u user.User
	var err error
	if err = row.Scan(&u.ID, &u.Name, &u.CreatedAt); err == sql.ErrNoRows {
		return nil, fmt.Errorf("User not found")
	}

	return &u, nil
}
