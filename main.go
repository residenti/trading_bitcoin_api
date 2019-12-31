package main

import (
	"fmt"
	"net/http"

	"github.com/residenti/trading_bitcoin_api/bitflyer"
	"github.com/residenti/trading_bitcoin_api/config"
	"github.com/residenti/trading_bitcoin_api/utils"
)

func init() {
	utils.InitSettingsOfLog(config.List.Logfile)
}

func main() {
	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":8080", nil)

	apiClinet := bitflyer.New()
	tickerChannel := make(chan bitflyer.Ticker)

	go apiClinet.SubscribeTicker(config.List.ProductCode, tickerChannel)

	for ticker := range tickerChannel {
		fmt.Println(ticker)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}
