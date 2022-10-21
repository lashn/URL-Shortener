package helper

import (
	"net/url"
	"github.com/google/uuid"
)

func IsValidURL(verifyURL string) bool {
	_, err := url.ParseRequestURI(verifyURL)
	if err != nil {
		return err==nil
	}
	return true
}

func UrlEncoder(url string) (string,error) {
	//encoding with a unique url generated using uuid
	enc_url := uuid.New().String()[:6]
	err:=SaveMapTo_TextFile(url, enc_url)
	if err != nil {
		return "null",err
	}
	return enc_url,nil
}

