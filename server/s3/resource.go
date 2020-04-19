package s3

import (
	"fmt"
	"path"
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
	Path        string
	CurrentDir  string
	CurrentPath string
	Data        string
}

func (r *S3Resource) String() string {
	return fmt.Sprintf(`
		Bucket: %s
		Name: %s
		Type: %s
		ParentDirs: %v
		Path: %s
		CurrentDir: %s
		CurrentPath: %s
	`, r.BucketName, r.Name, r.Type, r.ParentDirs, r.Path, r.CurrentDir, r.CurrentPath)
}

func (r *S3Resource) traversePath() {
	r.CurrentPath += "/" + r.ParentDirs[0]
	r.CurrentDir = r.ParentDirs[0]
	if len(r.ParentDirs) > 1 {
		r.ParentDirs = r.ParentDirs[1:]
	} else {
		r.ParentDirs = nil
	}
}

func (r *S3Resource) add(resource *S3Resource) {
	fmt.Println("Passed on: ", resource)

	if len(resource.ParentDirs) == 0 {
		// will be adding/replacing in this resource array at this currentPath.
		for index, eResource := range r.Resources {
			if eResource.Name == resource.Name {
				r.Resources[index] = resource
				return
			}
		}
		r.Resources = append(r.Resources, resource)
		return
	}

	dirToFind := resource.ParentDirs[0]

	// pass resource on to Dir resource.
	for index, eResource := range r.Resources {
		if eResource.Name == dirToFind && eResource.Type == "Directory" {
			resource.traversePath()
			r.Resources[index].add(resource)
			return
		}
	}

	r.Resources = append(r.Resources, &S3Resource{
		Name:       dirToFind,
		Type:       "Directory",
		BucketName: resource.BucketName,
		Path:       path.Join(resource.CurrentPath, dirToFind),
	})

	r.add(resource)
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
		Path:        a.Path,
		CurrentPath: "/" + path[0],
		Data:        a.Data,
	}
}
