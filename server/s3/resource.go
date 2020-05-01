package s3

import (
	"fmt"
	"strings"
)

type S3Resource struct {
	// TODO: add metadata property.
	Method      string `json:"-"`
	BucketName  string `json:"bucketName"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	parentDirs  []string
	Resources   []*S3Resource `json:"resources"`
	ParentPath  string        `json:"-"`
	Path        string        `json:"path"`
	currentPath string
	Data        string `json:"data"`
}

func (r *S3Resource) String() string {
	return fmt.Sprintf(`
		Bucket: %s
		Name: %s
		Type: %s
		parentDirs: %v
		Path: %s
		CurrentPath: %s
	`, r.BucketName, r.Name, r.Type, r.parentDirs, r.Path, r.currentPath)
}

func newS3Resource(genRequest *GenRequest) *S3Resource {
	splitFn := func(c rune) bool {
		return c == '/'
	}
	path := strings.FieldsFunc(genRequest.Path, splitFn)
	n := len(path)

	if len(path) == 1 {
		return &S3Resource{
			Method:      genRequest.Method,
			BucketName:  path[0],
			Type:        "Bucket",
			Path:        genRequest.Path,
			currentPath: genRequest.Path,
			Data:        genRequest.Data,
		}
	}

	return &S3Resource{
		Method:      genRequest.Method,
		BucketName:  path[0],
		Name:        path[n-1],
		Type:        "File",
		ParentPath:  strings.Join(path[:n-1], "/"),
		parentDirs:  path[1 : n-1],
		Path:        genRequest.Path,
		currentPath: "/" + path[0],
		Data:        genRequest.Data,
	}
}
