FROM golang:1.13.3

ENV TZ="Asia/Tokyo" \
    GO111MODULE=on \
    GOBIN="${PWD}/bin" \
    PATH="{$GOBIN}:$PATH"

# set -x: docker build時にコマンドが出力
RUN set -x && \
    apt-get update && \
    apt-get install -y default-mysql-client vim

WORKDIR /go/src/trading_bitcoin_api

COPY . .

RUN go build
