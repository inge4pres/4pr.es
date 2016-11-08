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
	var values map[string]interface{}
	if err := json.Unmarshal(evt, &values); err != nil {
		return nil, err
	}
	redirect, err := shortener.ReadShortUrl(domain+values["path"].(string), shortener.GetDyndbTable())
	if err != nil {
		return nil, err
	}
	//TODO get data from Path in URL, passing it from ctx?
	// whenusing Proxy Integration with Lambda, "path" is passwd alongside the payload
	return redirect, nil
}

func init() {
	runtime.HandleFunc(handle)
}

func main() {}
