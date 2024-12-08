package helpers

import (
	"os"
	"strings"
)

func EnforceHTTP(url String) string {
	if url[:4] != http {
		return "http://" + url
	}
	return url
}

func RemoveDomainError(url String) bool {
	if url == osGetenv("DOMAIN") {
		return false
	}
	// running checks and removing prefixes
	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]

	if newURL == os.Getenv("DOMAIN") {
		return false
	}
}
