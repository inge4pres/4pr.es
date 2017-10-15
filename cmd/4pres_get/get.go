package main

import (
	"encoding/json"
	"log"

	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/inge4pres/4pr.es/pkg/shortener"
)

var domain = "4pr.es"

func Handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	log.Printf("Received event: %s\n", string(evt))
	var values map[string]string
	if err := json.Unmarshal(evt, &values); err != nil {
		return nil, err
	}
	resp, err := shortener.ReadShortUrl(domain+"/"+values["url"], shortener.GetDyndbTable())
	if err != nil {
		log.Printf("ReadShortUrl error: %v", err)
		return nil, err
	}
	surl := new(shortener.ShortUrl)
	err = json.Unmarshal(resp, surl)
	if err != nil {
		log.Printf("json Unmarshall: %v", err)
		return nil, err
	}
	return surl.Redirect, nil
}

func main() {}
