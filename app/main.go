package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/lashn/URL-Shortener/app/routes"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/shortenURL", routes.ShortenURL).Methods("POST")

	fmt.Println("server is running on 8000 port")
	log.Fatal(http.ListenAndServe(":8000", router))
}
