package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/brilock/FirstService/handlers"
	"github.com/brilock/FirstService/version"
)

func main() {
	log.Printf(
		"Starting the service...\ncommit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release,
	)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	router := handlers.Router()
	//http.HandleFunc("/home", httpHandler)

	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func httpHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello! Your request was processed.")
}
