package s3

import (
	"fmt"
	"log"
)

type S3Node struct {
	Name     string `json:"name"`
	Path     string
	Type     string
	Data     string
	children []*S3Node
}

func (n *S3Node) Print() {
	fmt.Println(fmt.Sprintf(`
	Name: %s
	Type: %s
	Path: %s
`, n.Name, n.Type, n.Path))
	if len(n.children) > 0 {
		for _, childNode := range n.children {
			childNode.Print()
		}
	}
}

func (n *S3Node) addNode(resource *S3Resource) {
	resource.Path = resource.Path[1:]
	if len(resource.Path) == 1 {
		// will add at this level
		// will always be a file.
		fileNode := &S3Node{
			Name: resource.Name,
			Path: fmt.Sprintf("%s/%s", n.Path, resource.Path[0]),
			Type: "File",
		}
		n.children = append(n.children, fileNode)
		return
	}
	for i := 0; i < len(n.children); i++ {
		name := n.children[i].Name
		if name == resource.Path[0] {
			n.children[i].addNode(resource)
			return
		}
	}
	// definitely a nested resource.
	// create file node.
	dirNode := &S3Node{
		Name: resource.Path[0],
		Path: fmt.Sprintf("%s/%s", n.Path, resource.Path[0]),
		Type: "Directory",
	}
	dirNode.addNode(resource)
	n.children = append(n.children, dirNode)
}

func M() {

	s3Resources, err := Parse("./recorded_api_calls.mock.json")
	if err != nil {
		log.Fatal(err)
	}

	// initial node.
	s3Node := &S3Node{Name: "static-resources", Path: "static-resources"}

	for _, s3Resource := range s3Resources {
		if s3Resource.Type != "Bucket" {
			s3Node.addNode(s3Resource)
		}
	}

	s3Node.Print()
}
