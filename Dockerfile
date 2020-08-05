FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /go/src
COPY . /go/src

ENV GO111MODULE=on
RUN go mod download

RUN go get -u github.com/cosmtrek/air
ENTRYPOINT air
