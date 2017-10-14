package shortener

import (
	"encoding/json"
	"net/url"
	"os"
	"testing"
)

var (
	malicious    = "http://ing-diba.de-sicherheitsverifikation.com/script.php?fid=1234234"
	notmalicious = "https://answers.splunk.com/answers/46774/where-can-i-change-system-variable-from-splunk-db.html"
	longurl      = "https%3A%2F%2Fgist.github.com%2Finge4pres%2F597bb9350ff3e9cc43ecb476a10e636b"
	test         = domain + "/" + "abcdefg12"
	test2        = domain + "/" + "llwp48"
	testenc      = url.QueryEscape(test)
	testenc2     = url.QueryEscape(test2)
)

func TestNewShortUrl(t *testing.T) {
	surl := NewShortUrl(longurl)
	if err := json.NewEncoder(os.Stdout).Encode(&surl); err != nil {
		t.Error(err)
	}
}
