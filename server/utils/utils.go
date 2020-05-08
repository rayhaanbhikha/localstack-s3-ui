package utils

import (
	"log"
	"os"
)

// GetFilePath ...
func GetFilePath() string {
	filePath := os.Getenv("FILE_PATH")

	if filePath == "" {
		log.Fatalln("Missing env 'FILE_PATH'")
	}

	fileInfo, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		log.Fatalf("cannot find file at file path: '%s'", filePath)
	}

	if fileInfo.IsDir() {
		log.Fatalf("'%s' is a directory not a file.", filePath)
	}

	return filePath
}

// GetPort ...
func GetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		return "8080"
	}
	return port
}
