package shortener

import (
	"log"
	"math/rand"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	db "github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	letters    = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	urllength  = uint(8)
	domain     = "4pr.es"
	dyndbtable = "4pres_url"
	svc        *db.DynamoDB
	region     = "eu-central-1"
)

func init() {
	if err := connect(); err != nil {
		log.Fatalf("Dynamo DB connection: %v")
	}
}

type ShortUrl struct {
	Redirect string `json:"redirect"`
	Created  string `json:"created"`
}

type PostResp struct {
	Err error
	Url string
}

func shorten(c uint) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, c)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func NewShortUrl(url string) *ShortUrl {
	return &ShortUrl{
		Redirect: url,
		Created:  time.Now().Format(time.RFC1123Z),
	}
}

func SetShortUrlLength(l uint) {
	urllength = l
}

func SetDomain(d string) {
	domain = d
}

func SetDyndbTable(t string) {
	dyndbtable = t
}

func SetAwsRegion(region string) {
	region = region
}

func GetShortUrlLength() uint {
	return urllength
}

func GetDomain() string {
	return domain
}

func GetDyndbTable() string {
	return dyndbtable
}

func GetAwsRegion() string {
	return region
}

func connect() error {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(region),
		},
	})
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
		return err
	}
	svc = db.New(sess)
	return nil
}
