package controller

import (
	"encoding/json"
	"github.com/meromen/db-inteface-v2/db"
	"log"
	"net/http"
)

func (u *UserController) GetCategory(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	id := queryValues.Get("id")

	ct, err := db.GetCategory(u.db, id)
	if err != nil {
		log.Printf("Cannot scan root category, error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(ct)
	if err != nil {
		log.Printf("Error encoding root category to json, error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (u *UserController) GetSubcategories(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	id := queryValues.Get("id")

	categories, err := db.GetSubcategories(u.db, id)
	if err != nil {
		log.Printf("Cannot extract subcategories from database: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(categories)
	if err != nil {
		log.Printf("Error encoding subcategories to json, error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (u *UserController) GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categories, err := db.GetCategories(u.db)
	if err != nil {
		log.Printf("Cannot extract categories from database: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(categories)
	if err != nil {
		log.Printf("Error encoding categories to json, error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
