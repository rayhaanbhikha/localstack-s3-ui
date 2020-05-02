package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"path"

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
	http.HandleFunc("/page", pageHandler(rootNode))

	log.Printf("About to listen on 8080. Go to https://127.0.0.1:8080/")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func pageHandler(rootNode *s3.Node) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		queryParams := r.URL.Query()

		if v, ok := queryParams["path"]; ok {
			resourcePath := path.Clean(v[0])
			fmt.Println("Resource Path requested: ", resourcePath)
			node, ok := rootNode.Get(resourcePath)
			if !ok {
				http.NotFound(w, r)
				return
			}
			decoded, err := base64.StdEncoding.DecodeString(node.Data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// TODO: node should have content type. (chrome is smart enough to know what the mime type is.)
			// w.Header().Set("Content-Type", "text/javascript")
			w.Write([]byte(decoded))
		} else {
			http.NotFound(w, r)
			return
		}
	}
}

func resourceHandler(rootNode *s3.Node) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		queryParams := r.URL.Query()

		if v, ok := queryParams["path"]; ok {
			resourcePath := path.Clean(v[0])
			fmt.Println("Resource Path requested: ", resourcePath)
			json, err := rootNode.JSON(resourcePath)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(json)
		} else {
			http.NotFound(w, r)
			return
		}
	}
}
