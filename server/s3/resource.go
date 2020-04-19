package s3

import (
	"fmt"
	"strings"

	"github.com/rayhaanbhikha/localstack-s3-ui/api"
)

type S3Resource struct {
	// TODO: add metadata property.
	BucketName  string
	Name        string
	Type        string
	ParentDirs  []string
	Resources   []*S3Resource
	ActualPath  string
	CurrentPath string
	Data        string
}

func (r *S3Resource) String() string {
	return fmt.Sprintf(`
		Bucket: %s
		Name: %s
		Type: %s
		ParentDirs: %v
		ActualPath: %s
		CurrentPath: %s
	`, r.BucketName, r.Name, r.Type, r.ParentDirs, r.ActualPath, r.CurrentPath)
}

func (r *S3Resource) UpdatePath() {
	r.CurrentPath += "/" + r.ParentDirs[0]
	if len(r.ParentDirs) > 1 {
		r.ParentDirs = r.ParentDirs[1:]
	}
}

func (r *S3Resource) Add(resource *S3Resource) {
	resource.UpdatePath()

	for _, existingResource := range r.Resources {
		switch {
		case existingResource.Name == resource.Name:
			existingResource.Data = resource.Data
			return
		case existingResource.CurrentPath == resource.CurrentPath:
			existingResource.Add(resource)
			return
		}
	}

	// brand new resource which may need flattening.
	dirs := generateNestedDirResource(resource)
	r.Resources = append(r.Resources, dirs)
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
		ParentDirs:  path[1 : n-1],
		ActualPath:  a.Path,
		CurrentPath: "/" + path[0],
		Data:        a.Data,
	}
}

func EmptyDir(resource *S3Resource) *S3Resource {
	splitFn := func(c rune) bool {
		return c == '/'
	}
	path := strings.FieldsFunc(resource.ActualPath, splitFn)
	pathLen := len(path)
	parentDirLen := len(resource.ParentDirs)

	currentPath := "/" + strings.Join(path[:pathLen-1], "/")
	return &S3Resource{
		BucketName:  resource.BucketName,
		Name:        resource.ParentDirs[parentDirLen-1],
		ActualPath:  currentPath,
		CurrentPath: currentPath,
		ParentDirs:  resource.ParentDirs[:parentDirLen-1],
		Type:        "Directory",
		Resources: []*S3Resource{
			resource,
		},
	}
}
