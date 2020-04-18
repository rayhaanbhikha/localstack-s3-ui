package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rayhaanbhikha/localstack-s3-ui/s3"
)

func main() {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	// 	w.Write()
	// })

	http.HandleFunc("/data", dataHandler)

	log.Printf("About to listen on 8080. Go to https://127.0.0.1:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// TODO: should be parsing file from a particular place.
	s3ApiCall := s3.Parser()

	data, err := json.Marshal(s3ApiCall.Data)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
