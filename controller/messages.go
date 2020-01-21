package controller

import (
	"encoding/json"
	"github.com/meromen/db-inteface-v2/db"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

// Message represents message instance from DB

func (u *UserController) GetMessages(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	id := queryValues.Get("id")

	messages, err := db.GetMessages(u.db, id)
	if err != nil {
		log.Printf("Cannot extract messages from database: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(messages)
	if err != nil {
		log.Printf("Error encoding messages to json, error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (u *UserController) AddMessage(w http.ResponseWriter, r *http.Request) {
	var message db.Message
	message.Text = r.FormValue("text")
	message.CategoryId, _ = uuid.FromString(r.FormValue("category_id"))

	message.Id, _ = uuid.NewV4()

	err := db.InsertMessage(u.db, &message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Cannot execute message, error: %v", err)
		return
	}

	err = json.NewEncoder(w).Encode(message)
	if err != nil {
		log.Printf("Error encoding message to json, error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
