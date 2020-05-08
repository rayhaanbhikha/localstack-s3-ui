package main

import (
	"encoding/base64"
	"log"
	"net/http"
	"path"

	"github.com/rayhaanbhikha/localstack-s3-ui/s3"
)

func resourceHandler(rootNode *s3.Node) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resourcePath := path.Clean(r.URL.EscapedPath())
		log.Println("Resource: ", resourcePath)
		node, ok := rootNode.Get(resourcePath)
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
		json, err := rootNode.JSON(resourcePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	})
}
