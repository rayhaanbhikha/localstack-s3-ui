package s3

import (
	"encoding/json"
)

// RootNode ... return root s3 node.
func RootNode() *Node {
	return &Node{Name: "Root", Path: "/", Type: "Root", children: make(map[string]*Node)}
}

// LoadData ... load api requests from localstack.
func (n *Node) LoadData(filePath string) error {
	s3Requests, err := parse(filePath)
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

// JSON ... return json response.
func (n *Node) JSON(resourcePath string) ([]byte, error) {

	nodes := make([]*Node, 0)
	node, ok := n.getNode(resourcePath)

	if ok {
		for _, childNode := range node.children {
			nodes = append(nodes, childNode)
		}
	} else {
		// empty response.
		return []byte("[]"), nil
	}

	data, err := json.Marshal(struct {
		Name     string  `json:"name"`
		Path     string  `json:"path"`
		Data     string  `json:"data,omitempty"`
		Type     string  `json:"type"`
		Children []*Node `json:"children,omitempty"`
	}{
		Name:     node.Name,
		Path:     resourcePath,
		Data:     node.Data,
		Type:     node.Type,
		Children: nodes,
	})

	if err != nil {
		return nil, err
	}
	return data, nil
}
