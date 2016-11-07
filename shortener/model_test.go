package shortener

import (
	"encoding/json"
	"os"
	"testing"
)

var longurl = "https://gist.github.com/inge4pres/597bb9350ff3e9cc43ecb476a10e636b"
var test = domain + "/" + "abcdefg12"

func TestNewShortUrl(t *testing.T) {
	surl := NewShortUrl(longurl)
	if err := json.NewEncoder(os.Stdout).Encode(&surl); err != nil {
		t.Error(err)
	}
}
