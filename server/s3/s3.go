package s3

import (
	"encoding/json"

	"github.com/rayhaanbhikha/localstack-s3-ui/api"
)

type S3 struct {
	Buckets map[string][]*S3Resource `json:"buckets"`
}

func New() *S3 {
	return &S3{Buckets: make(map[string][]*S3Resource)}
}
func (s3 *S3) Json() []byte {
	data, err := json.Marshal(s3)
	if err != nil {
		panic(err)
	}
	return data
}

func (s3 S3) Add(a *api.ApiRequest) {
	resource := NewS3Resource(a)
	bucket, ok := s3.Buckets[resource.Bucket]
	if resource.Type == "Bucket" && !ok {
		s3.Buckets[resource.Bucket] = make(S3Resources, 0)
	} else {
		bucket.add(resource)
	}
}
