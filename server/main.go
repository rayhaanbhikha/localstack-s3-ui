package main

import (
	"log"

	"github.com/rayhaanbhikha/localstack-s3-ui/file"
	"github.com/rayhaanbhikha/localstack-s3-ui/s3"
	"github.com/rayhaanbhikha/localstack-s3-ui/utils"
)

func main() {

	filePath := utils.GetFilePath()
	rootNode := s3.RootNode()

	err := rootNode.LoadData(filePath)
	if err != nil {
		log.Fatal(err)
	}

	watcher, err := file.Watch(filePath, func() {
		rootNode.LoadData(filePath)
	})
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	startServer(rootNode)
}
