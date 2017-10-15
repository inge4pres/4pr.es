# __4pr.es__ 
## A URL shortener in Go

[![Build Status](https://travis-ci.org/inge4pres/4pr.es.svg?branch=master)](https://travis-ci.org/inge4pres/4pr.es)

[4pres](http://4pr.es/) is yet another URL shortener, it will give you a short link for a long URL.
Short links have no expiration, for now :D, you can get short URL via any browser or other HTTP client.

4pres is written in [Go](http://golang.org/) and leverages a serverless architecture powered by AWS and coded with [eawsy](https://github.com/eawsy/aws-lambda-go-shim). It uses [Google SafeBrowsing](https://safebrowsing.google.com/) to validate the request URL is not from a malicious domain, so to prevent malware spread.  

#### API mapping
* `GET /` returns the home page
* `GET /s?URLENCODEDLINKTOSHORTEN` get you a HTTP200 with the HTML containing the short URL
* `GET /ANYURL` get you a HTTP301 redirecting to the original lonh URL or a HTTP404.
