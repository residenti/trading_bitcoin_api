FROM golang:1.13.3

# set -x: docker build時にコマンドが出力
RUN set -x && \
    apt-get update && \
    apt-get install -y default-mysql-client vim

WORKDIR /go/src/trading_bitcoin_api

ENV TZ="Asia/Tokyo" \
    GO111MODULE=on \
    GOBIN="/go/src/trading_bitcoin_api/bin" \
    PATH="/go/src/trading_bitcoin_api/bin:$PATH"

COPY . .

RUN go build
