package main

import (
	"log"

	"github.com/rayhaanbhikha/localstack-s3-ui/s3"
)

func main() {
	// TODO: os.GetEnv file path.
	fileName := "./mock-data/recorded_api_calls.json"
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
	startServer(rootNode)
}
