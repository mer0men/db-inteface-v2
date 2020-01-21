package db

import "database/sql"

const (
	createUserTableQuery = `
		CREATE TABLE IF NOT EXISTS users (
			"id" 		uuid NOT NULL,
			"name" 	varchar(255) NOT NULL,
		CONSTRAINT "users_pk" PRIMARY KEY ("id")
		) WITH (
			OIDS=FALSE
		);
	`
	insertUserQuery = `
		INSERT INTO users ("id", "name")
		VALUES($1, $2);
	`

	dropUserTableQuery = `
		DROP TABLE IF EXISTS users;
	`
)

func CreateUserTable(conn *sql.DB) error {
	bdTx, err := conn.Begin()
	if err != nil {
		return err
	}

	_, err = bdTx.Exec(createUserTableQuery)
	if err != nil {
		return err
	}

	err = bdTx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func InsertUser(conn *sql.DB, row *User) error {
	bdTx, err := conn.Begin()
	if err != nil {
		return err
	}

	_, err = bdTx.Exec(insertUserQuery, row.Id, row.Name)
	if err != nil {
		return err
	}

	err = bdTx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func DropUserTable(conn *sql.DB) error {
	bdTx, err := conn.Begin()
	if err != nil {
		return err
	}

	_, err = bdTx.Exec(dropUserTableQuery)
	if err != nil {
		return err
	}

	err = bdTx.Commit()
	if err != nil {
		return err
	}

	return nil
}
