package db

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Message struct {
	Id         uuid.UUID `json:"id"`
	Text       string    `json:"text"`
	CategoryId uuid.UUID `json:"category_id"`
	PostedAt   time.Time `json:"posted_at"`
	AuthorId   uuid.UUID `json:"author_id"`
}



