package s3

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type S3Resource struct {
	// TODO: add metadata property.
	Method     string   `json:"-"`
	BucketName string   `json:"bucketName"`
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	Path       []string `json:"path"`
	Data       string   `json:"data"`
}

func (r *S3Resource) String() string {
	return fmt.Sprintf(`
		Bucket: %s
		Name: %s
		Type: %s
		Path: %s
	`, r.BucketName, r.Name, r.Type, r.Path)
}

func newS3Resource(genRequest *GenRequest) *S3Resource {
	splitFn := func(c rune) bool {
		return c == '/'
	}
	path := strings.FieldsFunc(genRequest.Path, splitFn)
	n := len(path)

	if len(path) == 1 {
		return &S3Resource{
			Method:     genRequest.Method,
			BucketName: path[0],
			Name:       path[0],
			Type:       "Bucket",
			Path:       path,
			Data:       genRequest.Data,
		}
	}

	return &S3Resource{
		Method:     genRequest.Method,
		BucketName: path[0],
		Name:       path[n-1],
		Type:       "File",
		Path:       path,
		Data:       genRequest.Data,
	}
}

type GenRequest struct {
	Type   string `json:"a"`
	Method string `json:"m"`
	Path   string `json:"p"`
	Data   string `json:"d"`
}

// Parse ... Parse API requests in file.
func Parse(filePath string) ([]*S3Resource, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	resources := make([]*S3Resource, 0)
	for scanner.Scan() {
		data := scanner.Bytes()
		genRequest := &GenRequest{}
		err := json.Unmarshal(data, genRequest)
		if err != nil {
			return nil, err
		}

		// if api request type is s3.
		if genRequest.Type == "s3" {
			resources = append(resources, newS3Resource(genRequest))
		}
	}
	return resources, nil
}
