package controller

import (
	"database/sql"
	"net/http"
)

type UserController struct {
	db *sql.DB
}

func New(db *sql.DB) *UserController {
	return &UserController{db: db}
}

func (c *UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := c.NewRouter()
	router.ServeHTTP(w, r)
}
