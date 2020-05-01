package db

import (
	"database/sql"

	"github.com/rayhaanbhikha/localstack-s3-ui/s3"
)

// DeleteBucket creates a new bucket
func (db *DB) DeleteBucket(bucketName string) (sql.Result, error) {
	stmt, err := db.Conn.Prepare("DELETE FROM bucket WHERE bucket.name=?;")
	if err != nil {
		return nil, err
	}
	return stmt.Exec(bucketName)
}

// DeleteResource creates a new resource
func (db *DB) DeleteResource(s3Resource *s3.S3Resource) (sql.Result, error) {
	stmt, err := db.Conn.Prepare("DELETE FROM resource WHERE resource.path=?")
	if err != nil {
		return nil, err
	}
	return stmt.Exec(s3Resource.Path)
}
