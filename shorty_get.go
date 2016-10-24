package 4pres

import (
	"log"
	"net/http"
	
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	db "github.com/aws/aws-sdk-go/service/dynamodb"
)

func connectAws()  error {
	sess, err := session.NewSession()
	if err != nil {
		log.Println("failed to create session,", err)
		return err
	}
	return nil
}

func readUrl() {}


//	svc := dynamodb.New(sess)
