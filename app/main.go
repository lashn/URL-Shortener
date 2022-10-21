package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lashn/URL-Shortener/app/routes"
)

func main() {
	router := NewApplicationHandler()
	log.Fatal(http.ListenAndServe(":8000", router))
}

// NewApplicationHandler provides a configured handler for the
// application.
func NewApplicationHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/shortenURL", routes.ShortenURL).Methods("POST")
	router.HandleFunc("/test", Test)
	return router
}

func Test(writer http.ResponseWriter, reader *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
    writer.WriteHeader(http.StatusOK)

    // In the future we could report back on the status of our DB, or our cache
    // (e.g. Redis) by performing a simple PING, and include them in the response.
    io.WriteString(writer, `{"alive": true}`)
}
