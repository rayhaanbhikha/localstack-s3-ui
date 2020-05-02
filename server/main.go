package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fsnotify/fsnotify"

	"github.com/rayhaanbhikha/localstack-s3-ui/s3"
)

func startFileWatcher(fileName string, rootNode *s3.Node) (*fsnotify.Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					rootNode.LoadData(fileName)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(fileName)
	if err != nil {
		return nil, err
	}

	return watcher, nil
}

func main() {

	fileName := "./recorded_api_calls.mock.json"
	rootNode := s3.RootNode()

	err := rootNode.LoadData(fileName)
	if err != nil {
		log.Fatal(err)
	}

	watcher, err := startFileWatcher(fileName, rootNode)
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	http.HandleFunc("/resource", resourceHandler(rootNode))

	log.Printf("About to listen on 8080. Go to https://127.0.0.1:8080/")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

// TODO: need to sanitize query param input and handle edge cases.
func resourceHandler(rootNode *s3.Node) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		queryParams := r.URL.Query()
		if v, ok := queryParams["path"]; ok {
			fmt.Println("Path: ", v[0])
			w.Header().Set("Content-Type", "application/json")
			json, err := rootNode.JSON(v[0])
			if err != nil {
				// w.Write([]byte(err.Error()))
				return
			}
			w.Write(json)
		}
	}
}
