package s3

import (
	"fmt"
	"strings"

	"github.com/rayhaanbhikha/localstack-s3-ui/api"
)

type S3Resource struct {
	// TODO: add metadata property.
	Bucket     string
	Type       string
	Path       []string
	ActualPath string
	Resources  []*S3Resource
	Data       string
}

func (r *S3Resource) String() string {
	return fmt.Sprintf(`
		Bucket: %s
		Type: %s
		Path: %v
		ActualPath: %s
	`, r.Bucket, r.Type, r.Path, r.ActualPath)
}

func (r *S3Resource) UpdatePath(path []string) {
	r.Path = path
}

func NewS3Resource(a *api.ApiRequest) *S3Resource {
	splitFn := func(c rune) bool {
		return c == '/'
	}
	path := strings.FieldsFunc(a.Path, splitFn)

	resourceType := "Resource"
	if len(path) == 1 {
		resourceType = "Bucket"
	}

	return &S3Resource{
		Bucket:     path[0],
		Type:       resourceType,
		Path:       path[1:],
		ActualPath: a.Path,
		Data:       a.Data,
	}
}
