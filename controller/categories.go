package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

type Category struct {
	Id       uuid.UUID      `json:"id"`
	Name     string         `json:"name"`
	ParentId sql.NullString `json:"parent_id,omitempty"`
}

func (u *UserController) GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	queryValues := r.URL.Query()
	id := queryValues.Get("id")

	var row *sql.Row
	if len(id) == 0 || id == "" {
		row = u.db.QueryRow("SELECT * FROM categories AS c WHERE c.name = 'Forum'")
	} else {
		row = u.db.QueryRow("SELECT * FROM categories AS c WHERE c.id = $1", id)
	}

	ct := Category{}
	if err := row.Scan(&ct.Id, &ct.Name, &ct.ParentId); err != nil {
		log.Printf("Cannot scan root category, error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := json.NewEncoder(w).Encode(ct)
	if err != nil {
		log.Printf("Error encoding root category to json, error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (u *UserController) GetSubcategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	queryValues := r.URL.Query()
	id := queryValues.Get("id")

	categories, err := u.getSubcategories(id)
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

func (u *UserController) getSubcategories(id string) ([]Category, error) {
	rows, err := u.db.Query(
		fmt.Sprintf(`
			SELECT * FROM categories AS c 
			WHERE c.parent_id = '%v' 
			ORDER BY c.name`, id,
		),
	)
	if err != nil {
		return make([]Category, 0), err
	}

	var categories []Category
	for rows.Next() {
		ct := Category{}
		err := rows.Scan(&ct.Id, &ct.Name, &ct.ParentId)
		if err != nil {
			return make([]Category, 0), err
		}
		categories = append(categories, ct)
	}

	//log.Println(len(categories))

	if err = rows.Err(); err != nil {
		return make([]Category, 0), err
	}

	if len(categories) == 0 {
		return make([]Category, 0), nil
	}
	return categories, nil
}

func (u *UserController) GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categories, err := u.getCategories()
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

func (u *UserController) getCategories() ([]Category, error) {
	rows, err := u.db.Query(
		"SELECT * FROM categories ORDER BY categories.name",
	)
	if err != nil {
		return make([]Category, 0), err
	}

	var categories []Category
	for rows.Next() {
		ct := Category{}
		err := rows.Scan(&ct.Id, &ct.Name, &ct.ParentId)
		if err != nil {
			return make([]Category, 0), err
		}
		categories = append(categories, ct)
	}

	if err = rows.Err(); err != nil {
		return make([]Category, 0), err
	}

	if len(categories) == 0 {
		return make([]Category, 0), nil
	}
	return categories, nil
}
