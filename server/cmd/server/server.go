package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/EugeniaKol/forums_system/server/forums"
)

// HTTPPortNumber alias for int
type HTTPPortNumber int

// ForumAPIServer configures necessary handlers and starts listening on a configured port
type ForumAPIServer struct {
	Port HTTPPortNumber

	ForumsHandler forums.HTTPHandlerFunc

	server *http.Server
}

// Start will set all handlers and start listening
// If this methods succeeds, it does not return until server is shut down
func (s *ForumAPIServer) Start() error {
	log.Printf("Listening...")

	if s.ForumsHandler == nil {
		return fmt.Errorf("Channel's HTTP handler is not defined")
	}
	if s.Port == 0 {
		return fmt.Errorf("Port is not defined")
	}

	// Create request router
	handler := new(http.ServeMux)
	handler.HandleFunc("/forums", s.ForumsHandler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

// Stop will shut down previously started HTTP server.
func (s *ForumAPIServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("Server was not started")
	}
	return s.server.Shutdown(context.Background())
}
