package main

import (
	"encoding/json"
	"log"

	"github.com/eawsy/aws-lambda-go/service/lambda/runtime"
	_ "github.com/inge4pres/4pr.es/shortener"
)

var domain = "4pr.es"

func handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	log.Printf("Received event: %s\n", string(evt))
	//	var values map[string]string
	//	if err := json.Unmarshal(evt, &values); err != nil {
	//		return nil, err
	//	}
	//	redirect, err := shortener.ReadShortUrl(domain+"/"+ctx., shortener.GetDyndbTable())
	//TODO get data from Path in URL, passing it from ctx?
	return string(evt), nil
}

func init() {
	runtime.HandleFunc(handle)
}

func main() {}
