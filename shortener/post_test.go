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
	ret, err := urlExists(test, dyndbtable)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Url Exists? ", ret)
}
