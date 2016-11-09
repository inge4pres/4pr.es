package main

import (
	"encoding/json"
	"log"

	"github.com/eawsy/aws-lambda-go/service/lambda/runtime"
	"github.com/inge4pres/4pr.es/shortener"
)

var domain = "4pr.es"

func handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	log.Printf("Received event: %s\n", string(evt))
	var values map[string]string
	if err := json.Unmarshal(evt, &values); err != nil {
		return nil, err
	}
	resp, err := shortener.ReadShortUrl(domain+"/"+values["url"], shortener.GetDyndbTable())
	if err != nil {
		return nil, err
	}
	surl := new(shortener.ShortUrl)
	err = json.Unmarshal(resp, &surl)
	if err != nil {
		return nil, err
	}
	return surl.Redirect, nil
}

func init() {
	runtime.HandleFunc(handle)
}

func main() {}
