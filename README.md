# trading_bitcoin_api

## Environment
- Go 1.13.3
- MySQL 8.0

## Setup
プロジェクトのルート直下で `docker-compose up -d` を実行すると、下記の2つのコンテナが立ち上がります。

- golang
- mysql

サーバーを起動する。

```
% docker exec -it golang bash
# go build -o main
# ./main
```

`go build` を実行すると `go.sum`, `go.mod` に記録されているパッケージが自動でインストールされます。

## Package management
Go Moudules の module-aware mode でパッケージを管理しています。  
module-aware mode を有効にするために `GO111MODULE=on` としています。  
(プロジェクト配下に `go.sum`, `go.mod` が配置されているため `GO111MODULE=auto` にしても module-aware mode となるが分かりやすいので明示的に `on` を指定している)

注意点として、ツール系のパッケージは `tools/tool.go` にて blank import してください。  

実行ファイルは `GOBIN="/go/src/trading_bitcoin_api/bin"` に配置されます。  
詳細は[こちら](https://marcofranssen.nl/manage-go-tools-via-go-modules/)を参考にしてください。
