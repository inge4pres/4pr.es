package main

import (
	"encoding/json"
	"log"

	"github.com/eawsy/aws-lambda-go/service/lambda/runtime"
	"github.com/inge4pres/4pr.es/shortener"
)

var domain = "4pr.es"

func handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	log.Println("Received event: ", string(evt))
	var values map[string]string
	if err := json.Unmarshal(evt, &values); err != nil {
		return nil, err
	}
	resp, err := shortener.SaveShortUrl(values["longurl"], shortener.GetDyndbTable())
	if err != nil {
		log.Println("Save Short Url err: ", err)
		return nil, err
	}
	return nil, nil
}

func init() {
	runtime.HandleFunc(handle)
}

func main() {}
