package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/rayhaanbhikha/localstack-s3-ui/s3"
	"github.com/rayhaanbhikha/localstack-s3-ui/utils"
)

func startServer(rootNode *s3.Node) {

	r := mux.NewRouter()
	r.PathPrefix("/api/resource/").
		Handler(http.StripPrefix("/api/resource", resourcesHandler(rootNode))).
		Methods("GET")

	spa := spaHandler{staticPath: "build", indexPath: "index.html", resourceHandler: resourceHandler(rootNode)}
	r.PathPrefix("/").Handler(spa)

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
