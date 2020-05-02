package s3

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"
)

type APIRequest struct {
	Type       string `json:"a"`
	Method     string `json:"m"`
	Path       string `json:"p"`
	Data       string `json:"d"`
	actualPath []string
}

// Parse ... Parse API requests in file.
func Parse(filePath string) ([]*APIRequest, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	apiRequests := make([]*APIRequest, 0)

	for scanner.Scan() {
		data := scanner.Bytes()
		genRequest := &APIRequest{}
		err := json.Unmarshal(data, genRequest)
		if err != nil {
			return nil, err
		}
		genRequest.actualPath = strings.FieldsFunc(genRequest.Path, func(c rune) bool {
			if c == '/' {
				return true
			}
			return false
		})
		// if api request type is s3.
		if genRequest.Type == "s3" {
			apiRequests = append(apiRequests, genRequest)
		}
	}
	return apiRequests, nil
}
