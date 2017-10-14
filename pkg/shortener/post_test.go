package shortener

import (
	"fmt"
	"testing"
)

func TestSaveShortUrl(t *testing.T) {
	surl, err := SaveShortUrl(longurl, dyndbtable)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Long: %s \nShort: %s", longurl, surl)
}

func TestUrlExists(t *testing.T) {
	ret, err := urlExists(test2, dyndbtable)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Url Exists? ", ret)
}

func TestCheckMaliciousURL(t *testing.T) {
	if err := checkMalicousURL(malicious); err == nil {
		t.Error("Expected malicious URL detection")
	}
	if err := checkMalicousURL(notmalicious); err != nil {
		t.Error(err)
	}
}
