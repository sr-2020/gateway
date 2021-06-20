FROM golang:1.16.5-alpine as builder

WORKDIR /go/src/github.com/sr2020/gateway

COPY ./src/go.mod .
COPY ./src/go.sum .

RUN go mod download

COPY ./src .

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/gateway ./cmd


FROM alpine:3.12.1

WORKDIR /root/

COPY --from=builder /go/bin/gateway .

EXPOSE 80

CMD ["./gateway"]
