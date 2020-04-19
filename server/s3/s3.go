package s3

import (
	"encoding/json"

	"github.com/rayhaanbhikha/localstack-s3-ui/api"
)

type S3 struct {
	Buckets []*S3Bucket `json:"buckets"`
}

func New() *S3 {
	return &S3{}
}

func (s3 *S3) Json() []byte {
	data, err := json.Marshal(s3)
	if err != nil {
		panic(err)
	}
	return data
}

func (s3 *S3) Add(a *api.ApiRequest) {
	resource := NewS3Resource(a)

	for index, bucket := range s3.Buckets {
		if bucket.Name == resource.Bucket {
			s3.Buckets[index].add(resource)
			return
		}
	}
	s3.Buckets = append(s3.Buckets, NewS3Bucket(resource))

}
