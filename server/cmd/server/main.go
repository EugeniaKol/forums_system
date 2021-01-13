package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/EugeniaKol/forums_system/server/db"
	"github.com/joho/godotenv"
)

// Define flag to set port number with console
var httpPortNumber = flag.Int("p", 8080, "HTTP port number")

// NewDbConnection creates connection to db
func NewDbConnection() (*sql.DB, error) {
	connection := &db.Connection{
		DbName:     os.Getenv("DbName"),
		User:       os.Getenv("User"),
		Password:   os.Getenv("Password"),
		Host:       os.Getenv("Host"),
		DisableSSL: true,
	}
	return connection.Open()
}

func main() {
	// Load .env
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Parse command line arguments. Port number can be defined with "-p" flag
	flag.Parse()

	// Create server
	if server, err := ComposeAPIServer(HTTPPortNumber(*httpPortNumber)); err == nil {
		// Start server
		go func() {
			log.Println("Starting forums server...")

			err := server.Start()
			if err == http.ErrServerClosed {
				log.Printf("HTTP server stopped")
			} else {
				log.Fatalf("Cannot start HTTP server: %s", err)
			}
		}()

		// Wait for Ctrl-C signal
		sigChannel := make(chan os.Signal, 1)
		signal.Notify(sigChannel, os.Interrupt)

		// Recieve value from sigChannel
		<-sigChannel

		if err := server.Stop(); err != nil && err != http.ErrServerClosed {
			log.Printf("Error stopping the server: %s", err)
		}
	} else {
		log.Fatalf("Cannot initialize forums server: %s", err)
	}
}
