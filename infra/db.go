package infra

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "chat.sqlite?_time_format=sqlite")
	return db, err
}

func NewTestDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "../../test_chat.sqlite")
	return db, err
}

func CleanDB(db *sql.DB) error {
	sqlTableNames := `SELECT name FROM sqlite_master WHERE type = "table";`

	var tables []string
	rows, err := db.Query(sqlTableNames)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var t string
		err = rows.Scan(&t)
		tables = append(tables, t)
	}
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()
	for _, t := range tables {
		q := fmt.Sprintf("DELETE FROM %s", t)
		_, err := tx.Exec(q, t)
		if err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
