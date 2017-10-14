package shortener

import (
	"testing"
)

func TestDeleteShortUrl(t *testing.T) {
	if err := DeleteShortUrl(testenc, dyndbtable); err != nil {
		t.Error(err)
	}
	return
}
