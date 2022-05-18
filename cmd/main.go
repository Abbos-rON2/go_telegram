package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"telegram_back/config"
	"telegram_back/rest"
	"telegram_back/storage"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var cfg = config.Load()

	// mongo
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://admin:password@localhost:27017"))
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(context.TODO())

	storage := storage.New(cfg, client.Database("telegram_back"))
	srv := rest.New(cfg, storage)

	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Printf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
