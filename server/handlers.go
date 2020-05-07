package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"path"

	"github.com/rayhaanbhikha/localstack-s3-ui/s3"
)

type contextKey string

func (c contextKey) String() string {
	return "mypackage context key " + string(c)
}

var resourcePathCtxKey = contextKey("path")

func queryPathMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		queryParams := r.URL.Query()
		if v, ok := queryParams["path"]; ok {
			path := path.Clean(v[0])
			fmt.Println("Resource Path requested: ", path)
			ctx := context.WithValue(r.Context(), resourcePathCtxKey, path)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.NotFound(w, r)
		}
	})
}

func pageHandler(rootNode *s3.Node) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resourcePath := r.Context().Value(resourcePathCtxKey).(string)
		node, ok := rootNode.Get(resourcePath)
		if !ok {
			http.NotFound(w, r)
			return
		}
		decoded, err := base64.StdEncoding.DecodeString(node.Data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// TODO: node should have content type. (chrome is smart enough to know what the mime type is.)
		// w.Header().Set("Content-Type", "text/javascript")
		w.Write([]byte(decoded))
	})
}

func resourceHandler(rootNode *s3.Node) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resourcePath := r.Context().Value(resourcePathCtxKey).(string)
		json, err := rootNode.JSON(resourcePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	})
}
