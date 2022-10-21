package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/lashn/URL-Shortener/app/helper"
	"github.com/lashn/URL-Shortener/app/util"
)

func ShortenURL(writer http.ResponseWriter, reader *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	urlreq := new(util.Request)

	err := json.NewDecoder(reader.Body).Decode(&urlreq)
	if err != nil {
		log.Fatalln("Error decoding the URL", err)
	}
	url := urlreq.URL

	//validator for URL - raise error if bad URL
	if !helper.IsValidURL(url) {
		writer.WriteHeader((http.StatusBadRequest))
		fmt.Fprint(writer, "{\"message\":\"Bad URL\"}")
		return
		
	}

	if enc_url, ok := helper.CheckURL_TextFile(url); ok {
		json.NewEncoder(writer).Encode(util.Host + enc_url)
	} else {
		enc_url,err := helper.UrlEncoder(url)
		if err != nil {
			log.Fatalln("Error encoding the URL", err)
		}
		json.NewEncoder(writer).Encode(util.Host + enc_url)
	}
}


