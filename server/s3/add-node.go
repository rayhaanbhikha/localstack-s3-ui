package s3

import (
	"fmt"
)

func (n *Node) addNode(s3Request *apiRequest) {
	if n.Name == "Root" && len(s3Request.actualPath) == 1 {
		bucketName := s3Request.actualPath[0]
		if _, ok := n.Children[bucketName]; !ok {
			bucketNode := &Node{
				Name:       bucketName,
				bucketName: bucketName,
				Path:       fmt.Sprintf("/%s", bucketName),
				Type:       "Bucket",
				Children:   make(map[string]*Node),
			}
			n.Children[bucketName] = bucketNode
		}
		return
	}

	if n.Name != "Root" && len(s3Request.actualPath) == 1 {
		fileName := s3Request.actualPath[0]
		if _, ok := n.Children[fileName]; !ok {
			fileNode := &Node{
				bucketName: n.bucketName,
				Name:       fileName,
				Type:       "File",
				Path:       fmt.Sprintf("%s/%s", n.Path, fileName),
				ResPath:    fmt.Sprintf("%s/%s", n.ResPath, fileName),
				Data:       s3Request.Data,
				Headers:    s3Request.Headers,
				Children:   make(map[string]*Node),
			}
			n.Children[fileName] = fileNode
		} else {
			n.Children[fileName].Data = s3Request.Data
		}
		return
	}

	for childPath, childNode := range n.Children {
		if childNode.Name == s3Request.actualPath[0] {
			s3Request.actualPath = s3Request.actualPath[1:]
			n.Children[childPath].addNode(s3Request)
			return
		}
	}

	// definitely a nested resource.
	// create file node.
	dirName := s3Request.actualPath[0]
	var path string
	var resPath string
	if n.Path == "/" {
		path = fmt.Sprintf("/%s", dirName)
		resPath = fmt.Sprintf("/%s", dirName)
	} else {
		path = fmt.Sprintf("%s/%s", n.Path, dirName)
		resPath = fmt.Sprintf("%s/%s", n.ResPath, dirName)
	}
	dirNode := &Node{
		Name:     dirName,
		Path:     path,
		ResPath:  resPath,
		Type:     "Directory",
		Children: make(map[string]*Node),
	}
	n.Children[dirName] = dirNode
	s3Request.actualPath = s3Request.actualPath[1:]
	dirNode.addNode(s3Request)
}
