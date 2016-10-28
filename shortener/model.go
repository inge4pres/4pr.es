package shortener

import (
	"math/rand"
	"time"
)

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var urllength = uint(8)
var domain = "4pr.es"
var dyndbtable = "4pres_url"

type ShortUrl struct {
	Redirect string `json:"redirect"`
	Created  string `json:"created"`
}

func shorten(c uint) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, c)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func NewShortUrl(url string) *ShortUrl {
	return &ShortUrl{
		Redirect: url,
		Created:  time.Now().Format(time.RFC1123Z),
	}
}

func SetShortUrlLength(l uint) {
	urllength = l
}

func SetDomain(d string) {
	domain = d
}

func SetDyndbTable(t string) {
	dyndbtable = t
}
