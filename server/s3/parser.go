package s3

import (
	"bufio"
	"encoding/json"
	"os"
)

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
