package db

import (
	"database/sql"

	"github.com/rayhaanbhikha/localstack-s3-ui/s3"
)

// AddBucket creates a new bucket
func (db *DB) AddBucket(bucketName string) (sql.Result, error) {
	statement := `REPLACE INTO bucket(name) VALUES(?) ON CONFLICT(name) DO NOTHING;`
	stmt, err := db.Conn.Prepare(statement)
	if err != nil {
		return nil, err
	}
	return stmt.Exec(bucketName, bucketName)
}

// AddResource creates a new resource
func (db *DB) AddResource(s3Resource *s3.S3Resource) (sql.Result, error) {
	statement := `
	REPLACE INTO resource_file(path, name, data, bucket_name, parent_path) 
	VALUES(?,?,?,
		(SELECT bucket.name FROM bucket WHERE bucket.name=?)
		(SELECT resource_dir.path FROM resource_dir WHERE resource_dir.path=?)
	) ON CONFLICT(path) DO UPDATE SET data=?;
`
	stmt, err := db.Conn.Prepare(statement)
	if err != nil {
		return nil, err
	}
	return stmt.Exec(
		s3Resource.Path,
		s3Resource.Name,
		s3Resource.Data,
		s3Resource.BucketName,
		s3Resource.ParentPath,
		s3Resource.Data)
}

// AddFolderResource creates a new empty 'directory' resource
func (db *DB) AddFolderResource(s3Resource *s3.S3Resource) (sql.Result, error) {
	statement := `
		REPLACE INTO resource_dir(path, name, bucket_name, parent_path) 
		VALUES(?,?,
			(SELECT bucket.name FROM bucket WHERE bucket.name=? ),
			(SELECT resource_dir.path FROM resource_dir WHERE resource_dir.path=?)
		)  ON CONFLICT(path) DO NOTHING;
	`
	stmt, err := db.Conn.Prepare(statement)
	if err != nil {
		return nil, err
	}
	return stmt.Exec(s3Resource.Path, s3Resource.Name, s3Resource.BucketName, s3Resource.ParentPath)
}
