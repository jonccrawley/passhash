package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/jonccrawley/passhash/backend"
	"github.com/jonccrawley/passhash/definition"
	"github.com/jonccrawley/passhash/handler"
	"github.com/jonccrawley/passhash/repository"
)

var (
	Port = flag.String("port", ":8080", "Server Port")
)

func main() {

	log.Printf("Starting server on port %v", *Port)

	log.Printf("Starting the Dispatcher with 5 workers.")
	backend.StartDispatcher(5)

	definition.HashRepo = &repository.HashRepository{}
	definition.ExecutionRepo = &repository.ExecutionRepository{}
	definition.StatisticsRepo = &repository.StatisticsRepository{}

	m := http.NewServeMux()
	s := http.Server{Addr: *Port  , Handler: m}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	m.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))

		// Cancel the context on request
		cancel()
	})

	m.HandleFunc("/hash", handler.SaveHashHandler)
	m.HandleFunc("/hash/", handler.RetrieveHashHandler)
	m.HandleFunc("/stats", handler.StatisticsHandler)

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	select {
	case <-ctx.Done():
		// Shutdown the server when the context is canceled
		s.Shutdown(ctx)
	}
	log.Printf("Finished")
}