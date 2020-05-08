package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"github.com/rayhaanbhikha/localstack-s3-ui/s3"
	"github.com/rayhaanbhikha/localstack-s3-ui/utils"
)

type spaHandler struct {
	resourceHandler http.Handler
	staticPath      string
	indexPath       string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// render index.html
	if path == "/" || path == "/s3" {
		http.ServeFile(w, r, "./build/index.html")
		return
	}

	path = filepath.Join(h.staticPath, path)

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// attempt to return s3 resource.
		h.resourceHandler.ServeHTTP(w, r)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func startServer(rootNode *s3.Node) {

	r := mux.NewRouter()
	r.PathPrefix("/api/resource/").
		Handler(http.StripPrefix("/api/resource", resourcesHandler(rootNode))).
		Methods("GET")

	spa := spaHandler{staticPath: "build", indexPath: "index.html", resourceHandler: resourceHandler(rootNode)}
	r.PathPrefix("/").Handler(spa)

	// TODO: use os.GetEnv to retieve PORT.
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", utils.GetPort()),
		Handler: r,
	}

	server.RegisterOnShutdown(func() {
		log.Println("Shutting down server")
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		log.Printf("Server starting on %s\n", server.Addr)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %s", err.Error())
		}
	}()

	sig := <-c
	log.Printf("Signal received: %s\n", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	err := server.Shutdown(ctx)

	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Err shutting down server: %s", err.Error())
	}
}
