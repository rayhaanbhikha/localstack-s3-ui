package db

import (
	"database/sql"
	"fmt"
	"os"

	// sqllite db driver.
	_ "github.com/mattn/go-sqlite3"
)

// DB struct as a wrapper around *sql.DB
type DB struct {
	Conn *sql.DB
}

// Init create connection with db
func Init(fileName string, reset bool) (*DB, error) {
	if reset {
		err := os.Remove(fileName)
		if err != nil {
			return nil, err
		}
	}
	DSN := fmt.Sprintf("%s?_foreign_keys=true", fileName)
	fmt.Println(DSN)
	db, err := sql.Open("sqlite3", DSN)
	if err != nil {
		return nil, err
	}
	return &DB{Conn: db}, nil
}
