package s3

import (
	"fmt"
)

// Node ...
type Node struct {
	Name       string
	BucketName string
	Path       string
	Type       string
	Data       string
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

func (n *Node) addNode(path []string, data string) {
	if n.Name == "Root" && len(path) == 1 {
		bucketName := path[0]
		if _, ok := n.children[bucketName]; !ok {
			bucketNode := &Node{
				Name:       bucketName,
				BucketName: bucketName,
				Path:       fmt.Sprintf("/%s", bucketName),
				Type:       "Bucket",
				children:   make(map[string]*Node),
			}
			n.children[bucketName] = bucketNode
		}
		return
	}

	if n.Name != "Root" && len(path) == 1 {

		fileName := path[0]
		if _, ok := n.children[fileName]; !ok {
			fileNode := &Node{
				BucketName: n.BucketName,
				Name:       fileName,
				Type:       "File",
				Path:       fmt.Sprintf("%s/%s", n.Path, fileName),
				Data:       data,
				children:   make(map[string]*Node),
			}
			n.children[fileName] = fileNode
		} else {
			// just update the data.
			n.children[fileName].Data = data
		}
		return
	}

	for childPath, childNode := range n.children {
		if childNode.Name == path[0] {
			path = path[1:]
			n.children[childPath].addNode(path, data)
			return
		}
	}

	// definitely a nested resource.
	// create file node.
	dirName := path[0]
	dirNode := &Node{
		Name:     dirName,
		Path:     fmt.Sprintf("%s/%s", n.Path, dirName),
		Type:     "Directory",
		children: make(map[string]*Node),
	}
	n.children[dirName] = dirNode
	path = path[1:]
	dirNode.addNode(path, data)
}
