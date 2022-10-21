package helper

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/lashn/URL-Shortener/app/util"
)

// saves the newly created shorten url to text file
func SaveMapTo_TextFile(url, enc_url string) error {
	file, err := os.OpenFile(util.DS_Path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Error opening the file", err)
		return err
	}
	defer file.Close()

	//add new URL and the short url to the text file
	fileWriter := bufio.NewWriter(file)
	newURL := (url + "::" + enc_url)
	fileWriter.WriteString(newURL + "\n")

	err = fileWriter.Flush()
	if err != nil {
		log.Fatalln("Error writing the URL to the file", err)
		return err
	}
	return nil
}

// checks if the given url already exists in the text file
func CheckURL_TextFile(url string) (string, bool) {
	file, err := os.Open(util.DS_Path)
	if err != nil {
		log.Fatalln("Error opening the file", err)
	}
	defer file.Close()

	//generalize url
	var GeneralizedUrl = url
	if strings.HasPrefix(url, "https://") {
		GeneralizedUrl = strings.Replace(url, "https://", "", 1)
	}
	if strings.HasPrefix(url, "http://") {
		GeneralizedUrl = strings.Replace(url, "http://", "", 1)
	}

	//check if the URL exists in the text file by file walk and return the shortURL or null
	found := "null"
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		if strings.Contains(fileScanner.Text(), GeneralizedUrl) {
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
