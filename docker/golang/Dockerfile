FROM golang:1.13.3

ENV TZ="Asia/Tokyo"
ENV GO111MODULE=on

RUN apt-get update && \
    apt-get install -y default-mysql-client vim

WORKDIR /go/src/trading_bitcoin_api

ADD . .

RUN go build
