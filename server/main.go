package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rayhaanbhikha/localstack-s3-ui/s3"
	"golang.org/x/net/context"
)

func main() {
	fileName := "./recorded_api_calls.json"
	rootNode := s3.RootNode()

	err := rootNode.LoadData(fileName)
	if err != nil {
		log.Fatal(err)
	}

	watcher, err := startFileWatcher(fileName, rootNode)
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	startServer(rootNode)
}

func startServer(rootNode *s3.Node) {

	mux := http.NewServeMux()
	mux.Handle("/resource", queryPathMiddleware(resourceHandler(rootNode)))
	mux.Handle("/page", queryPathMiddleware(pageHandler(rootNode)))

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	server.RegisterOnShutdown(func() {
		log.Println("Shutting down server")
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		log.Println("Server starting on PORT 8080")

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
