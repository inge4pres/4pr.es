package shortener

import (
	"errors"
	"log"
	"net/url"

	"github.com/aws/aws-sdk-go/aws"
	db "github.com/aws/aws-sdk-go/service/dynamodb"
)

func DeleteShortUrl(encurl, table string) error {
	if err := connect(); err != nil {
		log.Fatalf("Dynamo DB connection: %v", err)
	}
	decoded, err := url.QueryUnescape(encurl)
	if err != nil {
		log.Println("Decode URL err: ", err)
		return errors.New("HTTP 500 Internal Server Error")
	}
	//	obj := NewShortUrl(decoded)
	out, err := svc.DeleteItem(&db.DeleteItemInput{
		TableName: aws.String(table),
		Key: map[string]*db.AttributeValue{
			"url": &db.AttributeValue{
				S: aws.String(decoded),
			},
		},
	})
	if err != nil {
		log.Fatalf("Dynamo DB write: %v", err)
		return errors.New("HTTP 404 Not Found")
	}
	log.Printf("Deleted short url %s\nitem %s\n", out.String(), decoded)
	return nil
}
