FROM golang:1.12.7-alpine3.10
RUN apk add --update --no-cache make

COPY . /go/src/github.com/tcd/prjr
WORKDIR /go/src/github.com/tcd/prjr
