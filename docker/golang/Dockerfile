FROM golang:1.13.3

ENV TZ="Asia/Tokyo"
ENV GO111MODULE=on

RUN apt-get update && \
    apt-get install -y default-mysql-client vim && \
    export GOBIN=$PWD/bin $$ \ # ここ書き方が違うかも
    export PATH=$GOBIN:$PATH   # ここ書き方が違うかも

WORKDIR /go/src/trading_bitcoin_api

ADD . .

RUN go build
