package s3

import (
	"fmt"
)

// Node ...
type Node struct {
	Name       string `json:"name"`
	bucketName string
	Path       string `json:"path"`
	Type       string `json:"type"`
	Data       string `json:"-"`
	children   map[string]*Node
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
	if len(n.children) > 0 {
		for _, childNode := range n.children {
			childNode.Print()
		}
	}
}

func (n *Node) getNode(path string) (*Node, bool) {
	if n.Path == path {
		return n, true
	}

	for _, childNode := range n.children {
		if v, ok := childNode.getNode(path); ok {
			return v, ok
		}
	}
	return nil, false
}

func (n *Node) deleteNode(s3Request *apiRequest) {
	pathLen := len(s3Request.actualPath) - 1
	keyToDel := s3Request.actualPath[pathLen]

	if _, ok := n.children[keyToDel]; ok {
		delete(n.children, keyToDel)
		return
	}

	for key := range n.children {
		n.children[key].deleteNode(s3Request)
	}
}

func (n *Node) addNode(s3Request *apiRequest) {
	if n.Name == "Root" && len(s3Request.actualPath) == 1 {
		bucketName := s3Request.actualPath[0]
		if _, ok := n.children[bucketName]; !ok {
			bucketNode := &Node{
				Name:       bucketName,
				bucketName: bucketName,
				Path:       fmt.Sprintf("/%s", bucketName),
				Type:       "Bucket",
				children:   make(map[string]*Node),
			}
			n.children[bucketName] = bucketNode
		}
		return
	}

	if n.Name != "Root" && len(s3Request.actualPath) == 1 {
		fileName := s3Request.actualPath[0]
		if _, ok := n.children[fileName]; !ok {
			fileNode := &Node{
				bucketName: n.bucketName,
				Name:       fileName,
				Type:       "File",
				Path:       fmt.Sprintf("%s/%s", n.Path, fileName),
				Data:       s3Request.Data,
				children:   make(map[string]*Node),
			}
			n.children[fileName] = fileNode
		} else {
			n.children[fileName].Data = s3Request.Data
		}
		return
	}

	for childPath, childNode := range n.children {
		if childNode.Name == s3Request.actualPath[0] {
			s3Request.actualPath = s3Request.actualPath[1:]
			n.children[childPath].addNode(s3Request)
			return
		}
	}

	// definitely a nested resource.
	// create file node.
	dirName := s3Request.actualPath[0]
	var path string
	if n.Path == "/" {
		path = fmt.Sprintf("/%s", dirName)
	} else {
		path = fmt.Sprintf("%s/%s", n.Path, dirName)
	}
	dirNode := &Node{
		Name:     dirName,
		Path:     path,
		Type:     "Directory",
		children: make(map[string]*Node),
	}
	n.children[dirName] = dirNode
	s3Request.actualPath = s3Request.actualPath[1:]
	dirNode.addNode(s3Request)
}
