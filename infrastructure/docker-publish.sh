#!/bin/bash 
# build and publish on AWS

case "$1" in {
 "get"
}
#GET
docker run --rm -v ///c/Users/francesco.gualazzi/Documents/Software/Golang:/go -v ///c/Users/francesco.gualazzi/Documents/Software/Golang/src/github.com/inge4pres/4pr.es/cmd/4pres_get:/tmp eawsy/aws-lambda-go 
aws lambda update-function-code  --function-name 4pres_lambda_get   --zip-file fileb://handler.zip --publish

#POST
docker run --rm -v ///c/Users/francesco.gualazzi/Documents/Software/Golang:/go -v ///c/Users/francesco.gualazzi/Documents/Software/Golang/src/github.com/inge4pres/4pr.es/cmd/4pres_post:/tmp eawsy/aws-lambda-go
aws lambda update-function-code  --function-name 4pres_lambda_post   --zip-file fileb://handler.zip --publish 
