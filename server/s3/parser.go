package s3

import (
	"log"
)

// TODO: Child nodes could be a map with the path as a key

func M() {

	s3Requests, err := Parse("./recorded_api_calls.mock.json")
	if err != nil {
		log.Fatal(err)
	}

	rootNode := &S3Node{Name: "Root", Path: "/", Type: "Root"}

	for _, s3Request := range s3Requests {
		// fmt.Println(s3Resource)
		if s3Request.Method == "PUT" {
			rootNode.addNode(s3Request.actualPath, s3Request.Data)
		}
	}

	rootNode.Print()

	// rootNode.GetNodesAtPath("/static-resource")
}
