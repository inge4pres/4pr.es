package shortener

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	db "github.com/aws/aws-sdk-go/service/dynamodb"
)

func SaveShortUrl(url, table string) (string, error) {
	if err := connect(); err != nil {
		log.Fatalf("Dynamo DB connection: %v")
	}
	obj := NewShortUrl(url)
	surl := GetDomain() + "/" + shorten(urllength)
	for c, e := urlExists(surl, GetDyndbTable()); c; {
		if e != nil {
			log.Printf("Url exists err: %v\n", e)
		}
		surl = GetDomain() + "/" + shorten(urllength)
	}
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
	return surl, err
}

func urlExists(url, table string) (bool, error) {
	out, err := svc.Query(&db.QueryInput{
		TableName: aws.String(table),
		ExpressionAttributeNames: map[string]*string{
			"#U": aws.String("url"),
		},
		ExpressionAttributeValues: map[string]*db.AttributeValue{
			":u": &db.AttributeValue{
				S: aws.String(url),
			},
		},
		KeyConditionExpression: aws.String("#U = :u"),
	})
	if len(out.Items) > 0 {
		return true, err
	}
	return false, nil
}
