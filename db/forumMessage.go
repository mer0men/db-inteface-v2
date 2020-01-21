package db

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type ForumMessage struct {
	Id         uuid.UUID `json:"id"`
	AuthorId   string    `json:"author_id"`
	AuthorName string    `json:"author_name"`
	Text       string    `json:"text"`
	PostedAt   time.Time `json:"posted_at"`
}

