package db

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	createBucketStmt = `
	CREATE TABLE IF NOT EXISTS bucket(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE
	);
	`
	createResourceFileTableStmt = `
	CREATE TABLE resource_file(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		path TEXT UNIQUE NOT NULL,
		parent_path TEXT,
		name TEXT NOT NULL,
		type TEXT DEFAULT 'File',
		data BLOB,
		headers BLOB,
		bucket_name TEXT NOT NULL,
		FOREIGN KEY (bucket_name) REFERENCES bucket(name) ON DELETE CASCADE,
		FOREIGN KEY (parent_path) REFERENCES resource_dir(path) ON DELETE CASCADE
	);
`
	createResourceDirTableStmt = `
	CREATE TABLE resource_dir(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		path TEXT UNIQUE NOT NULL,
		name TEXT NOT NULL,
		type TEXT DEFAULT "Directory",
		bucket_name TEXT NOT NULL,
		FOREIGN KEY(bucket_name) REFERENCES bucket(name) ON DELETE CASCADE
	);
	`
	alterResourceDirTableStmt = `ALTER TABLE resource_dir ADD COLUMN parent_path TEXT REFERENCES resource_dir(path) ON DELETE CASCADE;`
)

// SetUp bucket and resource table.
func (db *DB) SetUp() (sql.Result, error) {
	// FIXME: error handling can be improved.

	tx, err := db.Conn.Begin()
	handleErr(err)

	res, err := tx.Exec(createBucketStmt)
	handleTxError(tx, err)

	id, err := res.LastInsertId()
	handleErr(err)
	fmt.Println(id)

	res, err = tx.Exec(createResourceDirTableStmt)
	handleTxError(tx, err)
	handleErr(err)
	fmt.Println(id)

	res, err = tx.Exec(alterResourceDirTableStmt)
	handleTxError(tx, err)
	handleErr(err)
	fmt.Println(id)

	res, err = tx.Exec(createResourceFileTableStmt)
	handleTxError(tx, err)
	handleErr(err)
	fmt.Println(id)

	handleErr(tx.Commit())

	return res, nil
}

func handleTxError(tx *sql.Tx, err error) {
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
