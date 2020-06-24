FROM golang:latest as go
FROM golangci/golangci-lint:v1.27-alpine

RUN apk update && apk add --no-cache --update python3 && apk add make && apk add bash

COPY . /go/src/github.com/kichyr/GoSearcher
WORKDIR /go/src/github.com/kichyr/GoSearcher

RUN pip3 install -r ./test/requirements.txt
CMD make test-local; make lint