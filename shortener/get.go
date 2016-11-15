package shortener

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	db "github.com/aws/aws-sdk-go/service/dynamodb"
)

func ReadShortUrl(surl, tname string) ([]byte, error) {
	if err := connect(); err != nil {
		log.Fatalf("Dynamo DB connection: %v")
	}
	out, err := svc.GetItem(&db.GetItemInput{
		ConsistentRead: aws.Bool(false),
		Key: map[string]*db.AttributeValue{
			"url": &db.AttributeValue{
				S: aws.String(surl),
			},
		},
		TableName: aws.String(tname),
	})
	if err != nil {
		log.Printf("Error getting item on DB for URL %s\n%v", surl, err)
		return nil, err
	}
	// Prevent empty
	if len(out.Item) < 1 {
		return nil, errors.New("Not Found")
	}
	return json.Marshal(&ShortUrl{
		Redirect: *out.Item["redirect"].S,
		Created:  *out.Item["created"].S,
	})
}
