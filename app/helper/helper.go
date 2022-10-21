package helper

import (
	"net/url"
)

func IsValidURL(verifyURL string) bool {
	_, err := url.ParseRequestURI(verifyURL)
	if err != nil {
		return false
	}
	return true
}

