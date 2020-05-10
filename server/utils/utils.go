package utils

import (
	"log"
	"os"
	"time"
)

func checkFileExists(filePath string) bool {
	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	if fileInfo.IsDir() {
		log.Fatalf("%s is a directory not a file", filePath)
	}
	return true
}

func waitForFile(filePath string) {
	start := time.Now()
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	log.Printf("Waiting for file %s", filePath)
	for {
		select {
		case t := <-ticker.C:
			duration := t.Sub(start)
			log.Printf("Waiting for file %s: %s", filePath, duration.Truncate(time.Second).String())
			if duration.Seconds() > 60.0 {
				log.Fatalf("Timed out waiting for file.")
			}
			ok := checkFileExists(filePath)
			if ok {
				log.Printf("file %s created.\n", filePath)
				return
			}
		}
	}
}

// GetFilePath ...
func GetFilePath() string {
	apiReqFilePath := os.Getenv("API_REQ_FILE_PATH")

	if apiReqFilePath == "" {
		log.Fatalln("Missing env 'API_REQ_FILE_PATH'")
	}

	// implement solution to wait for file to exist ... up until 1min.
	waitForFile(apiReqFilePath)

	return apiReqFilePath
}

// GetPort ...
func GetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		return "9000"
	}
	return port
}
