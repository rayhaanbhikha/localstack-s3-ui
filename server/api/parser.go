package api

import (
	"fmt"
)

type ApiRequest struct {
	Type   string `json:"a"`
	Method string `json:"m"`
	Path   string `json:"p"`
	Data   string `json:"d"`
}

func (a *ApiRequest) String() string {
	return fmt.Sprintf(`
		Type: %s,
		Method: %s,
		Path: %s,
		Data: %s,		
	`, a.Type, a.Method, a.Path, a.Data)
}

// func Parser(filePath string) []byte {
// 	file, err := os.Open(filePath)
// 	defer file.Close()
// 	if err != nil {
// 		panic(err)
// 	}
// 	scanner := bufio.NewScanner(file)
// 	s3 := &s3.S3{}
// 	for scanner.Scan() {
// 		data := scanner.Bytes()
// 		apiRequest := &ApiRequest{}
// 		err := json.Unmarshal(data, apiRequest)
// 		if err != nil {
// 			panic(err)
// 		}

// 		if strings.ToLower(apiRequest.Type) == "s3" {
// 			s3.AddResource(apiRequest)
// 		}
// 	}
// 	return s3.Json()
// }
