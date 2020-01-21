package db

import (
	"database/sql"
	uuid "github.com/satori/go.uuid"
)

type Category struct {
	Id       uuid.UUID      `json:"id"`
	Name     string         `json:"name"`
	ParentId sql.NullString `json:"parent_id,omitempty"`
}


