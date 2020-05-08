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
		switch s3Request.Method {
		case "PUT":
			n.addNode(s3Request)
		case "DELETE":
			n.deleteNode(s3Request)
		}
	}
	return nil
}

func (n *Node) Get(resourcePath string) (*Node, bool) {
	return n.getNode(resourcePath)
}

// JSON ... return json response.
func (n *Node) JSON(resourcePath string) ([]byte, error) {

	node, ok := n.getNode(resourcePath)
	if !ok {
		return []byte("[]"), nil
	}

	nodes := make([]*Node, 0)
	for _, childNode := range node.children {
		nodes = append(nodes, childNode)
	}

	data, err := json.Marshal(struct {
		Name     string     `json:"name"`
		Path     string     `json:"path"`
		Type     string     `json:"type"`
		Headers  ReqHeaders `json:"headers"`
		Children []*Node    `json:"children,omitempty"`
	}{
		Name:     node.Name,
		Path:     resourcePath,
		Type:     node.Type,
		Headers:  node.Headers,
		Children: nodes,
	})

	if err != nil {
		return nil, err
	}
	return data, nil
}
