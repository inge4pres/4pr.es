package shortener

import (
	"encoding/json"
	"net/url"
	"os"
	"testing"
)

var longurl = "https%3A%2F%2Fgist.github.com%2Finge4pres%2F597bb9350ff3e9cc43ecb476a10e636b"
var test = domain + "/" + "abcdefg12"
var testenc = url.QueryEscape(test)

func TestNewShortUrl(t *testing.T) {
	surl := NewShortUrl(longurl)
	if err := json.NewEncoder(os.Stdout).Encode(&surl); err != nil {
		t.Error(err)
	}
}
