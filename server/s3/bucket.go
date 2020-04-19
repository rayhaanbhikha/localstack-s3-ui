package s3

import "fmt"

type S3Bucket struct {
	resources []*S3Resource
}

func (r *S3Resource) add(resource *S3Resource) {
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
