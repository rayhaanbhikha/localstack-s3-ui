package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"path"

	"github.com/rayhaanbhikha/localstack-s3-ui/s3"
)

func resourceHandler(rootNode *s3.Node) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resourcePath := path.Clean(r.URL.EscapedPath())
		log.Println("Resource: ", resourcePath)
		node, ok := rootNode.GetNode(resourcePath)
		// TODO: node may not be a child node.
		if !ok {
			http.NotFound(w, r)
			return
		}
		decoded, err := base64.StdEncoding.DecodeString(node.Data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", node.Headers.ContentType)
		w.Write([]byte(decoded))
	})
}

func resourcesHandler(rootNode *s3.Node) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resourcePath := path.Clean(r.URL.EscapedPath())
		log.Println("Resources: ", resourcePath)
		node, ok := rootNode.GetNode(resourcePath)

		if !ok {
			http.NotFound(w, r)
			return
		}

		nodes := make([]*s3.Node, 0)
		for _, childNode := range node.Children {
			nodes = append(nodes, childNode)
		}

		data, err := json.Marshal(struct {
			Name     string        `json:"name"`
			Path     string        `json:"path"`
			Type     string        `json:"type"`
			Headers  s3.ReqHeaders `json:"headers"`
			Children []*s3.Node    `json:"children,omitempty"`
		}{
			Name:     node.Name,
			Path:     resourcePath,
			Type:     node.Type,
			Headers:  node.Headers,
			Children: nodes,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})
}
