package shortener

import (
	"errors"
	"log"
	"net/url"

	"github.com/aws/aws-sdk-go/aws"
	db "github.com/aws/aws-sdk-go/service/dynamodb"
)

func SaveShortUrl(encurl, table string) (string, error) {
	var ret string
	if err := connect(); err != nil {
		log.Fatalf("Dynamo DB connection: %v")
	}
	decoded, err := url.QueryUnescape(encurl)
	if err != nil {
		log.Println("Decode URL err: ", err)
		return ret, errors.New("HTTP 500 Internal Server Error")
	}
	obj := NewShortUrl(decoded)
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
		return ret, errors.New("HTTP 500 Internal Server Error")
	}
	log.Printf("Saved short url %s\nitem %s\n", out.String(), encurl)
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
