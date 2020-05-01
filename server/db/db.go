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

// SetUp bucket and resource table.
func (db *DB) SetUp() (sql.Result, error) {
	// create bucket table
	sqlStatement := `
		CREATE TABLE IF NOT EXISTS bucket(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT UNIQUE
		);
		`

	stmt, err := db.Conn.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec()
	if err != nil {
		return nil, err
	}

	fmt.Println(res.RowsAffected())

	// create resource table
	sqlStatement = `
		CREATE TABLE IF NOT EXISTS resource(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			path TEXT UNIQUE NOT NULL,
			parent_path TEXT NOT NULL,
			name TEXT NOT NULL,
			typeof TEXT NOT NULL,
			data BLOB,
			headers BLOB,
			bucket_name TEXT NOT NULL,
			FOREIGN KEY (bucket_name) REFERENCES bucket(name) ON DELETE CASCADE,
			CHECK(typeof = "File" OR typeof = "Directory")
		);
		`

	stmt, err = db.Conn.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.Exec()
}
