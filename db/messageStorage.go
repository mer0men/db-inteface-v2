package db

import (
	"database/sql"
	"time"
)

const (
	createMessageTableQuery = `
		CREATE TABLE IF NOT EXISTS messages (
			"id" uuid NOT NULL,
			"text" TEXT NOT NULL,
			"category_id" uuid NOT NULL,
			"posted_at" TIME NOT NULL,
			"author_id" uuid NOT NULL,
		CONSTRAINT "messages_pk" PRIMARY KEY ("id")
		) WITH (
			OIDS=FALSE
		);

		ALTER TABLE "messages" ADD CONSTRAINT "messages_fk0" FOREIGN KEY ("category_id") REFERENCES "categories"("id");
		ALTER TABLE "messages" ADD CONSTRAINT "messages_fk1" FOREIGN KEY ("author_id") REFERENCES "users"("id");
	`
	insertMessageQuery = `
		INSERT INTO messages ("id", "text", "category_id", "posted_at", "author_id")
		VALUES($1, $2, $3, $4, $5);
	`

	dropMessageTableQuery = `
		DROP TABLE IF EXISTS messages;
	`

	getMessagesQuery = `
		SELECT m.id, u.id, u.name, m.text, m.posted_at 
		FROM messages m
		INNER JOIN users u ON m.author_id = u.id AND m.category_id = $1
		ORDER BY m.posted_at DESC
	`
)

func CreateMessageTable(conn *sql.DB) error {
	bdTx, err := conn.Begin()
	if err != nil {
		return err
	}

	_, err = bdTx.Exec(createMessageTableQuery)
	if err != nil {
		return err
	}

	err = bdTx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func InsertMessage(conn *sql.DB, row *Message) error {
	bdTx, err := conn.Begin()
	if err != nil {
		return err
	}

	_, err = bdTx.Exec(insertMessageQuery, row.Id, row.Text, row.CategoryId, time.Now(), row.AuthorId)
	if err != nil {
		return err
	}

	err = bdTx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func DropMessageTable(conn *sql.DB) error {
	bdTx, err := conn.Begin()
	if err != nil {
		return err
	}

	_, err = bdTx.Exec(dropMessageTableQuery)
	if err != nil {
		return err
	}

	err = bdTx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func GetMessages(conn *sql.DB, id string) ([]ForumMessage, error) {
	rows, err := conn.Query(getMessagesQuery, id)
	if err != nil {
		return make([]ForumMessage, 0), err
	}

	var messages []ForumMessage
	for rows.Next() {
		msg := ForumMessage{}
		err := rows.Scan(&msg.Id, &msg.AuthorId, &msg.AuthorName, &msg.Text, &msg.PostedAt)
		if err != nil {
			return make([]ForumMessage, 0), err
		}
		messages = append(messages, msg)
	}

	if err = rows.Err(); err != nil {
		return make([]ForumMessage, 0), err
	}

	if len(messages) == 0 {
		return make([]ForumMessage, 0), nil
	}
	return messages, nil
}
