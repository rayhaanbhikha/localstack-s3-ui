package s3

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type S3APICalls struct {
	Data []*S3APICall `json:"data"`
}

func (a *S3APICalls) add(apicall *S3APICall) {
	a.Data = append(a.Data, apicall)
}

// TODO: need to write transformer.
type S3APICall struct {
	Type   string `json:"a"`
	Method string `json:"m"`
	Path   string `json:"p"`
	Data   string `json:"d"`
}

func (a *S3APICall) String() string {
	return fmt.Sprintf(`
		Type: %s,
		Method: %s,
		Path: %s,
		Data: %s,		
	`, a.Type, a.Method, a.Path, a.Data)
}

func Parser() *S3APICalls {
	file, err := os.Open("./recorded_api_calls.json")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	s3APICalls := &S3APICalls{}

	for scanner.Scan() {
		data := scanner.Bytes()
		s3APICall := &S3APICall{}
		err := json.Unmarshal(data, s3APICall)
		// TODO: remove
		if err != nil {
			panic(err)
		}
		if s3APICall.Type == "s3" {
			s3APICalls.add(s3APICall)
		}
	}
	return s3APICalls
}
