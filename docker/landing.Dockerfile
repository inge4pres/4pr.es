FROM golang:1.9
MAINTAINER inge4pres

RUN mkdir -p /go/src/github.com/inge4pres/4pr.es

WORKDIR /go/src/github.com/inge4pres/4pr.es
COPY cmd cmd
COPY static static
RUN GOOS=linux go build ./cmd/landing/...

EXPOSE 8080
CMD ["./landing"]