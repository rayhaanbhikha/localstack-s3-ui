package s3

import (
	"path"
)

type S3Bucket struct {
	Name      string
	Resources []*S3Resource
}

func NewS3Bucket(name string) *S3Bucket {
	return &S3Bucket{Name: name, Resources: make([]*S3Resource, 0)}
}

func (s3B *S3Bucket) add(resource *S3Resource) {

	if len(resource.parentDirs) == 0 {
		// will be adding/replacing in this resource array at this currentPath.
		for index, eResource := range s3B.Resources {
			if eResource.Name == resource.Name {
				s3B.Resources[index] = resource
				return
			}
		}
		s3B.Resources = append(s3B.Resources, resource)
		return
	}

	dirToFind := resource.parentDirs[0]

	// pass resource on to Dir resource.
	for index, eResource := range s3B.Resources {
		if eResource.Name == dirToFind && eResource.Type == "Directory" {
			resource.traversePath()
			s3B.Resources[index].add(resource)
			return
		}
	}

	// add Empty Dir.
	s3B.Resources = append(s3B.Resources, &S3Resource{
		Name:       dirToFind,
		Type:       "Directory",
		BucketName: resource.BucketName,
		Path:       path.Join(resource.currentPath, dirToFind),
	})

	s3B.add(resource)
}

func (s3B *S3Bucket) delete(resource *S3Resource) {
	if len(resource.parentDirs) == 0 {
		// delete at this level.
		for index, eResource := range s3B.Resources {
			if eResource.Name == resource.Name {
				s3B.Resources = updateResource(index, s3B.Resources)
				return
			}
		}
	}

	dirToFind := resource.parentDirs[0]

	// pass resource on to Dir resource.
	for index, eResource := range s3B.Resources {
		if eResource.Name == dirToFind && eResource.Type == "Directory" {
			resource.traversePath()
			s3B.Resources[index].delete(resource)
			return
		}
	}
}
