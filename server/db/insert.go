package db

import (
	"database/sql"
)

// AddBucket creates a new bucket
func (db *DB) AddBucket(bucketName string) (sql.Result, error) {
	stmt, err := db.Conn.Prepare("INSERT INTO bucket(name) VALUES(?)")
	if err != nil {
		return nil, err
	}
	return stmt.Exec(bucketName)
}

// AddResource creates a new resource
func (db *DB) AddResource(data string) (sql.Result, error) {
	var (
		path       = "static-resources/folder1/index.html"
		name       = "index.html"
		typeof     = "file"
		bucketName = "static-resources"
	)
	statement := `
	REPLACE INTO resource(path, name, typeof, bucket_name, data) 
	values(?,?,?,(SELECT bucket.name FROM bucket WHERE bucket.name=?),?) 
	ON CONFLICT(path) DO UPDATE SET data=?
	`
	stmt, err := db.Conn.Prepare(statement)
	if err != nil {
		return nil, err
	}
	return stmt.Exec(path, name, typeof, bucketName, data, data)
}
