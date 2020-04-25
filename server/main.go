package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// http.HandleFunc("/data", dataHandler)
	http.HandleFunc("/echo", echoHandler)

	log.Printf("About to listen on 8080. Go to https://127.0.0.1:8080/")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

type echoReq struct {
	Data string `json:"data"`
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	echo := &echoReq{}
	b := make([]byte, 0)
	r.Body.Read(b)
	// check if error on close.
	r.Body.Close()
	fmt.Println(string(b))
	json.Unmarshal(b, echo)

	fmt.Fprintf(w, " %s %s %s", "hello world", r.Method, echo.Data)
}

// func dataHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")

// 	// TODO: should be parsing file from a particular place.
// 	s3 := &s3.New()

// 	data, err := json.Marshal(s3ApiCall.Data)
// 	if err != nil {
// 		panic(err)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(data)
// }
