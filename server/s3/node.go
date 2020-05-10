package s3

import "fmt"

// Node ... creates s3 node
type Node struct {
	Name       string `json:"name"`
	bucketName string
	Path       string           `json:"path"`
	ResPath    string           `json:"resourcePath"`
	Type       string           `json:"type"`
	Data       string           `json:"-"`
	Headers    ReqHeaders       `json:"h"`
	Children   map[string]*Node `json:"-"`
}

// RootNode ... return root s3 node.
func RootNode() *Node {
	return &Node{Name: "Root", Path: "/", Type: "Root", Children: make(map[string]*Node)}
}

func (n *Node) String() string {
	return fmt.Sprintf(`
	Name: %s
	Type: %s
	Path: %s
	Data: %s
`, n.Name, n.Type, n.Path, n.Data)
}

// Print ... node tree.
func (n *Node) Print() {
	fmt.Println(n)
	if len(n.Children) > 0 {
		for _, childNode := range n.Children {
			childNode.Print()
		}
	}
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

// GetResourceNode ... returns resource node at given path.
func (n *Node) GetResourceNode(resourcePath string) (*Node, bool) {
	if n.ResPath == resourcePath {
		return n, true
	}

	for _, childNode := range n.Children {
		if v, ok := childNode.GetResourceNode(resourcePath); ok {
			return v, ok
		}
	}
	return nil, false
}

// GetNode ... returns node at given path.
func (n *Node) GetNode(path string) (*Node, bool) {
	if n.Path == path {
		return n, true
	}

	for _, childNode := range n.Children {
		if v, ok := childNode.GetNode(path); ok {
			return v, ok
		}
	}
	return nil, false
}

func (n *Node) deleteNode(s3Request *apiRequest) {

	if s3Request.Path == "/" {
		return
	}

	pathLen := len(s3Request.actualPath) - 1
	keyToDel := s3Request.actualPath[pathLen]

	if _, ok := n.Children[keyToDel]; ok {
		delete(n.Children, keyToDel)
		return
	}

	for key := range n.Children {
		n.Children[key].deleteNode(s3Request)
	}
}
