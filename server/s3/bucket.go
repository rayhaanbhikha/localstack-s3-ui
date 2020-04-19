package s3

import "fmt"

type S3Bucket struct {
	Name      string
	Resources []*S3Resource
}

func NewS3Bucket(resource *S3Resource) *S3Bucket {
	return &S3Bucket{Name: resource.Bucket, Resources: make([]*S3Resource, 0)}
}

func (s3B *S3Bucket) add(resource *S3Resource) {
	resourceName := resource.Path[0]

	// for _, existingResource := range s3Bucket

	// foundResource, ok := s3Bucket[resourceName]
	// if !ok {
	// 	resource.UpdatePath(resource.Path[1:])
	// 	s3Bucket[resourceName] = resource
	// 	if len(resource.Path) != 0 {
	// 		resource
	// 	}
	// }

	fmt.Println(resourceName, resource.Path[1:])
}
