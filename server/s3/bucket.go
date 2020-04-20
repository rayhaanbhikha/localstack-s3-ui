package s3

type S3Bucket struct {
	Name      string
	Resources []*S3Resource
}

func NewS3Bucket(name string) *S3Bucket {
	return &S3Bucket{Name: name, Resources: make([]*S3Resource, 0)}
}

func (s3B *S3Bucket) Add(resource *S3Resource) {
	s3B.Resources = addResource(s3B.Resources, resource)
}
