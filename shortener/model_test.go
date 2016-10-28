package shortener

import (
	"encoding/json"
	"os"
	"testing"
)

var longurl = "http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_CreateTable.html"
var test = domain + "/" + "abcdefg12"

func TestNewShortUrl(t *testing.T) {
	surl := NewShortUrl(longurl)
	if err := json.NewEncoder(os.Stdout).Encode(&surl); err != nil {
		t.Error(err)
	}
}
