package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/google/uuid"
	"github.com/lashn/URL-Shortener/app/helper"
)

func ShortenURL(writer http.ResponseWriter, reader *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	urlreq := new(request)

	err := json.NewDecoder(reader.Body).Decode(&urlreq)
	if err != nil {
		log.Fatalln("Error decoding the URL", err)
	}
	url := urlreq.URL
	fmt.Println(url)

	//validator for URL - raise error if bad URL
	if !helper.IsValidURL(url) {
		writer.WriteHeader((http.StatusBadRequest))
		fmt.Fprint(writer, "{\"message\":\"Bad URL\"}")
		return
	}

	//check datastore and return same short url if exists
	if val, ok := url_map[url]; ok {
		json.NewEncoder(writer).Encode(host+val)
	} else {
		enc_url := encode(url)
		json.NewEncoder(writer).Encode(host+enc_url)
	}
}

func encode(url string) string {
	//encoding with a unique url generated using uuid
	enc_url := uuid.New().String()[:6]
	url_map[url] = enc_url
	return enc_url
}
