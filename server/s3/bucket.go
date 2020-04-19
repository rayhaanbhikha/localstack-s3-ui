package s3

import "fmt"

type S3Bucket struct {
	Name      string
	Resources []*S3Resource
}

func NewS3Bucket(name string) *S3Bucket {
	return &S3Bucket{Name: name, Resources: make([]*S3Resource, 0)}
}

func (s3B *S3Bucket) add(resource *S3Resource) {

	if len(resource.ParentDirs) == 0 {
		s3B.Resources = append(s3B.Resources, resource)
		return
	}

	fmt.Println(resource)
	resource.UpdatePath()

	for _, existingResource := range s3B.Resources {
		fmt.Println(existingResource.CurrentPath, resource.CurrentPath)
		if existingResource.CurrentPath == resource.CurrentPath {
			existingResource.Add(resource)
			fmt.Println("nested resource: ", resource)
			return
		}
	}

	// brand new resource which may need flattening.
	dirs := generateNestedDirResource(resource)
	s3B.Resources = append(s3B.Resources, dirs)

}

func generateNestedDirResource(resource *S3Resource) *S3Resource {
	parentDir := EmptyDir(resource)
	if len(parentDir.ParentDirs) != 0 {
		return generateNestedDirResource(parentDir)
	}
	return parentDir
}
