package entities

import (
	"time"
)

type Message struct {
	ID       string
	Sender   User
	Content  string
	Room     *Room
	SentAt   time.Time
	EditedAt time.Time
}
