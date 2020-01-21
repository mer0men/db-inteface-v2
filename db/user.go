package db

import uuid "github.com/satori/go.uuid"

type User struct {
	Id   uuid.UUID
	Name string
}
