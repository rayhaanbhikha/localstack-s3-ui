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
	parentDirs  []string
	Resources   []*S3Resource
	Path        string
	currentPath string
	Data        string
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

func (r *S3Resource) add(resource *S3Resource) {

	if len(resource.parentDirs) == 0 {
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

	dirToFind := resource.parentDirs[0]

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
		Path:       path.Join(resource.currentPath, dirToFind),
	})

	r.add(resource)
}

func (r *S3Resource) delete(resource *S3Resource) {
	fmt.Println("Passed on: ", resource)

	if len(resource.parentDirs) == 0 {
		// delete at this level.
		for index, eResource := range r.Resources {
			if eResource.Name == resource.Name {
				r.Resources = updateResource(index, r.Resources)
				return
			}
		}
	}

	dirToFind := resource.parentDirs[0]

	// pass resource on to Dir resource.
	for index, eResource := range r.Resources {
		if eResource.Name == dirToFind && eResource.Type == "Directory" {
			resource.traversePath()
			r.Resources[index].delete(resource)
			return
		}
	}
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

func updateResource(index int, resource []*S3Resource) []*S3Resource {
	rLen := len(resource)
	switch {
	case index == 0:
		return resource[1:]
	case index == rLen-1:
		return resource[:index]
	default:
		return append(resource[:index], resource[index+1:]...)
	}
}
