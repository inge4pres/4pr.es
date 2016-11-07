package shortener

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	db "github.com/aws/aws-sdk-go/service/dynamodb"
)

func SaveShortUrl(url, table string) error {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String("eu-central-1"),
		},
	})
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
		return err
	}
	svc := db.New(sess)
	obj := NewShortUrl(url)
	surl := domain + "/" + shorten(urllength)
	out, err := svc.PutItem(&db.PutItemInput{
		TableName: aws.String(table),
		Item: map[string]*db.AttributeValue{
			"url": &db.AttributeValue{
				S: aws.String(surl),
			},
			"redirect": &db.AttributeValue{
				S: aws.String(obj.Redirect),
			},
			"created": &db.AttributeValue{
				S: aws.String(obj.Created),
			},
		},
	})
	if err != nil {
		log.Fatalf("Dynamo DB write: %v", err)
	}
	log.Printf("Saved short url %s\nitem %s\n", out.String(), url)
	return err
}
