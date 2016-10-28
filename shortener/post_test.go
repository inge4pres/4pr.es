package shortener

import (
	"testing"
)

func TestSaveShortUrl(t *testing.T) {
	err := saveShortUrl(longurl, dyndbtable)
	if err != nil {
		t.Error(err)
	}
}
