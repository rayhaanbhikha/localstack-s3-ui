package s3

import "fmt"

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
				Headers:    s3Request.Headers,
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
