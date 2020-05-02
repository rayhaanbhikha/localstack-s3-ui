package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rayhaanbhikha/localstack-s3-ui/s3"
)

func main() {

	fmt.Println("hello")
	rootNode, err := s3.Init("./recorded_api_calls.mock.json")
	if err != nil {
		panic(err)
	}

	// rootNode.Print()
	// json, _ := rootNode.Json("/static-resources")
	// file, _ := os.Create("data_2.json")

	// defer file.Close()

	// file.Write(json)
	http.HandleFunc("/resource", resourceHandler(rootNode))

	log.Printf("About to listen on 8080. Go to https://127.0.0.1:8080/")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func resourceHandler(rootNode *s3.S3Node) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		queryParams := r.URL.Query()
		if v, ok := queryParams["path"]; ok {
			fmt.Println("Path: ", v[0])
			w.Header().Set("Content-Type", "application/json")
			json, err := rootNode.Json(v[0])
			if err != nil {
				// w.Write([]byte(err.Error()))
				return
			}
			w.Write(json)
		}
	}
}
