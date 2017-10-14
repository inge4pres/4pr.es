package delete

import (
	"encoding/json"
	"log"

	"github.com/eawsy/aws-lambda-go/service/lambda/runtime"
	"github.com/inge4pres/4pr.es/pkg/shortener"
)

func handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	log.Println("Received event: ", string(evt))
	var values map[string]string
	if err := json.Unmarshal(evt, &values); err != nil {
		return nil, err
	}
	if err := shortener.DeleteShortUrl(values["url"], shortener.GetDyndbTable()); err != nil {
		return nil, err
	}
	return "DELETED", nil
}

func init() {
	runtime.HandleFunc(handle)
}

func main() {}
