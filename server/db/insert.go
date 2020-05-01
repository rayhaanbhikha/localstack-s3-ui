package db

import (
	"database/sql"

	"github.com/rayhaanbhikha/localstack-s3-ui/s3"
)

// AddBucket creates a new bucket
func (db *DB) AddBucket(bucketName string) (sql.Result, error) {
	statement := `
		INSERT INTO bucket(name) 
		SELECT ? WHERE NOT EXISTS(SELECT bucket.name from bucket WHERE bucket.name=?);
	`
	stmt, err := db.Conn.Prepare(statement)
	if err != nil {
		return nil, err
	}
	return stmt.Exec(bucketName, bucketName)
}

// AddResource creates a new resource
func (db *DB) AddResource(s3Resource *s3.S3Resource) (sql.Result, error) {
	statement := `
	REPLACE INTO resource(path, parent_path, name, typeof, bucket_name, data) 
	values(?,?,?,?,(SELECT bucket.name FROM bucket WHERE bucket.name=?),?) 
	ON CONFLICT(path) DO UPDATE SET data=?
	`
	stmt, err := db.Conn.Prepare(statement)
	if err != nil {
		return nil, err
	}
	return stmt.Exec(s3Resource.Path, s3Resource.ParentPath, s3Resource.Name, s3Resource.Type, s3Resource.BucketName, s3Resource.Data, s3Resource.Data)
}
