package helpers

import (
	"os"
	"strings"
)

func EnforceHTTP() string {
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}

func RemoveDomainError(url string) bool {
	if url == os.Getenv("DOMAIN") {
		return false
	}

	neURL := strings.Replace(url, "http://", "", 1) /* remove the commonly found prefixes, to ensure there;s no way user can hack into our system */
	newURL = strings.Replace(newURL, "https://", "", 1)
	neURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]

	if newURL == os.Getenv("DOMAIN") {
		return false
	}

	return true
}
