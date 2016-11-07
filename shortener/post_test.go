package shortener

import (
	"testing"
)

func TestSaveShortUrl(t *testing.T) {
	err := SaveShortUrl(longurl, dyndbtable)
	if err != nil {
		t.Error(err)
	}
}
