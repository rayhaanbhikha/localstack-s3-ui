package s3

import (
	"fmt"
	"strings"
)

type S3Resource struct {
	// TODO: add metadata property.
	Method     string `json:"-"`
	BucketName string `json:"bucketName"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	ParentDirs []string
	Resources  []*S3Resource `json:"resources"`
	ParentPath string        `json:"-"`
	Path       string        `json:"path"`
	Data       string        `json:"data"`
}

func (r *S3Resource) String() string {
	return fmt.Sprintf(`
		Bucket: %s
		Name: %s
		Type: %s
		ParentDirs: %v
		Path: %s
		ParentPath: %s
	`, r.BucketName, r.Name, r.Type, r.ParentDirs, r.Path, r.ParentPath)
}

func newS3Resource(genRequest *GenRequest) *S3Resource {
	splitFn := func(c rune) bool {
		return c == '/'
	}
	path := strings.FieldsFunc(genRequest.Path, splitFn)
	n := len(path)

	if len(path) == 1 {
		return &S3Resource{
			Method:     genRequest.Method,
			BucketName: path[0],
			Type:       "Bucket",
			Path:       genRequest.Path,
			Data:       genRequest.Data,
		}
	}

	return &S3Resource{
		Method:     genRequest.Method,
		BucketName: path[0],
		Name:       path[n-1],
		Type:       "File",
		ParentPath: strings.Join(path[:n-1], "/"),
		ParentDirs: path[1 : n-1],
		Path:       genRequest.Path,
		Data:       genRequest.Data,
	}
}

func EmptyDirResources(resource *S3Resource) []*S3Resource {
	resources := make([]*S3Resource, 0)
	n := len(resource.ParentDirs)
	createResource := func(bucketName, name string, path []string) *S3Resource {
		return &S3Resource{
			Type:       "Directory",
			Name:       name,
			BucketName: bucketName,
			Path:       strings.Join(path, "/"),
			ParentPath: strings.Join(path[:len(path)-1], "/"),
		}
	}
	currentPath := []string{resource.BucketName}
	for _, parentDir := range resource.ParentDirs[:n-1] {
		currentPath = append(currentPath, parentDir)
		fmt.Println(createResource(resource.BucketName, parentDir, currentPath))
	}

	return resources
}
