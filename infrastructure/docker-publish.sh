#!/bin/bash 
# build code and publish on AWS Lambda
# parameters 
#   1- build-type: get|post|all
#   2- godir: pass it with $GOPATH on *nix, with ///c/Users/user.name/go/path on Win+Myngw

if [[ $# -ne 2 ]]
  then
    echo "Usage: $0 (get|post|all) gopath"
    exit 1	
fi

godir=$2
work=$(dirname $0)
mkdir -p $work/published

function buildandpublish() {
	docker run --rm -v $godir:/go -v $godir/src/github.com/inge4pres/4pr.es/cmd/4pres_${1}:/tmp eawsy/aws-lambda-go
        mv $work/../cmd/4pres_${1}/handler.zip $work/
        aws lambda update-function-code  --function-name 4pres_lambda_${1}   --zip-file fileb://$work/handler.zip --publish
        mv $work/handler.zip $work/published/$(date +%s)_4pres_${1}_handler.zip
}


case "$1" in 
 "get")
   	buildandpublish get
	;;
 "post")
	buildandpublish post
	;;
 "delete")
 	buildandpublish delete
	;;
 "all")
	buildandpublish get
	buildandpublish post
	buildandpublish delete
	;;
esac

exit 0
