package s3

import (
	"fmt"
	"strings"

	"github.com/rayhaanbhikha/localstack-s3-ui/api"
)

type S3Resource struct {
	// TODO: add metadata property.
	BucketName  string        `json:"bucketName"`
	Name        string        `json:"name"`
	Type        string        `json:"type"`
	parentDirs  []string      `json:"-"`
	Resources   []*S3Resource `json:"resources"`
	Path        string        `json:"path"`
	currentPath string        `json:"-"`
	Data        string        `json:"data"`
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

func (r *S3Resource) traversePath() {
	r.currentPath += "/" + r.parentDirs[0]
	if len(r.parentDirs) > 1 {
		r.parentDirs = r.parentDirs[1:]
	} else {
		r.parentDirs = nil
	}
}

func (r *S3Resource) Add(resource *S3Resource) {
	r.Resources = addResource(r.Resources, resource)
}

func NewS3Resource(a *api.ApiRequest) *S3Resource {
	splitFn := func(c rune) bool {
		return c == '/'
	}
	path := strings.FieldsFunc(a.Path, splitFn)
	n := len(path)

	return &S3Resource{
		BucketName:  path[0],
		Name:        path[n-1],
		Type:        "File",
		parentDirs:  path[1 : n-1],
		Path:        a.Path,
		currentPath: "/" + path[0],
		Data:        a.Data,
	}
}
