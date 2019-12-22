FROM golang:1.8

RUN apt-get update && \
    apt-get install -y mysql-client vim

WORKDIR /go/src/trading_bitcoin_api

ADD . .

RUN go build
