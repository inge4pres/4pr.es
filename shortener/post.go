package shortener

import (
	"errors"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	db "github.com/aws/aws-sdk-go/service/dynamodb"
	sb "github.com/google/safebrowsing"
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
	if malicious := checkMalicousURL(decoded); malicious != nil {
		return ret, malicious
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

func checkMalicousURL(url string) error {
	sbdb := filepath.Join(filepath.Dir("/"), "tmp", "safebrowsing.db")
	k := os.Getenv("GOOGLE_SB_KEY")
	if k == "" {
		log.Println("No Google API credentials found! Exiting")
		return errors.New("Cannot authenticate to Safebrowsing API")
	}
	if _, err := os.Open(sbdb); err != nil {
		_, err := os.Create(sbdb)
		if err != nil {
			return err
		}
	}

	conf := &sb.Config{
		//Lambda offers a 512MB filesystem on /tmp
		DBPath: sbdb,
		APIKey: k,
		Logger: os.Stderr,
	}
	client, err := sb.NewSafeBrowser(*conf)
	if err != nil {
		return err
	}
	defer client.Close()
	threats, err := client.LookupURLs([]string{url})
	if err != nil {
		return err
	}
	if len(threats[0]) > 0 {
		log.Printf("Detected threats for URL %s: %v", url, threats)
		return errors.New("HTTP 412 Precondition Failed")
	}
	return nil
}
