package s3

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
