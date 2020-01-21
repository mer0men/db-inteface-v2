package db

import (
	"database/sql"
)

const (
	createCategoryTableQuery = `
		CREATE TABLE IF NOT EXISTS categories (
			"id" 					uuid NOT NULL,
			"name" 				varchar(255) NOT NULL,
			"parent_id" 	uuid NOT NULL,
		CONSTRAINT "categories_pk" PRIMARY KEY ("id")
		) WITH (
			OIDS=FALSE
		);

		ALTER TABLE "categories" ADD CONSTRAINT "categories_fk0" FOREIGN KEY ("parent_id") REFERENCES "categories"("id");
	`

	insertCategoryQuery = `
		INSERT INTO categories ("id", "name", "parent_id")
		VALUES($1, $2, $3);
	`

	dropCategoryTableQuery = `
		DROP TABLE IF EXISTS categories;
	`

	getCategoryQuery = `
		SELECT * FROM categories AS c WHERE c.id = $1
	`

	getCategoriesQuery = `
		SELECT * FROM categories ORDER BY categories.name
	`

	getSubcategoriesQuery = `
		SELECT * FROM categories AS c 
		WHERE c.parent_id = $1 
		ORDER BY c.name
	`
)

func CreateCategoryTable(conn *sql.DB) error {
	bdTx, err := conn.Begin()
	if err != nil {
		return err
	}

	_, err = bdTx.Exec(createCategoryTableQuery)
	if err != nil {
		return err
	}

	err = bdTx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func InsertCategory(conn *sql.DB, row *Category) error {
	bdTx, err := conn.Begin()
	if err != nil {
		return err
	}

	_, err = bdTx.Exec(insertCategoryQuery, row.Id, row.Name, row.ParentId)
	if err != nil {
		return err
	}

	err = bdTx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func DropCategoryTable(conn *sql.DB) error {
	bdTx, err := conn.Begin()
	if err != nil {
		return err
	}

	_, err = bdTx.Exec(dropCategoryTableQuery)
	if err != nil {
		return err
	}

	err = bdTx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func GetCategory(conn *sql.DB, id string) (Category, error) {
	row := conn.QueryRow(getCategoryQuery, id)

	ct := Category{}

	if err := row.Scan(&ct.Id, &ct.Name, &ct.ParentId); err != nil {
		return ct, err
	}

	return ct, nil
}

func GetCategories(conn *sql.DB) ([]Category, error) {
	rows, err := conn.Query(getCategoriesQuery)
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

func GetSubcategories(conn *sql.DB, id string) ([]Category, error) {
	rows, err := conn.Query(getSubcategoriesQuery, id)

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
