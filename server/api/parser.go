package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Request ...
type Request struct {
	Type   string `json:"a"`
	Method string `json:"m"`
	Path   string `json:"p"`
	Data   string `json:"d"`
}

func (a *Request) String() string {
	return fmt.Sprintf(`
		Type: %s,
		Method: %s,
		Path: %s,
		Data: %s,		
	`, a.Type, a.Method, a.Path, a.Data)
}

// Parse ... Parse API requests in file.
func Parse(filePath string) ([]*Request, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	apiRequests := make([]*Request, 0)
	for scanner.Scan() {
		data := scanner.Bytes()
		apiRequest := &Request{}
		err := json.Unmarshal(data, apiRequest)
		if err != nil {
			return nil, err
		}
		apiRequests = append(apiRequests, apiRequest)
	}
	return apiRequests, nil
}
