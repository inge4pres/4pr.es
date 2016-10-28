package shortener

import (
	"log"
	"testing"
)

func TestReadShortUrl(t *testing.T) {
	redir, err := readShortUrl(test, dyndbtable)
	if err != nil {
		t.Error(err)
		return
	}
	log.Printf("Read url %s\nfor short url %s\n", redir, test)
}
