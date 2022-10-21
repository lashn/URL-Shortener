package routes

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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

	//validator for URL - raise error if bad URL
	if !helper.IsValidURL(url) {
		writer.WriteHeader((http.StatusBadRequest))
		fmt.Fprint(writer, "{\"message\":\"Bad URL\"}")
		return
	}

	if val, ok := checkURL_TextFile(url); ok {
		json.NewEncoder(writer).Encode(host + val)
	} else {
		enc_url := encode(url)
		json.NewEncoder(writer).Encode(host + enc_url)
	}
}

func encode(url string) string {
	//encoding with a unique url generated using uuid
	enc_url := uuid.New().String()[:6]
	saveMapTo_TextFile(url, enc_url)
	return enc_url
}

func saveMapTo_TextFile(url, enc_url string) {
	file, err := os.OpenFile(ds_path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Error opening the file", err)
	}
	defer file.Close()

	//add new URL and the short url to the text file
	fileWriter := bufio.NewWriter(file)
	newURL := (url + "::" + enc_url)
	fileWriter.WriteString(newURL + "\n")

	err = fileWriter.Flush()
	if err != nil {
		log.Fatalln("Error writing the URL to the file", err)
	}
}

func checkURL_TextFile(url string) (string, bool) {
	file, err := os.Open(ds_path)
	if err != nil {
		log.Fatalln("Error opening the file", err)
	}
	defer file.Close()

	//check if the URL exists in the text file by file walk and return the shortURL or null
	found := "null"
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		if strings.Contains(fileScanner.Text(), url) {
			found = fileScanner.Text()
		}
	}

	if found != "null" {
		result := (strings.Split(found, "::"))
		Enc_url := result[1]
		return Enc_url, true
	} else {
		return found, false
	}
}
