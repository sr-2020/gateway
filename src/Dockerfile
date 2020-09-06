FROM golang:1.13-alpine as builder

WORKDIR /go/src/github.com/sr2020/gateway

ENV CGO_ENABLED=0

COPY . .

RUN go get github.com/smartystreets/goconvey

CMD cd tests && $GOPATH/bin/goconvey -host 0.0.0.0 -cover=false
