package s3

import (
	"log"
)

// TODO: Child nodes could be a map with the path as a key

func M() {

	s3Resources, err := Parse("./recorded_api_calls.mock.json")
	if err != nil {
		log.Fatal(err)
	}

	rootNode := &S3Node{Name: "Root", Path: "/", Type: "Root"}

	for _, s3Resource := range s3Resources {
		// fmt.Println(s3Resource)
		rootNode.addNode(s3Resource)
	}

	rootNode.Print()

	rootNode.GetNodesAtPath("/static-resource")
}
