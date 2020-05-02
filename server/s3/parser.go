package s3

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type apiRequest struct {
	Type       string `json:"a"`
	Method     string `json:"m"`
	Path       string `json:"p"`
	Data       string `json:"d"`
	actualPath []string
}

// parse ... parse API requests in file.
func parse(filePath string) ([]*apiRequest, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	apiRequests := make([]*apiRequest, 0)

	for scanner.Scan() {
		data := scanner.Bytes()
		genRequest := &apiRequest{}
		err := json.Unmarshal(data, genRequest)
		if err != nil {
			fmt.Println("Err parsing api requests: ", err.Error())
			continue
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
