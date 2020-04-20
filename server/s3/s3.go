package s3

import (
	"encoding/json"
	"strings"

	"github.com/rayhaanbhikha/localstack-s3-ui/api"
)

type S3 map[string]*S3Bucket

func New() S3 {
	return make(S3)
}

func (s3 S3) Json() []byte {
	// FIXME: remove indentation
	data, err := json.MarshalIndent(s3, "", "  ")
	if err != nil {
		panic(err)
	}
	return data
}

func (s3 S3) Add(a *api.ApiRequest) {
	splitFn := func(c rune) bool {
		return c == '/'
	}
	path := strings.FieldsFunc(a.Path, splitFn)

	bucketName := path[0]
	if len(path) == 1 {
		if _, ok := s3[bucketName]; !ok {
			s3[bucketName] = NewS3Bucket(bucketName)
		}
		return
	}

	// TODO: need to check if "d" is empty or not. represents a delete file
	// TODO: need to consider type of method. i.e. DELETE || PUT
	s3[bucketName].Add(NewS3Resource(a))
}
