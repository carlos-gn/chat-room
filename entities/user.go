package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

func NewUser(name string) *User {
	return &User{
		ID:        uuid.NewString(),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
