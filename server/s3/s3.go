package s3

import (
	"encoding/json"
)

func RootNode() *S3Node {
	return &S3Node{Name: "Root", Path: "/", Type: "Root", children: make(map[string]*S3Node)}
}

func (n *S3Node) Init(filePath string) error {
	s3Requests, err := Parse(filePath)
	if err != nil {
		return err
	}

	for _, s3Request := range s3Requests {
		if s3Request.Method == "PUT" {
			n.addNode(s3Request.actualPath, s3Request.Data)
		}
	}
	return nil
}

func (n *S3Node) Json(resourcePath string) ([]byte, error) {

	nodes := make([]*S3Node, 0)
	node, ok := n.getNode(resourcePath)

	if ok {
		for _, childNode := range node.children {
			nodes = append(nodes, childNode)
		}
	}

	data, err := json.Marshal(struct {
		Name     string    `json:"name"`
		Path     string    `json:"path"`
		Children []*S3Node `json:"children"`
	}{
		Name:     node.Name,
		Path:     resourcePath,
		Children: nodes,
	})
	// data, err := json.Marshal(nodes)

	if err != nil {
		return nil, err
	}
	return data, nil
}
