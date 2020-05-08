package s3

import "fmt"

type Node struct {
	Name       string `json:"name"`
	bucketName string
	Path       string     `json:"path"`
	Type       string     `json:"type"`
	Data       string     `json:"-"`
	Headers    ReqHeaders `json:"h"`
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
