package s3

func (n *Node) deleteNode(s3Request *apiRequest) {

	if s3Request.Path == "/" {
		return
	}

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
