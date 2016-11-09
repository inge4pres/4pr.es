package shortener

import (
	"fmt"
	"testing"
)

func TestSaveShortUrl(t *testing.T) {
	err := SaveShortUrl(longurl, dyndbtable)
	if err != nil {
		t.Error(err)
	}
}

func TestUrlExists(t *testing.T) {
	ret, err := UrlExists(test, dyndbtable)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Url Exists? ", ret)
}
